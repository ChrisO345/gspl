package main

import (
	"github.com/chriso345/gspl/lp"
	"github.com/chriso345/gspl/solver"
)

func main() {
	// Objective: Minimise - 6 * x1 + 7 * x2 + 4 * x3
	// Constraints: 2 * x1 + 5 * x2 - 1 * x3 <= 18
	// Constraints: 1 * x1 - 1 * x2 - 2 * x3 <= -14
	// Constraints: 3 * x1 + 2 * x2 + 2 * x3 = 26

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
	solver.AddObjective(&minProg, lp.LpMaximise, objective)
	solver.AddConstraint(&minProg, lp.NewExpression(terms2), lp.LpConstraintLE, -4)
	solver.AddConstraint(&minProg, lp.NewExpression(terms3), lp.LpConstraintLE, -4)
	solver.AddConstraint(&minProg, lp.NewExpression(terms4), lp.LpConstraintLE, -8)

	solver.Solve(&minProg)
	minProg.PrintSolution()
}
