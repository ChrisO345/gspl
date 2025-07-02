package simplex

import (
	"math"

	"github.com/chriso345/gspl/internal/matrix"
	"gonum.org/v1/gonum/mat"
)

// Simplex solves the LP problem using the standard simplex method.
//
// It calls RevisedSimplex for both phase 1 and phase 2 internally.
func Simplex(A *mat.Dense, b, c *mat.VecDense, m, n int) (z float64, x, piValues, indices *mat.VecDense, exitflag int) {
	// TODO: Replace exitflag integer implementation with proper error handling or struct
	// Create identity matrix I of size m
	I := matrix.Eye(m)

	// Construct A_phase1 = [A | I]
	A_phase1 := mat.NewDense(m, n+m, nil)
	for i := range m {
		for j := range n {
			A_phase1.Set(i, j, A.At(i, j))
		}
		for j := range m {
			A_phase1.Set(i, n+j, I.At(i, j))
		}
	}

	// Construct c_phase1 = [0's for x vars; 1's for artificial vars]
	c_phase1 := mat.NewVecDense(n+m, nil)
	for i := range n + m {
		if i < n {
			c_phase1.SetVec(i, 0.0) // 0 for original variables
		} else {
			c_phase1.SetVec(i, 1.0) // 1 for artificial variables
		}
	}

	// Initial basic indices: the artificial variables
	indicesInit := mat.NewVecDense(m, nil)
	for i := range m {
		indicesInit.SetVec(i, float64(n+i)) // Artificial variables indices
	}

	// Initial B matrix from A_phase1 columns at artificial indices
	// Bmatrix := A_phase1.ExtractColumns(indicesInit)
	Bmatrix := matrix.ExtractColumns(A_phase1, indicesInit)

	// Run Phase 1 revised simplex
	z1, xPhase1, piValues1, indicesPhase1, exitflag1 := RevisedSimplex(A_phase1, b, c_phase1, m, n+m, Bmatrix, indicesInit, 1)

	if exitflag1 == 0 && z1 > 0 {
		// Infeasible
		return z1, xPhase1, piValues1, indicesPhase1, 1
	}

	finalIndices := indicesPhase1

	// Phase 2: Solve actual problem
	A_phase2 := mat.NewDense(m, n+m, nil)
	for i := range m {
		// TODO: Optimize this into a single loop
		for j := range n {
			A_phase2.Set(i, j, A.At(i, j))
		}
		for j := range m {
			A_phase2.Set(i, n+j, I.At(i, j))
		}
	}

	// Extend c to include 0s for artificial variables
	cExtended := mat.NewVecDense(n+m, nil)
	for i := range n {
		cExtended.SetVec(i, c.At(i, 0))
	}

	Bmatrix = matrix.ExtractColumns(A_phase2, finalIndices)

	// Run Phase 2 revised simplex
	z, x, piValues, indices, exitflag2 := RevisedSimplex(A_phase2, b, cExtended, m, n, Bmatrix, finalIndices, 2)

	if exitflag2 == -1 {
		exitflag = -1 // unbounded
	} else {
		exitflag = 0 // success
	}

	return
}

