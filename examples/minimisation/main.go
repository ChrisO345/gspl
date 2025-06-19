package main

import (
	"github.com/chriso345/gspl"
)

func main() {
	// Objective: Minimise - 6 * x1 + 7 * x2 + 4 * x3
	// Constraints: 2 * x1 + 5 * x2 - 1 * x3 <= 18
	// Constraints: 1 * x1 - 1 * x2 - 2 * x3 <= -14
	// Constraints: 3 * x1 + 2 * x2 + 2 * x3 = 26

	variables := []gspl.LpVariable{
		gspl.NewVariable("x1"),
		gspl.NewVariable("x2"),
		gspl.NewVariable("x3"),
		gspl.NewVariable("x4"),
		gspl.NewVariable("x5"),
	}

	terms := []gspl.LpTerm{
		gspl.NewTerm(15, variables[0]),
		gspl.NewTerm(10, variables[1]),
		gspl.NewTerm(-10, variables[2]),
		gspl.NewTerm(1, variables[3]),
		gspl.NewTerm(2, variables[4]),
	}
	objective := gspl.NewExpression(terms)

	terms2 := []gspl.LpTerm{
		gspl.NewTerm(-1, variables[0]),
		gspl.NewTerm(-1, variables[1]),
		gspl.NewTerm(-1, variables[2]),
		gspl.NewTerm(-1, variables[3]),
		gspl.NewTerm(0, variables[4]),
	}

	terms3 := []gspl.LpTerm{
		gspl.NewTerm(0, variables[0]),
		gspl.NewTerm(1, variables[1]),
		gspl.NewTerm(0, variables[2]),
		gspl.NewTerm(1, variables[3]),
		gspl.NewTerm(-1, variables[4]),
	}

	terms4 := []gspl.LpTerm{
		gspl.NewTerm(-1, variables[0]),
		gspl.NewTerm(0, variables[1]),
		gspl.NewTerm(-1, variables[2]),
		gspl.NewTerm(0, variables[3]),
		gspl.NewTerm(-1, variables[4]),
	}

	lp := gspl.NewLinearProgram("Minimisation Example", variables)
	lp.AddObjective(gspl.LpMaximise, objective).
		AddConstraint(gspl.NewExpression(terms2), gspl.LpConstraintLE, -4).
		AddConstraint(gspl.NewExpression(terms3), gspl.LpConstraintLE, -4).
		AddConstraint(gspl.NewExpression(terms4), gspl.LpConstraintLE, -8)

	lp.Solve().PrintSolution()
}
