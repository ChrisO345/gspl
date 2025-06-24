package main

import (
	"github.com/chriso345/gspl/solver"
)

func main() {
	// Create decision variables
	variables := []solver.LpVariable{
		solver.NewVariable("x1"),
		solver.NewVariable("x2"),
		solver.NewVariable("x3"),
	}

	x1 := &variables[0]
	x2 := &variables[1]
	x3 := &variables[2]

	// Objective function: Minimize -6 * x1 + 7 * x2 + 4 * x3
	objective := solver.NewExpression([]solver.LpTerm{
		solver.NewTerm(-6, *x1),
		solver.NewTerm(7, *x2),
		solver.NewTerm(4, *x3),
	})

	// Set up the LP problem
	lp := solver.NewLinearProgram("README Example", variables)
	lp.AddObjective(solver.LpMinimise, objective)

	// Add constraints
	lp.AddConstraint(solver.NewExpression([]solver.LpTerm{
		solver.NewTerm(2, *x1),
		solver.NewTerm(5, *x2),
		solver.NewTerm(-1, *x3),
	}), solver.LpConstraintLE, 18)

	lp.AddConstraint(solver.NewExpression([]solver.LpTerm{
		solver.NewTerm(1, *x1),
		solver.NewTerm(-1, *x2),
		solver.NewTerm(-2, *x3),
	}), solver.LpConstraintLE, -14)

	lp.AddConstraint(solver.NewExpression([]solver.LpTerm{
		solver.NewTerm(3, *x1),
		solver.NewTerm(2, *x2),
		solver.NewTerm(2, *x3),
	}), solver.LpConstraintEQ, 26)

	// Solve it
	lp.Solve().PrintSolution()
}
