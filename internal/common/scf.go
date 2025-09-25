package common

import "gonum.org/v1/gonum/mat"

// StandardComputationalForm represents a linear programming problem in standard form.
type StandardComputationalForm struct {
	Objective   *mat.VecDense // c
	Constraints *mat.Dense    // A
	RHS         *mat.VecDense // b

	PrimalSolution *mat.VecDense // x*

	ObjectiveValue *float64
	Status         *SolverStatus // Optimal, Infeasible, Unbounded, etc.
}
