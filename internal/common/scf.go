package common

import (
	"gonum.org/v1/gonum/mat"
)

// StandardComputationalForm represents a linear programming problem in standard form.
type StandardComputationalForm struct {
	Objective   *mat.VecDense // c
	Constraints *mat.Dense    // A
	RHS         *mat.VecDense // b

	PrimalSolution *mat.VecDense // x*

	ObjectiveValue *float64
	Status         *SolverStatus // Optimal, Infeasible, Unbounded, etc.
	SlackIndices   []int         // Indices of slack variables in the solution
	NumPrimals     int           // Number of primal variables (non-slack)
}

// Copy creates a deep copy of the SCF
func (scf *StandardComputationalForm) Copy() *StandardComputationalForm {
	return &StandardComputationalForm{
		Objective:      mat.VecDenseCopyOf(scf.Objective),
		Constraints:    mat.DenseCopyOf(scf.Constraints),
		RHS:            mat.VecDenseCopyOf(scf.RHS),
		PrimalSolution: mat.VecDenseCopyOf(scf.PrimalSolution),
		ObjectiveValue: scf.ObjectiveValue,
		Status:         scf.Status,
		SlackIndices:   scf.SlackIndices,
		NumPrimals:     scf.NumPrimals,
	}
}

// AddBranch adds a new constraint to the SCF
func (scf *StandardComputationalForm) AddBranch(idx int, rhs float64, dir int) {
	numRows, numCols := scf.Constraints.Dims()
	newConstraints := mat.NewDense(numRows+1, numCols, nil)
	newRHS := mat.NewVecDense(numRows+1, nil)
	for i := range numRows {
		for j := range numCols {
			newConstraints.Set(i, j, scf.Constraints.At(i, j))
		}
		newRHS.SetVec(i, scf.RHS.AtVec(i))
	}

	for j := range numCols {
		if j == idx {
			switch dir {
			case 1:
				newConstraints.Set(numRows, j, 1)
			case 2:
				newConstraints.Set(numRows, j, -1)
			}
		} else {
			newConstraints.Set(numRows, j, 0)
		}
	}
	switch dir {
	case 1:
		newRHS.SetVec(numRows, rhs)
	case 2:
		newRHS.SetVec(numRows, -rhs)
	}
	scf.Constraints = newConstraints
	scf.RHS = newRHS
}
