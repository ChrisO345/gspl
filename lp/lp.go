package lp

import "gonum.org/v1/gonum/mat"

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
	VariablesMap     []string
	Status           LpStatus
	Sense            LpSense
	ConstraintVector []LpConstraintType
}

// NewLinearProgram Create a new Linear Program
func NewLinearProgram(desc string, vars []LpVariable) LinearProgram {

	lp := LinearProgram{
		Description:  desc,
		VariablesMap: make([]string, len(vars)),
	}

	for i, v := range vars {
		lp.VariablesMap[i] = v.Name
	}

	lp.Status = LpStatusNotSolved
	return lp
}
