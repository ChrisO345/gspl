package solver

import "testing"

func Test_fillDefaults(t *testing.T) {
	opts := SolverOption{}
	opts.fillDefaults()

	if *opts.MaxIterations != 1000 {
		t.Errorf("Expected MaxIterations to be 1000, got %d", *opts.MaxIterations)
	}
	if *opts.Tolerance != 1e-6 {
		t.Errorf("Expected Tolerance to be 1e-6, got %.1e", *opts.Tolerance)
	}
	if *opts.SolverMethod != SimplexMethod {
		t.Errorf("Expected SolverType to be Simplex, got %v", *opts.SolverMethod)
	}
}

func Test_partialDefaults(t *testing.T) {
	opts := SolverOption{
		Tolerance: ptr(1e-5),
	}
	opts.fillDefaults()

	if *opts.MaxIterations != 1000 {
		t.Errorf("Expected MaxIterations to be 1000, got %d", *opts.MaxIterations)
	}
	if *opts.Tolerance != 1e-5 {
		t.Errorf("Expected Tolerance to be 1e-5, got %.1e", *opts.Tolerance)
	}
	if *opts.SolverMethod != SimplexMethod {
		t.Errorf("Expected SolverType to be Simplex, got %v", *opts.SolverMethod)
	}
}
