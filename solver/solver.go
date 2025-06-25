package solver

import (
	"fmt"
	"math"

	"github.com/chriso345/gspl/internal/matrix"
	"github.com/chriso345/gspl/internal/simplex"
	"github.com/chriso345/gspl/lp"
	"gonum.org/v1/gonum/mat"
)

// AddObjective Add an objective to the linear program
func AddObjective(prog *lp.LinearProgram, sense lp.LpSense, objective lp.LpExpression) {
	prog.Sense = sense
	// lp.ObjectiveFunction = objective
	prog.ObjectiveFunc = mat.NewVecDense(len(objective.Terms), nil)
	for i, v := range objective.Terms {
		mappedIndex := -1
		for j, varName := range prog.VariablesMap {
			if varName == v.Variable.Name {
				mappedIndex = j
				break
			}
		}
		if mappedIndex == -1 {
			panic(fmt.Sprintf("Variable %s not found in Linear Program", v.Variable.Name))
		}
		prog.ObjectiveFunc.SetVec(i, v.Coefficient)
	}

	if prog.Sense == lp.LpMaximise {
		for i := range prog.ObjectiveFunc.RawVector().N {
			prog.ObjectiveFunc.SetVec(i, -prog.ObjectiveFunc.At(i, 0))
		}
	}
}

// AddConstraint Add a constraint to the linear program
func AddConstraint(prog *lp.LinearProgram, constraint lp.LpExpression, constraintType lp.LpConstraintType, rightHandSide float64) {
	// Panic if objective function is not set
	if prog.ObjectiveFunc == nil {
		panic("Objective function not set")
	}

	prog.ConstraintVector = append(prog.ConstraintVector, constraintType)

	if rightHandSide < 0 {
		// Multiply the constraint by -1, flip equality sign
		rightHandSide = math.Abs(rightHandSide)
		for i := range constraint.Terms {
			constraint.Terms[i].Coefficient *= -1
		}
		constraintType = -constraintType
	}

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
		for j, varName := range prog.VariablesMap {
			if varName == v.Variable.Name {
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

func Solve(prog *lp.LinearProgram, opts ...SolverOption) *lp.LinearProgram {
	// Add slacks for non-equality constraints
	for i, constraintType := range prog.ConstraintVector {
		if constraintType != lp.LpConstraintEQ {
			slack := lp.NewVariable(fmt.Sprintf("s%d", i))
			prog.VariablesMap = append(prog.VariablesMap, slack.Name)
			unitVector := mat.NewDense(prog.Constraints.RawMatrix().Rows, 1, nil)
			one := 1.0
			if constraintType == lp.LpConstraintGE {
				one = -1.0
			}
			unitVector.Set(i, 0, one)
			prog.Constraints = matrix.ResizeMatDense(prog.Constraints, prog.Constraints.RawMatrix().Rows, len(prog.VariablesMap))

			prog.ObjectiveFunc = matrix.ResizeVecDense(prog.ObjectiveFunc, prog.ObjectiveFunc.RawVector().N+1)
			prog.ObjectiveFunc.SetVec(prog.ObjectiveFunc.RawVector().N-1, 0)
		}
	}

	m := prog.Constraints.RawMatrix().Rows
	n := len(prog.VariablesMap)

	z, x, _, idx, flag := simplex.Simplex(prog.Constraints, prog.RHS, prog.ObjectiveFunc, m, n)

	prog.Solution = z
	prog.Variables = x
	prog.Indices = idx

	switch flag {
	case 0:
		prog.Status = lp.LpStatusOptimal
	case 1:
		prog.Status = lp.LpStatusInfeasible
	case -1:
		prog.Status = lp.LpStatusUnbounded
	}

	if prog.Sense == lp.LpMaximise {
		prog.Solution *= -1
	}

	return prog
}
