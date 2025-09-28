package simplex

import (
	"errors"
	"fmt"
	"math"

	"github.com/chriso345/gspl/internal/common"
	"github.com/chriso345/gspl/internal/matrix"
	"gonum.org/v1/gonum/mat"
)

func Simplex(scf *common.StandardComputationalForm) error {
	m, n := scf.Constraints.Dims()
	sm := &simplexMethod{
		m: m,
		n: n,
	}

	I := matrix.Eye(m)

	// Phase 1: Set up the auxilary problem
	sm.A = mat.NewDense(m, n+m, nil)
	for i := range m {
		for j := range n {
			sm.A.Set(i, j, scf.Constraints.At(i, j))
		}
		for j := range m {
			sm.A.Set(i, n+j, I.At(i, j))
		}
	}

	// Construct the cost vector for Phase 1, [0,...,0,1,...,1]
	sm.c = mat.NewVecDense(n+m, nil)
	for i := range n + m {
		if i < n {
			sm.c.SetVec(i, 0.)
		} else {
			sm.c.SetVec(i, 1.)
		}
	}

	sm.cb = mat.NewVecDense(m, nil)
	for i := range m {
		sm.cb.SetVec(i, float64(n+i)) // Artificial variables as initial basis
	}

	sm.B = matrix.ExtractColumns(sm.A, sm.cb)
	sm.b = scf.RHS

	// Run Phase 1 of the RSM
	err := RSM(sm, 1)
	if err != nil {
		return fmt.Errorf("error in Phase 1 of Simplex: %w", err)
	}

	// Check infeasibility
	if sm.rsmResult.flag == common.SolverStatusOptimal && sm.rsmResult.value > 1e-8 { // TODO: replace with tolerance option
		*scf.Status = common.SolverStatusInfeasible
		return nil
	}

	// Phase 2: Set up the original problem

	// OPTIM: Is the A matrix changed by RSM? if not we can skip this reconstruction
	sm.A = mat.NewDense(m, n+m, nil)
	for i := range m {
		for j := range n {
			sm.A.Set(i, j, scf.Constraints.At(i, j))
		}
		for j := range m {
			sm.A.Set(i, n+j, I.At(i, j))
		}
	}

	sm.c = mat.NewVecDense(n+m, nil) // FIXME:???
	for i := range n {
		sm.c.SetVec(i, scf.Objective.AtVec(i))
	}
	for i := range m {
		sm.c.SetVec(n+i, 0.)
	}

	sm.cb = sm.rsmResult.indices
	sm.B = matrix.ExtractColumns(sm.A, sm.cb)
	sm.b = scf.RHS // Should be unchanged? OPTIM: reuse from Phase 1?

	// Run Phase 2 of the RSM
	err = RSM(sm, 2)
	if err != nil {
		return fmt.Errorf("error in Phase 2 of Simplex: %w", err)
	}
	*scf.Status = sm.rsmResult.flag
	if sm.rsmResult.flag == common.SolverStatusOptimal {
		*scf.ObjectiveValue = sm.rsmResult.value
		scf.PrimalSolution = sm.rsmResult.x
	}

	return nil
}

