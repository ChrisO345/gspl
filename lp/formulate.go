package lp

import (
	"fmt"
	"math"

	"github.com/chriso345/gspl/internal/matrix"
	"gonum.org/v1/gonum/mat"
)

// AddObjective Add an objective to the linear program
func (prog *LinearProgram) AddObjective(sense LpSense, objective LpExpression) {
	prog.Sense = sense
	// ObjectiveFunction = objective
	prog.ObjectiveFunc = mat.NewVecDense(len(objective.Terms), nil)
	for i, v := range objective.Terms {
		mappedIndex := -1
		for j, lpVar := range prog.VariablesMap {
			if lpVar.Name == v.Variable.Name {
				mappedIndex = j
				break
			}
		}
		if mappedIndex == -1 {
			panic(fmt.Sprintf("Variable %s not found in Linear Program", v.Variable.Name))
		}
		prog.ObjectiveFunc.SetVec(i, v.Coefficient)
	}

	if prog.Sense == LpMinimise {
		for i := range prog.ObjectiveFunc.RawVector().N {
			prog.ObjectiveFunc.SetVec(i, -prog.ObjectiveFunc.At(i, 0))
		}
	}
}

// AddConstraint Add a constraint to the linear program
func (prog *LinearProgram) AddConstraint(constraint LpExpression, constraintType LpConstraintType, rightHandSide float64) {
	// Panic if objective function is not set
	if prog.ObjectiveFunc == nil {
		panic("Objective function not set")
	}

	if rightHandSide < 0 {
		// Multiply the constraint by -1, flip equality sign
		rightHandSide = math.Abs(rightHandSide)
		for i := range constraint.Terms {
			constraint.Terms[i].Coefficient *= -1
		}
		constraintType = -constraintType
	}

	prog.ConstraintVector = append(prog.ConstraintVector, constraintType)

	currentRow := 0
	if prog.Constraints == nil {
		prog.Constraints = mat.NewDense(1, len(prog.VariablesMap), nil)
		prog.RHS = mat.NewVecDense(1, nil)
	} else {
		currentRow = prog.Constraints.RawMatrix().Rows
		prog.Constraints = matrix.ResizeMatDense(prog.Constraints, currentRow+1, len(prog.VariablesMap))
		prog.RHS = matrix.ResizeVecDense(prog.RHS, currentRow+1)
	}

	newRow := make([]float64, len(prog.VariablesMap))
	for _, v := range constraint.Terms {
		mappedIndex := -1
		for j, lpVar := range prog.VariablesMap {
			if lpVar.Name == v.Variable.Name {
				mappedIndex = j
				break
			}
		}
		if mappedIndex == -1 {
			panic(fmt.Sprintf("Variable %s not found in Linear Program", v.Variable.Name))
		}
		newRow[mappedIndex] = v.Coefficient
	}

	prog.Constraints.SetRow(currentRow, newRow)
	prog.RHS.SetVec(currentRow, rightHandSide)
}

// AddIPConstraints adds integer programming constraints to the linear program.
func (prog *LinearProgram) AddIPConstraints() {
	if prog.VariablesMap == nil {
		return
	}

	for _, v := range prog.VariablesMap {
		if v.Category == LpCategoryBinary {
			constraint := LpExpression{
				Terms: []LpTerm{
					{Variable: v, Coefficient: 1},
				},
			}
			prog.AddConstraint(constraint, LpConstraintEQ, 1)
		}
	}
}
