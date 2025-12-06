package main

import (
	"fmt"

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
		lp.NewTerm(1, variables[0]),
		lp.NewTerm(2, variables[1]),
		lp.NewTerm(3, variables[2]),
		lp.NewTerm(1, variables[3]),
		lp.NewTerm(4, variables[4]),
	}
	objective := lp.NewExpression(terms)

	terms2 := []lp.LpTerm{
		lp.NewTerm(1, variables[0]),
		lp.NewTerm(1, variables[1]),
		lp.NewTerm(1, variables[2]),
		lp.NewTerm(1, variables[3]),
		lp.NewTerm(1, variables[4]),
	}

	terms3 := []lp.LpTerm{
		lp.NewTerm(1, variables[0]),
		lp.NewTerm(2, variables[1]),
		lp.NewTerm(1, variables[2]),
		lp.NewTerm(0, variables[3]),
		lp.NewTerm(0, variables[4]),
	}

	terms4 := []lp.LpTerm{
		lp.NewTerm(0, variables[0]),
		lp.NewTerm(1, variables[1]),
		lp.NewTerm(0, variables[2]),
		lp.NewTerm(1, variables[3]),
		lp.NewTerm(1, variables[4]),
	}

	terms5 := []lp.LpTerm{
		lp.NewTerm(1, variables[0]),
		lp.NewTerm(0, variables[1]),
		lp.NewTerm(1, variables[2]),
		lp.NewTerm(0, variables[3]),
		lp.NewTerm(1, variables[4]),
	}

	terms6 := []lp.LpTerm{
		// lp.NewTerm(0, variables[0]),
		// lp.NewTerm(0, variables[1]),
		// lp.NewTerm(0, variables[2]),
		lp.NewTerm(1, variables[3]),
		lp.NewTerm(1, variables[4]),
	}

	minProg := lp.NewLinearProgram("Minimisation Example", variables)
	minProg.AddObjective(lp.LpMinimise, objective)
	minProg.AddConstraint(lp.NewExpression(terms2), lp.LpConstraintGE, 10)
	minProg.AddConstraint(lp.NewExpression(terms3), lp.LpConstraintLE, 8)
	minProg.AddConstraint(lp.NewExpression(terms4), lp.LpConstraintLE, 7)
	minProg.AddConstraint(lp.NewExpression(terms5), lp.LpConstraintGE, 4)
	minProg.AddConstraint(lp.NewExpression(terms6), lp.LpConstraintLE, 6)

	fmt.Printf("%s\n", minProg.String())
	sol, err := solver.Solve(&minProg)
	if err != nil {
		fmt.Println("solve error:", err)
		return
	}
	fmt.Printf("Optimal Objective Value: %.2f\n", sol.ObjectiveValue)
	fmt.Printf("Primal Solution: %v\n", sol.PrimalSolution.RawVector().Data)
}
