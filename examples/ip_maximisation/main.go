package main

import (
	"fmt"

	"github.com/chriso345/gspl/lp"
	"github.com/chriso345/gspl/solver"
)

func main() {
	variables := []lp.LpVariable{
		lp.NewVariable("x1", lp.LpCategoryInteger),
		lp.NewVariable("x2", lp.LpCategoryInteger),
	}

	terms := []lp.LpTerm{
		lp.NewTerm(3, variables[0]),
		lp.NewTerm(2, variables[1]),
	}
	objective := lp.NewExpression(terms)

	terms2 := []lp.LpTerm{
		lp.NewTerm(1, variables[0]),
		lp.NewTerm(1, variables[1]),
	}
	terms3 := []lp.LpTerm{
		lp.NewTerm(0.67, variables[0]),
		lp.NewTerm(0, variables[1]),
	}

	// This is a naturally integer problem... FIXME: Change to non-natural integer problem
	minProg := lp.NewLinearProgram("Minimisation Example", variables)
	minProg.AddObjective(lp.LpMaximise, objective)
	minProg.AddConstraint(lp.NewExpression(terms2), lp.LpConstraintLE, 4)
	minProg.AddConstraint(lp.NewExpression(terms3), lp.LpConstraintLE, 2)

	fmt.Printf("%s\n", minProg.String())
	solver.Solve(&minProg)
}
