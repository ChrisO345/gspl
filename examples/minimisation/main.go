package main

import (
	"github.com/chriso345/gspl/lp"
	"github.com/chriso345/gspl/solver"
)

func main() {
	variables := []lp.LpVariable{
		lp.NewVariable("x1"),
		lp.NewVariable("x2"),
		lp.NewVariable("x3"),
		lp.NewVariable("x4"),
		lp.NewVariable("x5"),
	}

	terms := []lp.LpTerm{
		lp.NewTerm(15, variables[0]),
		lp.NewTerm(10, variables[1]),
		lp.NewTerm(-10, variables[2]),
		lp.NewTerm(1, variables[3]),
		lp.NewTerm(2, variables[4]),
	}
	objective := lp.NewExpression(terms)

	terms2 := []lp.LpTerm{
		lp.NewTerm(-1, variables[0]),
		lp.NewTerm(-1, variables[1]),
		lp.NewTerm(-1, variables[2]),
		lp.NewTerm(-1, variables[3]),
		lp.NewTerm(0, variables[4]),
	}

	terms3 := []lp.LpTerm{
		lp.NewTerm(0, variables[0]),
		lp.NewTerm(1, variables[1]),
		lp.NewTerm(0, variables[2]),
		lp.NewTerm(1, variables[3]),
		lp.NewTerm(-1, variables[4]),
	}

	terms4 := []lp.LpTerm{
		lp.NewTerm(-1, variables[0]),
		lp.NewTerm(0, variables[1]),
		lp.NewTerm(-1, variables[2]),
		lp.NewTerm(0, variables[3]),
		lp.NewTerm(-1, variables[4]),
	}

	minProg := lp.NewLinearProgram("Minimisation Example", variables)
	minProg.AddObjective(lp.LpMinimise, objective)
	minProg.AddConstraint(lp.NewExpression(terms2), lp.LpConstraintLE, -4)
	minProg.AddConstraint(lp.NewExpression(terms3), lp.LpConstraintLE, -4)
	minProg.AddConstraint(lp.NewExpression(terms4), lp.LpConstraintLE, -8)

	solver.Solve(&minProg, solver.WithLogging(true))
	minProg.PrintSolution()
}
