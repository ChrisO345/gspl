package lp

import "gonum.org/v1/gonum/mat"

// LinearProgram represents a linear programming problem in standard form.
type LinearProgram struct {
	// Problem definition
	Objective   *mat.VecDense      // c
	Constraints *mat.Dense         // A
	RHS         *mat.VecDense      // b
	Sense       LpSense            // Minimize or Maximize
	ConTypes    []LpConstraintType // metadata for constraints
	Vars        []LpVariable       // metadata for variables

	// Solution
	ObjectiveValue float64
	PrimalSolution *mat.VecDense // x*
	// DualSolution   *mat.VecDense // y*
	Status LpStatus // Optimal, Infeasible, Unbounded, etc.

	// Simplex internal state (keep unexported)
	indices  *mat.VecDense
	piValues *mat.VecDense
	basis    *mat.Dense
	cb       *mat.VecDense

	// Metadata
	Description string
}

// NewLinearProgram Create a new Linear Program
func NewLinearProgram(desc string, vars []LpVariable) LinearProgram {

	lp := LinearProgram{
		Description: desc,
		Vars:        make([]LpVariable, len(vars)),
	}

	copy(lp.Vars, vars)

	lp.Status = LpStatusNotSolved
	return lp
}