// RevisedSimplex implements the revised simplex algorithm.
//
// It solves the LP in the specified phase (1 or 2) given an initial basis.
// Returns optimal objective, solution, duals, basis indices, and exit flag.
func RevisedSimplex(A *mat.Dense, b, c *mat.VecDense, m, n int, Bmatrix *mat.Dense, indices_ *mat.VecDense, phase int) (z float64, x, pivalues, indices *mat.VecDense, exitflag int) {
	exitflag = 0
	x = mat.NewVecDense(n, nil) // Initialize x as a vector of size n
	B := Bmatrix
	indices = indices_

	// Calculate the cb vector (cost of basic variables)
	cb := mat.NewVecDense(m, nil)

	for i := range m {
		index := int(indices.AtVec(i)) // Get the index of the basic variable from indices vector
		cb.SetVec(i, c.AtVec(index))   // Assign the cost from c based on the index
	}

	for {

		xb := mat.NewVecDense(b.Len(), nil)
		err := xb.SolveVec(Bmatrix, b)
		if err != nil {
			exitflag = -1 // Singular matrix or unbounded
			return
		}

		var BT mat.Dense
		BT.CloneFrom(B.T()) // B transpose

		pivalues = mat.NewVecDense(cb.Len(), nil)
		err = pivalues.SolveVec(&BT, cb)
		if err != nil {
			exitflag = -1 // Singular matrix
		}

		isbasic := mat.NewVecDense(n, nil) // Initialize isbasic as a row vector of size n
		for i := range m {
			index := int(indices.AtVec(i)) // Get the index of the basic variable from indices vector
			if index < n {
				isbasic.SetVec(index, 1) // Basic variable
			}
		}

		as, cs, s := findEnter(A, pivalues, c, isbasic)

		if s == -1 {
			for i := range m {
				idx := int(indices.AtVec(i)) // Get the index of the basic variable from indices vector
				if idx < n {
					x.SetVec(idx, xb.AtVec(i)) // Assign the value from xb to x at index idx
				}
			}
			z = 0
			for i := range n {
				z += c.At(i, 0) * x.AtVec(i) // Calculate the objective value
			}
			return
		}

		leave := findLeave(B, as, xb, indices, phase, n)
		if leave == -1 {
			for i := range m {
				idx := int(indices.At(i, i)) // Get the index of the basic variable from indices vector
				if idx < n {
					x.SetVec(idx, xb.AtVec(i)) // Assign the value from xb to x at index idx
				}
			}
			z = 0
			for i := range n {
				z += c.At(i, 0) * x.AtVec(i) // Calculate the objective value
			}
			exitflag = -1
			return
		}

		bUpdate(B, indices, cb, as, s, leave, cs)
	}
}

// findEnter identifies the entering variable in the revised simplex method.
func findEnter(A *mat.Dense, pi, c, isbasic *mat.VecDense) (as *mat.VecDense, cs float64, s int) {
	s = -1
	as = nil
	cs = 0.0
	minrc := math.Inf(1)
	tolerance := -1.0e-6

	n := isbasic.Len()

	for j := range n {
		if isbasic.AtVec(j) == 0 {
			m, _ := A.Dims()
			aj := mat.NewVecDense(m, nil)
			for i := range m {
				aj.SetVec(i, A.At(i, j))
			}

			dot := mat.Dot(pi, aj)
			rc := c.AtVec(j) - dot

			if rc < minrc {
				minrc = rc
				s = j
				cs = c.AtVec(j)

				as = mat.NewVecDense(m, nil)
				for i := range m {
					as.SetVec(i, A.At(i, j))
				}
			}
		}
	}

	if minrc > tolerance {
		m, _ := A.Dims()
		as = mat.NewVecDense(m, nil) // zero vector column
		cs = 0.0
		s = -1
	}

	return as, cs, s
}

// findLeave identifies the leaving variable in the revised simplex method.
func findLeave(B *mat.Dense, as, xb, indices *mat.VecDense, phase, n int) int {
	leave := -1

	var Binv mat.Dense
	if err := Binv.Inverse(B); err != nil {
		return -1
	}

	colVec := mat.NewVecDense(as.RawVector().N, mat.Col(nil, 0, as))
	directionVec := mat.NewVecDense(as.RawVector().N, nil)
	directionVec.MulVec(&Binv, colVec)

	m := xb.Len()
	theta := math.Inf(1)

	for i := range m {
		dirVal := directionVec.AtVec(i)
		indexVal := int(indices.AtVec(i))

		if phase == 2 && indexVal > n {
			if dirVal != 0 {
				leave = i
				return leave
			}
		} else {
			if dirVal > 0 {
				ratio := xb.AtVec(i) / dirVal
				if ratio < theta {
					theta = ratio
					leave = i
				}
			}
		}
	}

	return leave
}

// bUpdate updates the B matrix, indices, cb, and as vectors during the revised simplex method.
func bUpdate(Bmatrix *mat.Dense, indices, cb, as *mat.VecDense, s, leave int, cs float64) {
	m, _ := Bmatrix.Dims()

	for i := range m {
		Bmatrix.Set(i, leave, as.AtVec(i))
	}

	indices.SetVec(leave, float64(s))
	cb.SetVec(leave, cs)
}
