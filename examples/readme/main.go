package main

import "github.com/chriso345/gspl"

func main() {
	// Create decision variables
	variables := []gspl.LpVariable{
		gspl.NewVariable("x1"),
		gspl.NewVariable("x2"),
		gspl.NewVariable("x3"),
	}

	x1 := &variables[0]
	x2 := &variables[1]
	x3 := &variables[2]

	// Objective function: Minimize -6 * x1 + 7 * x2 + 4 * x3
	objective := gspl.NewExpression([]gspl.LpTerm{
		gspl.NewTerm(-6, *x1),
		gspl.NewTerm(7, *x2),
		gspl.NewTerm(4, *x3),
	})

	// Set up the LP problem
	lp := gspl.NewLinearProgram("README Example", variables)
	lp.AddObjective(gspl.LpMinimise, objective)

	// Add constraints
	lp.AddConstraint(gspl.NewExpression([]gspl.LpTerm{
		gspl.NewTerm(2, *x1),
		gspl.NewTerm(5, *x2),
		gspl.NewTerm(-1, *x3),
	}), gspl.LpConstraintLE, 18)

	lp.AddConstraint(gspl.NewExpression([]gspl.LpTerm{
		gspl.NewTerm(1, *x1),
		gspl.NewTerm(-1, *x2),
		gspl.NewTerm(-2, *x3),
	}), gspl.LpConstraintLE, -14)

	lp.AddConstraint(gspl.NewExpression([]gspl.LpTerm{
		gspl.NewTerm(3, *x1),
		gspl.NewTerm(2, *x2),
		gspl.NewTerm(2, *x3),
	}), gspl.LpConstraintEQ, 26)

	// Solve it
	lp.Solve().PrintSolution()
}
