package solver

import (
	"testing"

	"github.com/chriso345/gspl/lp"
)

func BenchmarkSolve_Small(b *testing.B) {
	for b.Loop() {
		vars := []lp.LpVariable{
			lp.NewVariable("x1"),
			lp.NewVariable("x2"),
			lp.NewVariable("x3"),
		}
		objective := lp.NewExpression([]lp.LpTerm{
			lp.NewTerm(1, vars[0]),
			lp.NewTerm(2, vars[1]),
			lp.NewTerm(3, vars[2]),
		})
		prog := lp.NewLinearProgram("bench", vars)
		prog.AddObjective(lp.LpMaximise, objective)
		prog.AddConstraint(lp.NewExpression([]lp.LpTerm{
			lp.NewTerm(1, vars[0]),
			lp.NewTerm(1, vars[1]),
		}), lp.LpConstraintLE, 10)
		if _, err := Solve(&prog); err != nil {
			b.Fatal(err)
		}
	}
}
