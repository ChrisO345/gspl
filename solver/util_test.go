package solver

import (
	"testing"

	"github.com/chriso345/gore/assert"
	"github.com/chriso345/gspl/lp"
)

func TestHasIPConstraints(t *testing.T) {
	// Continuous variables only
	prog1 := &lp.LinearProgram{
		Vars: []lp.LpVariable{
			{Category: lp.LpCategoryContinuous},
			{Category: lp.LpCategoryContinuous},
		},
	}
	assert.False(t, hasIPConstraints(prog1))

	// Mixed variables
	prog2 := &lp.LinearProgram{
		Vars: []lp.LpVariable{
			{Category: lp.LpCategoryContinuous},
			{Category: lp.LpCategoryInteger},
		},
	}
	assert.True(t, hasIPConstraints(prog2))
}
