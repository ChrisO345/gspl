package gspl

import (
	"math"
)

func Simplex(A, b, c *Matrix, m, n int) (z float64, x, piValues, indices *Matrix, exitflag int) {
	// Create identity matrix I of size m
	I := Eye(m)

	// Construct A_phase1 = [A | I]
	A_phase1 := NewMatrix(m, n+m)
	for i := range m {
		for j := range n {
			A_phase1.Values[i][j] = A.Values[i][j]
		}
		for j := range m {
			A_phase1.Values[i][n+j] = I.Values[i][j]
		}
	}

	// Construct c_phase1 = [0's for x vars; 1's for artificial vars]
	c_phase1 := NewMatrix(n+m, 1)
	for i := n; i < n+m; i++ {
		c_phase1.Values[i][0] = 1.0
	}

	// Initial basic indices: the artificial variables
	indicesInit := NewMatrix(m, 1)
	for i := range m {
		indicesInit.Values[i][0] = float64(n + i)
	}

	// Initial B matrix from A_phase1 columns at artificial indices
	Bmatrix := A_phase1.ExtractColumns(indicesInit)

	// Run Phase 1 revised simplex
	z1, xPhase1, piValues1, indicesPhase1, exitflag1 := RevisedSimplex(A_phase1, b, c_phase1, m, n+m, Bmatrix, indicesInit, 1)

	if exitflag1 == 0 && z1 > 0 {
		// Infeasible
		return z1, xPhase1, piValues1, indicesPhase1, 1
	}

	finalIndices := indicesPhase1

	// Phase 2: Solve actual problem
	A_phase2 := NewMatrix(m, n+m)
	for i := range m {
		for j := range n {
			A_phase2.Values[i][j] = A.Values[i][j]
		}
		for j := range m {
			A_phase2.Values[i][n+j] = I.Values[i][j]
		}
	}

	// Extend c to include 0s for artificial variables
	cExtended := NewMatrix(n+m, 1)
	for i := range n {
		cExtended.Values[i][0] = c.Values[i][0]
	}

	Bmatrix = A_phase2.ExtractColumns(finalIndices)

	// Run Phase 2 revised simplex
	z, x, piValues, indices, exitflag2 := RevisedSimplex(A_phase2, b, cExtended, m, n, Bmatrix, finalIndices, 2)

	if exitflag2 == -1 {
		exitflag = -1 // unbounded
	} else {
		exitflag = 0 // success
	}

	return
}

func RevisedSimplex(A, b, c *Matrix, m, n int, Bmatrix, indices_ *Matrix, phase int) (z float64, x, pivalues, indices *Matrix, exitflag int) {
	exitflag = 0
	x = NewMatrix(n, 1)
	B := Bmatrix
	indices = indices_

	// Calculate the cb vector (cost of basic variables)
	cb := NewMatrix(m, 1) // Initialize a column vector for cb (same number of rows as the number of basic variables)

	for i := range m {
		index := int(indices.Values[i][0])   // Get the index of the basic variable from indices matrix
		cb.Values[i][0] = c.Values[index][0] // Assign the cost from c based on the index
	}

	for {

		xb := B.Inv().Mul(b)

		pivalues = B.Transpose().Inv().Mul(cb)

		isbasic := NewMatrix(1, n)
		for i := range m {
			index := int(indices.Values[i][0])
			if index < n {
				isbasic.Values[0][index] = 1 // Basic variable
			}
		}

		as, cs, s := findEnter(A, pivalues, c, isbasic, phase)

		if s == -1 {
			for i := range m {
				idx := int(indices.Values[i][0])
				if idx < n {
					x.Values[idx][0] = xb.Values[i][0]
				}
			}
			z = 0
			for i := range n {
				z += c.Values[i][0] * x.Values[i][0]
			}
			return
		}

		leave := findLeave(B, as, xb, phase, n, indices)
		if leave == -1 {
			for i := range m {
				idx := int(indices.Values[i][0])
				if idx < n {
					x.Values[idx][0] = xb.Values[i][0]
				}
			}
			z = 0
			for i := range n {
				z += c.Values[i][0] * x.Values[i][0]
			}
			exitflag = -1
			return
		}

		bUpdate(B, indices, cb, as, s, leave, cs)
	}
}

func findEnter(A, pi, c, isbasic *Matrix, phase int) (as *Matrix, cs float64, s int) {
	s = -1
	as = nil
	cs = 0.0
	minrc := math.Inf(1)
	tolerance := -1.0e-6

	n := isbasic.Length()

	for j := range n {
		if isbasic.Get(0, j) == 0 {
			aj := &Matrix{Rows: A.Rows, Columns: 1, Values: make([][]float64, A.Rows)}
			for i := range A.Rows {
				aj.Values[i] = []float64{A.Values[i][j]}
			}
			rc := c.Values[j][0] - Dot(pi, aj)
			if rc < minrc {
				minrc = rc
				s = j
				as = aj
				cs = c.Values[j][0]
			}
		}
	}

	if minrc > tolerance {
		as = NewMatrix(A.Rows, 1)
		cs = 0.0
		s = -1
	}

	return as, cs, s
}

func findLeave(B *Matrix, as *Matrix, xb *Matrix, phase int, n int, indices *Matrix) (leave int) {
	leave = -1
	direction := B.Inv().Mul(as)
	theta := math.Inf(1)
	m := xb.Rows

	for i := range m {
		dirVal := direction.Values[i][0]
		indexVal := int(indices.Values[i][0])

		if phase == 2 && indexVal > n {
			if dirVal != 0 {
				leave = i
				return
			}
		} else {
			if dirVal > 0 {
				ratio := xb.Values[i][0] / dirVal
				if ratio < theta {
					theta = ratio
					leave = i
				}
			}
		}
	}
	return
}

func bUpdate(Bmatrix, indices, cb, as *Matrix, s, leave int, cs float64) {
	// Replace column `leave` in Bmatrix with `as` (assuming `as` is a column vector)
	for i := range Bmatrix.Rows {
		Bmatrix.Values[i][leave] = as.Values[i][0]
	}

	// Update `indices` at the row `leave` with entering variable index `s`
	indices.Values[leave][0] = float64(s)

	// Update `cb` at the row `leave` with cost `cs`
	cb.Values[leave][0] = cs
}
