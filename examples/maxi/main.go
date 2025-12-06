package main

import (
	"fmt"

	"github.com/chriso345/gspl/lp"
	"github.com/chriso345/gspl/solver"
)

func main() {
	// Continuous decision variables
	variables := []lp.LpVariable{
		lp.NewVariable("x1", lp.LpCategoryContinuous),
		lp.NewVariable("x2", lp.LpCategoryContinuous),
	}

	// Objective: Maximize 5*x1 + 4*x2
	objTerms := []lp.LpTerm{
		lp.NewTerm(5, variables[0]),
		lp.NewTerm(4, variables[1]),
	}
	objective := lp.NewExpression(objTerms)

	// Constraints:
	// 2*x1 + 3*x2 <= 12
	con1Terms := []lp.LpTerm{
		lp.NewTerm(2, variables[0]),
		lp.NewTerm(3, variables[1]),
	}
	// x1 + x2 <= 5
	con2Terms := []lp.LpTerm{
		lp.NewTerm(1, variables[0]),
		lp.NewTerm(1, variables[1]),
	}

	// Build LP maximization problem
	lpProg := lp.NewLinearProgram("Maximization Example", variables)
	lpProg.AddObjective(lp.LpMaximise, objective)
	lpProg.AddConstraint(lp.NewExpression(con1Terms), lp.LpConstraintLE, 12)
	lpProg.AddConstraint(lp.NewExpression(con2Terms), lp.LpConstraintLE, 5)

	fmt.Printf("%s\n", lpProg.String())

	// Solve it
	sol, err := solver.Solve(&lpProg)
	if err != nil {
		fmt.Println("solve error:", err)
		return
	}
	fmt.Printf("Optimal Objective Value: %.2f\n", sol.ObjectiveValue)
	fmt.Printf("Primal Solution: %v\n", sol.PrimalSolution.RawVector().Data)
}
