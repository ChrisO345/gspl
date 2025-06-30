package lp_test

import (
	"testing"

	"github.com/chriso345/gspl/lp"
	"gonum.org/v1/gonum/mat"
)

func makeSimpleLP() *lp.LinearProgram {
	return &lp.LinearProgram{
		VariablesMap: []string{"x", "y", "z"},
	}
}

func TestAddObjective_SetsObjectiveFunc(t *testing.T) {
	prog := makeSimpleLP()

	obj := lp.NewExpression([]lp.LpTerm{
		lp.NewTerm(1.0, lp.NewVariable("x")),
		lp.NewTerm(2.0, lp.NewVariable("y")),
		lp.NewTerm(3.0, lp.NewVariable("z")),
	})

	prog.AddObjective(lp.LpMinimise, obj)

	if prog.ObjectiveFunc.Len() != 3 {
		t.Fatalf("Expected ObjectiveFunc length 3, got %d", prog.ObjectiveFunc.Len())
	}

	expected := []float64{1, 2, 3}
	for i := range 3 {
		got := prog.ObjectiveFunc.AtVec(i)
		if got != expected[i] {
			t.Errorf("ObjectiveFunc at %d: got %f, want %f", i, got, expected[i])
		}
	}
}

func TestAddObjective_Maximise_NegatesObjective(t *testing.T) {
	prog := makeSimpleLP()

	obj := lp.NewExpression([]lp.LpTerm{
		lp.NewTerm(1.5, lp.NewVariable("x")),
		lp.NewTerm(-2.5, lp.NewVariable("y")),
		lp.NewTerm(0, lp.NewVariable("z")),
	})

	prog.AddObjective(lp.LpMaximise, obj)

	expected := []float64{-1.5, 2.5, 0}
	for i := range 3 {
		got := prog.ObjectiveFunc.AtVec(i)
		if got != expected[i] {
			t.Errorf("ObjectiveFunc at %d: got %f, want %f", i, got, expected[i])
		}
	}
}

func TestAddObjective_UnknownVariable_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for unknown variable, but did not panic")
		}
	}()

	prog := makeSimpleLP()

	obj := lp.NewExpression([]lp.LpTerm{
		lp.NewTerm(1, lp.NewVariable("unknown")),
	})

	prog.AddObjective(lp.LpMinimise, obj)
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

	if len(prog.ConstraintVector) != 1 {
		t.Fatalf("Expected ConstraintVector length 1, got %d", len(prog.ConstraintVector))
	}

	gotRow := mat.Row(nil, 0, prog.Constraints)
	expectedRow := []float64{1, 0, 2}
	for i, val := range expectedRow {
		if gotRow[i] != val {
			t.Errorf("Constraint matrix row[%d]: got %f, want %f", i, gotRow[i], val)
		}
	}

	gotRHS := prog.RHS.AtVec(0)
	if gotRHS != 10 {
		t.Errorf("RHS[0]: got %f, want 10", gotRHS)
	}
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
		if gotRow[i] != val {
			t.Errorf("Constraint matrix row[%d]: got %f, want %f", i, gotRow[i], val)
		}
	}

	gotRHS := prog.RHS.AtVec(0)
	if gotRHS != 5 {
		t.Errorf("RHS[0]: got %f, want 5", gotRHS)
	}

	if prog.ConstraintVector[0] != -lp.LpConstraintGE {
		t.Errorf("ConstraintVector[0]: got %v, want %v", prog.ConstraintVector[0], -lp.LpConstraintGE)
	}
}

func TestAddConstraint_WithoutObjective_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when adding constraint without objective set")
		}
	}()

	prog := makeSimpleLP()

	constraint := lp.NewExpression([]lp.LpTerm{
		lp.NewTerm(1, lp.NewVariable("x")),
	})

	prog.AddConstraint(constraint, lp.LpConstraintLE, 10)
}
