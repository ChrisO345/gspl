package lp

import (
	"testing"

	"github.com/chriso345/gore/assert"
)

func TestAddObjective(t *testing.T) {
	x1 := NewVariable("x1")
	x2 := NewVariable("x2")
	lp := NewLinearProgram("Test LP", []LpVariable{x1, x2})

	expr := NewExpression([]LpTerm{
		NewTerm(3, x1),
		NewTerm(5, x2),
	})

	lp.AddObjective(LpMinimise, expr)
	assert.Equal(t, lp.Objective.AtVec(0), 3.0)
	assert.Equal(t, lp.Objective.AtVec(1), 5.0)

	lpMax := NewLinearProgram("Max LP", []LpVariable{x1, x2})
	lpMax.AddObjective(LpMaximise, expr)
	assert.Equal(t, lpMax.Objective.AtVec(0), -3.0)
	assert.Equal(t, lpMax.Objective.AtVec(1), -5.0)
}

func TestAddObjectivePanic(t *testing.T) {
	lp := NewLinearProgram("Test LP", []LpVariable{})
	expr := NewExpression([]LpTerm{
		NewTerm(1, NewVariable("x1")),
	})

	assert.Panic(t, func() {
		lp.AddObjective(LpMinimise, expr)
	})
}

func TestAddConstraintWithoutObjectivePanic(t *testing.T) {
	x1 := NewVariable("x1")
	lp := NewLinearProgram("Test LP", []LpVariable{x1})
	expr := NewExpression([]LpTerm{
		NewTerm(1, x1),
	})
	assert.Panic(t, func() {
		lp.AddConstraint(expr, LpConstraintLE, 5)
	})
}

func TestAddConstraint(t *testing.T) {
	x1 := NewVariable("x1")
	x2 := NewVariable("x2")
	lp := NewLinearProgram("Test LP", []LpVariable{x1, x2})
	lp.AddObjective(LpMinimise, NewExpression([]LpTerm{
		NewTerm(1, x1),
		NewTerm(1, x2),
	}))

	expr := NewExpression([]LpTerm{
		NewTerm(1, x1),
		NewTerm(2, x2),
	})
	lp.AddConstraint(expr, LpConstraintLE, 4)

	assert.Equal(t, lp.Constraints.RawMatrix().Rows, 1)
	assert.Equal(t, len(lp.Vars), 3) // slack variable added

	// GE constraint with negative RHS flips
	expr2 := NewExpression([]LpTerm{
		NewTerm(2, x1),
		NewTerm(1, x2),
	})
	lp.AddConstraint(expr2, LpConstraintGE, -3)
	assert.Equal(t, lp.Constraints.RawMatrix().Rows, 2)
}
