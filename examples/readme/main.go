package main

import (
	"fmt"

	"github.com/chriso345/gspl/lp"
	"github.com/chriso345/gspl/solver"
)

func main() {
	// Create decision variables
	variables := []lp.LpVariable{
		lp.NewVariable("x1"),
		lp.NewVariable("x2"),
		lp.NewVariable("x3"),
	}

	x1 := &variables[0]
	x2 := &variables[1]
	x3 := &variables[2]

	// Objective function: Minimize -6 * x1 + 7 * x2 + 4 * x3
	objective := lp.NewExpression([]lp.LpTerm{
		lp.NewTerm(-6, *x1),
		lp.NewTerm(7, *x2),
		lp.NewTerm(4, *x3),
	})

	// Set up the LP problem
	example := lp.NewLinearProgram("README Example", variables)
	example.AddObjective(lp.LpMinimise, objective)

	// Add constraints
	example.AddConstraint(lp.NewExpression([]lp.LpTerm{
		lp.NewTerm(2, *x1),
		lp.NewTerm(5, *x2),
		lp.NewTerm(-1, *x3),
	}), lp.LpConstraintLE, 18)

	example.AddConstraint(lp.NewExpression([]lp.LpTerm{
		lp.NewTerm(1, *x1),
		lp.NewTerm(-1, *x2),
		lp.NewTerm(-2, *x3),
	}), lp.LpConstraintLE, -14)

	example.AddConstraint(lp.NewExpression([]lp.LpTerm{
		lp.NewTerm(3, *x1),
		lp.NewTerm(2, *x2),
		lp.NewTerm(2, *x3),
	}), lp.LpConstraintEQ, 26)

	// Solve it
	// solver.Solve(&example, solver.WithLogging(true)).PrintSolution()

	fmt.Printf("%s\n", example.String())
	solver.Solve(&example)
}
