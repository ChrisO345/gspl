package main

import (
	"github.com/chriso345/gspl/lp"
	"github.com/chriso345/gspl/solver"
)

func main() {
	// Define variables (all as Integer)
	variables := []lp.LpVariable{
		lp.NewVariable("x1", lp.LpCategoryInteger),
		lp.NewVariable("x2", lp.LpCategoryInteger),
	}

	// Objective: Minimize 3x1 + 4x2
	objectiveTerms := []lp.LpTerm{
		lp.NewTerm(3, variables[0]),
		lp.NewTerm(4, variables[1]),
	}
	objective := lp.NewExpression(objectiveTerms)

	// Constraint 1: 2x1 + x2 <= 10
	constraint1 := lp.NewExpression([]lp.LpTerm{
		lp.NewTerm(2, variables[0]),
		lp.NewTerm(1, variables[1]),
	})

	// Constraint 2: x1 + 2x2 >= 6
	constraint2 := lp.NewExpression([]lp.LpTerm{
		lp.NewTerm(1, variables[0]),
		lp.NewTerm(2, variables[1]),
	})

	// Build and solve the model
	prog := lp.NewLinearProgram("Feasible Pure IP Example", variables)
	prog.AddObjective(lp.LpMinimise, objective)
	prog.AddConstraint(constraint1, lp.LpConstraintLE, 10)
	prog.AddConstraint(constraint2, lp.LpConstraintGE, 6)

	solver.Solve(&prog)
	prog.PrintSolution()
}
