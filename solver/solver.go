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

		return nil
	}

	// Create the SCF instance
	scf := newSCF(prog)

	// Call the Simplex solver
	err := simplex.Simplex(scf)
	if err != nil {
		fmt.Println("Error during Solving:", err)
	}

	return nil
}

// newSCF creates a new SCF instance for the linear program
func newSCF(prog *lp.LinearProgram) *common.StandardComputationalForm {
	return &common.StandardComputationalForm{
		Objective:   prog.Objective,
		Constraints: prog.Constraints,
		RHS:         prog.RHS,

		PrimalSolution: prog.PrimalSolution,

		// Link back to the original problem
		ObjectiveValue: &prog.ObjectiveValue,
		Status:         &prog.Status,
	}
}

// newIP creates a new IP instance for the linear program
func newIP(prog *lp.LinearProgram) *common.IntegerProgram {
	return &common.IntegerProgram{
		SCF: newSCF(prog),
	}
}
