package simplex

import (
	"errors"
	"fmt"

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

	sm.cb = mat.NewVecDense(n+m, nil)
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

	sm.c = mat.NewVecDense(n+m, nil)
	for i := range n + m {
		sm.c.SetVec(i, scf.Objective.AtVec(i))
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

	return errors.New("Simplex algorithm not yet implemented")
}

func RSM(sm *simplexMethod, phase int) error {
	return errors.New("Revised Simplex Method not yet implemented")
}

func printFormatMatrix(A *mat.Dense) {
	frmt := mat.Formatted(A, mat.Prefix(""), mat.Excerpt(0))
	fmt.Printf("\n%v\n", frmt)
}
