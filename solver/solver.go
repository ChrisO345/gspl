package solver

import (
	"fmt"

	"github.com/chriso345/gspl/internal/brancher"
	"github.com/chriso345/gspl/internal/common"
	"github.com/chriso345/gspl/internal/simplex"
	"github.com/chriso345/gspl/lp"
)

// Solve solves the given linear program and returns the optimal objective value
func Solve(prog *lp.LinearProgram, opts ...SolverOption) error {
	if hasIPConstraints(prog) {

		ip := newIP(prog)

		// Call the Integer Programming solver
		err := brancher.BranchAndBound(ip)
		if err != nil {
			fmt.Println("Error during Solving:", err)
		}

		if prog.Sense == lp.LpMaximise {
			ip.BestObj = -ip.BestObj
		}
		prog.ObjectiveValue = ip.BestObj

		fmt.Printf("[DEBUG] Solved IP: Status=%s, Objective=%.4f\n", prog.Status, ip.BestObj)
		fmt.Printf("[DEBUG] Primal Solution: %v\n", ip.BestSolution.RawVector().Data)

		return nil
	}

	// Create the SCF instance
	scf := newSCF(prog)

	// Call the Simplex solver
	err := simplex.Simplex(scf)
	if err != nil {
		fmt.Println("Error during Solving:", err)
	}

	fmt.Printf("[DEBUG] Solved LP: Status=%s, Objective=%.4f\n", prog.Status, prog.ObjectiveValue)

	return nil
}

// newSCF creates a new SCF instance for the linear program
func newSCF(prog *lp.LinearProgram) *common.StandardComputationalForm {
	// slackIndices := int{}
	slackIndices := make([]int, len(prog.Vars))
	for i, constr := range prog.Vars {
		if constr.IsSlack {
			slackIndices[i] = i
		} else {
			slackIndices[i] = -1
		}
	}

	return &common.StandardComputationalForm{
		Objective:   prog.Objective,
		Constraints: prog.Constraints,
		RHS:         prog.RHS,

		PrimalSolution: prog.PrimalSolution,

		// Link back to the original problem
		ObjectiveValue: &prog.ObjectiveValue,
		Status:         &prog.Status,
		SlackIndices:   slackIndices,
	}
}

// newIP creates a new IP instance for the linear program
func newIP(prog *lp.LinearProgram) *common.IntegerProgram {
	return &common.IntegerProgram{
		SCF: newSCF(prog),
	}
}
