package solver_test

import (
	"testing"

	"github.com/chriso345/gspl/solver"
)

func TestDefaultConfig(t *testing.T) {
	cfg := solver.NewSolverConfig()

	if cfg.MaxIterations != 1000 {
		t.Errorf("Expected MaxIterations to be 1000, got %d", cfg.MaxIterations)
	}
	if cfg.Tolerance != 1e-6 {
		t.Errorf("Expected Tolerance to be 1e-6, got %.1e", cfg.Tolerance)
	}
	if cfg.SolverMethod != solver.SimplexMethod {
		t.Errorf("Expected SolverMethod to be Simplex, got %v", cfg.SolverMethod)
	}
}

func TestPartialOverrideConfig(t *testing.T) {
	cfg := solver.NewSolverConfig(
		solver.WithTolerance(1e-5),
	)

	if cfg.MaxIterations != 1000 {
		t.Errorf("Expected MaxIterations to be 1000, got %d", cfg.MaxIterations)
	}
	if cfg.Tolerance != 1e-5 {
		t.Errorf("Expected Tolerance to be 1e-5, got %.1e", cfg.Tolerance)
	}
	if cfg.SolverMethod != solver.SimplexMethod {
		t.Errorf("Expected SolverMethod to be Simplex, got %v", cfg.SolverMethod)
	}
}
