package solver_test

import (
	"testing"

	"github.com/chriso345/gore/assert"
	"github.com/chriso345/gspl/solver"
)

func TestDefaultConfig(t *testing.T) {
	cfg := solver.NewSolverConfig()

	assert.Equal(t, cfg.Tolerance, 1e-6)
	assert.Equal(t, cfg.MaxIterations, 1000)
	assert.Equal(t, cfg.SolverMethod, solver.SimplexMethod)
}

func TestPartialOverrideConfig(t *testing.T) {
	cfg := solver.NewSolverConfig(
		solver.WithTolerance(1e-5),
	)

	assert.Equal(t, cfg.Tolerance, 1e-5)
	assert.Equal(t, cfg.MaxIterations, 1000)
	assert.Equal(t, cfg.SolverMethod, solver.SimplexMethod)
}
