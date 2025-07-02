package lp

import "gonum.org/v1/gonum/mat"

// LinearProgram represents a linear programming problem in standard form.
type LinearProgram struct {
	// Solution
	Solution float64 // z

	// Matrix Representation
	Variables     *mat.VecDense // x
	ObjectiveFunc *mat.VecDense // c
	Constraints   *mat.Dense    // A
	RHS           *mat.VecDense // b

	// Simplex Internal Variables
	Indices  *mat.VecDense // Indices of basic variables
	pivalues *mat.VecDense // Pivot values
	bMatrix  *mat.Dense    // Basis matrix
	cb       *mat.VecDense // Coefficients of the basis variables

	// Others
	Description      string
	VariablesMap     []LpVariable
	Status           LpStatus
	Sense            LpSense
	ConstraintVector []LpConstraintType
}

// NewLinearProgram Create a new Linear Program
func NewLinearProgram(desc string, vars []LpVariable) LinearProgram {

	lp := LinearProgram{
		Description:  desc,
		VariablesMap: make([]LpVariable, len(vars)),
	}

	for i, v := range vars {
		lp.VariablesMap[i] = v
	}

	lp.Status = LpStatusNotSolved
	return lp
}
