package solver

import (
	"github.com/chriso345/gspl/lp"
)

func Solve(prog *lp.LinearProgram) (float64, []float64, error) {
	if hasIPConstraints(prog) {
		panic("IP solving not yet implemented")
	}
	panic("LP solving not yet implemented")
}
