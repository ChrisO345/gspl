package lp

import (
	"testing"

	"github.com/chriso345/gore/assert"
)

func TestNewTermVariableExpression(t *testing.T) {
	x := NewVariable("x")
	assert.Equal(t, x.Name, "x")
	assert.Equal(t, x.IsSlack, false)
	assert.Equal(t, x.IsArtificial, false)
	assert.Equal(t, x.Category, LpCategoryContinuous)

	xInt := NewVariable("y", LpCategoryInteger)
	assert.Equal(t, xInt.Category, LpCategoryInteger)

	term := NewTerm(5, x)
	assert.Equal(t, term.Coefficient, 5.0)
	assert.Equal(t, term.Variable.Name, "x")

	expr := NewExpression([]LpTerm{term})
	assert.Equal(t, len(expr.Terms), 1)
}
