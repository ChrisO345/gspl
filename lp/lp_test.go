package lp

import (
	"testing"

	"github.com/chriso345/gore/assert"
	"github.com/chriso345/gspl/internal/common"
)

func TestNewLinearProgram(t *testing.T) {
	vars := []LpVariable{
		NewVariable("x1"),
		NewVariable("x2", LpCategoryInteger),
	}
	lp := NewLinearProgram("Test LP", vars)

	assert.Equal(t, lp.Description, "Test LP")
	assert.Equal(t, len(lp.Vars), 2)
	assert.Equal(t, lp.Status, common.SolverStatusNotSolved)
}
