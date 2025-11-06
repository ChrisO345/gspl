package common

import (
	"testing"

	"github.com/chriso345/gore/assert"
)

func TestSolverStatusString(t *testing.T) {
	assert.Equal(t, SolverStatusNotSolved.String(), "Not Solved")
	assert.Equal(t, SolverStatusOptimal.String(), "Optimal")
	assert.Equal(t, SolverStatusInfeasible.String(), "Infeasible")
	assert.Equal(t, SolverStatusUnbounded.String(), "Unbounded")
	assert.Equal(t, SolverStatus(999).String(), "Unknown")
}
