package lp_test

import (
	"testing"

	"github.com/chriso345/gore/assert"
	"github.com/chriso345/gspl/lp"
	"gonum.org/v1/gonum/mat"
)

func makeSimpleLP() *lp.LinearProgram {
	return &lp.LinearProgram{
		VariablesMap: []lp.LpVariable{
			{"x", 0, false, false, lp.LpCategoryContinuous},
			{"y", 0, false, false, lp.LpCategoryContinuous},
			{"z", 0, false, false, lp.LpCategoryContinuous},
		},
	}
}

func TestAddObjective_SetsObjectiveFunc(t *testing.T) {
	prog := makeSimpleLP()

	obj := lp.NewExpression([]lp.LpTerm{
		lp.NewTerm(1.0, lp.NewVariable("x")),
		lp.NewTerm(2.0, lp.NewVariable("y")),
		lp.NewTerm(3.0, lp.NewVariable("z")),
	})

	prog.AddObjective(lp.LpMaximise, obj)

	assert.Equal(t, prog.ObjectiveFunc.Len(), 3)

	expected := []float64{1, 2, 3}
	for i := range 3 {
		got := prog.ObjectiveFunc.AtVec(i)
		assert.Equal(t, got, expected[i])
	}
}

func TestAddObjective_Minimise_NegatesObjective(t *testing.T) {
	prog := makeSimpleLP()

	obj := lp.NewExpression([]lp.LpTerm{
		lp.NewTerm(1.5, lp.NewVariable("x")),
		lp.NewTerm(-2.5, lp.NewVariable("y")),
		lp.NewTerm(0, lp.NewVariable("z")),
	})

	prog.AddObjective(lp.LpMinimise, obj)

	expected := []float64{-1.5, 2.5, 0}
	for i := range 3 {
		got := prog.ObjectiveFunc.AtVec(i)
		assert.Equal(t, got, expected[i])
	}
}

func TestAddObjective_UnknownVariable_Panics(t *testing.T) {
	prog := makeSimpleLP()

	obj := lp.NewExpression([]lp.LpTerm{
		lp.NewTerm(1, lp.NewVariable("unknown")),
	})

	assert.Panic(t, func() {
		prog.AddObjective(lp.LpMinimise, obj)
	})
}

func TestAddConstraint_AppendsConstraint(t *testing.T) {
	prog := makeSimpleLP()

	prog.AddObjective(lp.LpMinimise, lp.NewExpression([]lp.LpTerm{
		lp.NewTerm(1, lp.NewVariable("x")),
		lp.NewTerm(1, lp.NewVariable("y")),
		lp.NewTerm(1, lp.NewVariable("z")),
	}))

	constraint := lp.NewExpression([]lp.LpTerm{
		lp.NewTerm(1, lp.NewVariable("x")),
		lp.NewTerm(2, lp.NewVariable("z")),
	})

	prog.AddConstraint(constraint, lp.LpConstraintLE, 10)

	assert.Equal(t, len(prog.ConstraintVector), 1)

	gotRow := mat.Row(nil, 0, prog.Constraints)
	expectedRow := []float64{1, 0, 2}
	for i, val := range expectedRow {
		assert.Equal(t, gotRow[i], val)
	}

	assert.Equal(t, prog.RHS.AtVec(0), 10.0)
}

func TestAddConstraint_NegativeRHS_FlipsConstraint(t *testing.T) {
	prog := makeSimpleLP()

	prog.AddObjective(lp.LpMinimise, lp.NewExpression([]lp.LpTerm{
		lp.NewTerm(1, lp.NewVariable("x")),
		lp.NewTerm(1, lp.NewVariable("y")),
		lp.NewTerm(1, lp.NewVariable("z")),
	}))

	constraint := lp.NewExpression([]lp.LpTerm{
		lp.NewTerm(1, lp.NewVariable("x")),
		lp.NewTerm(-1, lp.NewVariable("y")),
	})

	prog.AddConstraint(constraint, lp.LpConstraintGE, -5)

	gotRow := mat.Row(nil, 0, prog.Constraints)
	expectedRow := []float64{-1, 1, 0}
	for i, val := range expectedRow {
		assert.Equal(t, val, gotRow[i])
	}

	assert.Equal(t, prog.RHS.AtVec(0), 5.0)

	assert.Equal(t, prog.ConstraintVector[0], -lp.LpConstraintGE)
}

func TestAddConstraint_WithoutObjective_Panics(t *testing.T) {
	prog := makeSimpleLP()

	constraint := lp.NewExpression([]lp.LpTerm{
		lp.NewTerm(1, lp.NewVariable("x")),
	})

	assert.Panic(t, func() {
		prog.AddConstraint(constraint, lp.LpConstraintLE, 10)
	})
}