func RSM(sm *simplexMethod, phase int) error {
	_maxIter := 1000 // Simple safeguard

	n := sm.n
	if phase == 1 {
		n += sm.m
	}

	// Initialise the rsmResult
	sm.rsmResult = rsmResult{
		flag:    common.SolverStatusNotSolved,
		value:   0., // z
		x:       mat.NewVecDense(n, nil),
		indices: sm.cb,
	}

	// Initialise other variables
	B := sm.B // OPTIM: do we need to copy this?
	cb := mat.NewVecDense(sm.m, nil)
	for i := range sm.m {
		index := int(sm.rsmResult.indices.AtVec(i))
		cb.SetVec(i, sm.c.AtVec(index))
	}

	for range _maxIter {
		xb := mat.NewVecDense(sm.m, nil) // Basic solution
		err := xb.SolveVec(B, sm.b)
		if err != nil {
			// Basis is singular, return error
			return fmt.Errorf("error solving for basic solution: %w", err)
		}

		// Finding the leaving variable
		var BT mat.Dense
		BT.CloneFrom(B.T())

		sm.rsmResult.pi = mat.NewVecDense(sm.m, nil) // Dual variables
		err = sm.rsmResult.pi.SolveVec(&BT, cb)
		if err != nil {
			// Basis is singular, return error
			return fmt.Errorf("error solving for dual variables: %w", err)
		}

		fe := enterStruct{
			A:       sm.A,
			pi:      sm.rsmResult.pi,
			c:       sm.c,
			isbasic: mat.NewVecDense(n, nil),
		}

		for i := range sm.m {
			index := int(sm.rsmResult.indices.AtVec(i))
			if index < n {
				fe.isbasic.SetVec(index, 1.)
			}
		}

		err = findEnter(&fe)
		if err != nil {
			return fmt.Errorf("error finding entering variable: %w", err)
		}

		if fe.s == -1 {
			// Optimal solution found
			sm.rsmResult.flag = common.SolverStatusOptimal
			sm.rsmResult.value = 0.
			for i := range sm.m {
				index := int(sm.rsmResult.indices.AtVec(i))
				sm.rsmResult.x.SetVec(index, xb.AtVec(i))
				sm.rsmResult.value += cb.AtVec(i) * xb.AtVec(i)
			}
			return nil
		}

		// Finding the leaving variable
		fl := leaveStruct{
			B:       B,
			indices: sm.rsmResult.indices,
			as:      fe.as,
			xb:      xb,
			phase:   phase,
			n:       n,
		}

		err = findLeave(&fl)
		if err != nil {
			return fmt.Errorf("error finding leaving variable: %w", err)
		}

		if fl.r == -1 {
			// Unbounded solution
			sm.rsmResult.flag = common.SolverStatusUnbounded
			return nil
		}

		// Update B, cb, and indices
		bu := bUpdateStruct{
			BMat:    B,
			indices: sm.rsmResult.indices,
			cb:      cb,
			as:      fe.as,
			s:       fe.s,
			r:       fl.r,
			cs:      fe.cs,
		}

		err = updateB(&bu)

	}
	return errors.New("max iterations reached in RSM")
}

func findEnter(fe *enterStruct) error {
	fe.s = -1
	fe.as = nil
	fe.cs = 0.
	minrc := math.Inf(1)
	tol := -1e-6 // TODO: replace with tolerance option

	n := fe.isbasic.Len()

	for j := range n {
		if fe.isbasic.AtVec(j) == 0 {
			m, _ := fe.A.Dims()
			aj := mat.NewVecDense(m, nil)
			for i := range m {
				aj.SetVec(i, fe.A.At(i, j))
			}

			dot := mat.Dot(fe.pi, aj)
			rc := fe.c.AtVec(j) - dot

			if rc < minrc {
				minrc = rc
				fe.s = j
				fe.cs = fe.c.AtVec(j)

				// OPTIM: don't reallocate this every loop...
				fe.as = mat.NewVecDense(m, nil)
				for i := range m {
					fe.as.SetVec(i, fe.A.At(i, j))
				}
			}
		}
	}

	if minrc >= tol {
		m, _ := fe.A.Dims()
		fe.s = -1
		fe.as = mat.NewVecDense(m, nil)
		fe.cs = 0.
	}

	return nil
}

func findLeave(fl *leaveStruct) error {
	fl.r = -1

	var Binv mat.Dense
	if err := Binv.Inverse(fl.B); err != nil {
		return fmt.Errorf("error inverting basis matrix: %w", err)
	}

	directionVec := mat.NewVecDense(fl.as.Len(), nil)
	directionVec.MulVec(&Binv, fl.as)

	m := fl.xb.Len()
	theta := math.Inf(1)

	for i := range m {
		dirVal := directionVec.AtVec(i)
		indexVal := int(fl.indices.AtVec(i))

		if fl.phase == 2 && indexVal > fl.n {
			if dirVal != 0 {
				fl.r = i
				return nil
			}
		} else {
			if dirVal > 0 {
				ratio := fl.xb.AtVec(i) / dirVal
				if ratio < theta {
					theta = ratio
					fl.r = i
				}
			}
		}
	}

	return nil
}

func updateB(bu *bUpdateStruct) error {
	m, _ := bu.BMat.Dims()

	for i := range m {
		bu.BMat.Set(i, bu.r, bu.as.AtVec(i))
	}

	bu.indices.SetVec(bu.r, float64(bu.s))
	bu.cb.SetVec(bu.r, bu.cs)

	return nil
}
