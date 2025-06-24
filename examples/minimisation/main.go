package main

import (
	"github.com/chriso345/gspl/solver"
)

func main() {
	// Objective: Minimise - 6 * x1 + 7 * x2 + 4 * x3
	// Constraints: 2 * x1 + 5 * x2 - 1 * x3 <= 18
	// Constraints: 1 * x1 - 1 * x2 - 2 * x3 <= -14
	// Constraints: 3 * x1 + 2 * x2 + 2 * x3 = 26

	variables := []solver.LpVariable{
		solver.NewVariable("x1"),
		solver.NewVariable("x2"),
		solver.NewVariable("x3"),
		solver.NewVariable("x4"),
		solver.NewVariable("x5"),
	}

	terms := []solver.LpTerm{
		solver.NewTerm(15, variables[0]),
		solver.NewTerm(10, variables[1]),
		solver.NewTerm(-10, variables[2]),
		solver.NewTerm(1, variables[3]),
		solver.NewTerm(2, variables[4]),
	}
	objective := solver.NewExpression(terms)

	terms2 := []solver.LpTerm{
		solver.NewTerm(-1, variables[0]),
		solver.NewTerm(-1, variables[1]),
		solver.NewTerm(-1, variables[2]),
		solver.NewTerm(-1, variables[3]),
		solver.NewTerm(0, variables[4]),
	}

	terms3 := []solver.LpTerm{
		solver.NewTerm(0, variables[0]),
		solver.NewTerm(1, variables[1]),
		solver.NewTerm(0, variables[2]),
		solver.NewTerm(1, variables[3]),
		solver.NewTerm(-1, variables[4]),
	}

	terms4 := []solver.LpTerm{
		solver.NewTerm(-1, variables[0]),
		solver.NewTerm(0, variables[1]),
		solver.NewTerm(-1, variables[2]),
		solver.NewTerm(0, variables[3]),
		solver.NewTerm(-1, variables[4]),
	}

	lp := solver.NewLinearProgram("Minimisation Example", variables)
	lp.AddObjective(solver.LpMaximise, objective).
		AddConstraint(solver.NewExpression(terms2), solver.LpConstraintLE, -4).
		AddConstraint(solver.NewExpression(terms3), solver.LpConstraintLE, -4).
		AddConstraint(solver.NewExpression(terms4), solver.LpConstraintLE, -8)

	lp.Solve().PrintSolution()
}
