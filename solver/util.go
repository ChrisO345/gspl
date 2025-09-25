package solver

import "github.com/chriso345/gspl/lp"

func hasIPConstraints(prog *lp.LinearProgram) bool {
	for _, v := range prog.Vars {
		if v.Category != lp.LpCategoryContinuous {
			return true
		}
	}
	return false
}
