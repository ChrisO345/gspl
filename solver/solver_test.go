package solver_test

import (
	"testing"

	"github.com/chriso345/gspl/internal/testutils/assert"
	"github.com/chriso345/gspl/lp"
	"github.com/chriso345/gspl/solver"
)

func constructTestProgram(prog *lp.LinearProgram, variables []lp.LpVariable) {
	x1 := &variables[0]
	x2 := &variables[1]

	objective := lp.NewExpression([]lp.LpTerm{
		lp.NewTerm(1, *x1),
		lp.NewTerm(1, *x2),
	})

	prog.AddObjective(lp.LpMaximise, objective)

	prog.AddConstraint(lp.NewExpression([]lp.LpTerm{
		lp.NewTerm(1, *x1),
		lp.NewTerm(2, *x2),
	}), lp.LpConstraintLE, 4)

	prog.AddConstraint(lp.NewExpression([]lp.LpTerm{
		lp.NewTerm(2, *x1),
		lp.NewTerm(1, *x2),
	}), lp.LpConstraintLE, 4)
}

func TestSolve_LP(t *testing.T) {
	variables := []lp.LpVariable{
		lp.NewVariable("x1"),
		lp.NewVariable("x2", lp.LpCategoryContinuous),
	}

	prog := lp.NewLinearProgram("Linear Problem", variables)
	constructTestProgram(&prog, variables)

	solver.Solve(&prog, solver.WithMaxIterations(3), solver.WithLogging(true))

	assert.AssertEqual(t, prog.Status, lp.LpStatusOptimal)
	assert.AssertIsClose(t, prog.Solution, 2.666666666, 1e-5)
	assert.AssertIsClose(t, prog.Variables.AtVec(0), 1.333333333, 1e-5) // x1
	assert.AssertIsClose(t, prog.Variables.AtVec(1), 1.333333333, 1e-5) // x2
	assert.AssertEqual(t, prog.VariablesMap[0].Category, lp.LpCategoryContinuous)
	assert.AssertEqual(t, prog.VariablesMap[1].Category, lp.LpCategoryContinuous)

}

func TestSolve_IP(t *testing.T) {
	variables := []lp.LpVariable{
		lp.NewVariable("x1", lp.LpCategoryInteger),
		lp.NewVariable("x2", lp.LpCategoryInteger),
	}

	prog := lp.NewLinearProgram("Integer Problem", variables)
	constructTestProgram(&prog, variables)

	solver.Solve(&prog, solver.WithMaxIterations(3), solver.WithLogging(true))

	assert.AssertEqual(t, prog.Status, lp.LpStatusOptimal)
	assert.AssertEqual(t, prog.Solution, 2.0)
	assert.AssertEqual(t, prog.Variables.AtVec(0), 0.0) // x1
	assert.AssertEqual(t, prog.Variables.AtVec(1), 2.0) // x2
	assert.AssertEqual(t, prog.VariablesMap[0].Category, lp.LpCategoryInteger)
	assert.AssertEqual(t, prog.VariablesMap[1].Category, lp.LpCategoryInteger)
}
