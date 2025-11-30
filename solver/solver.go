package solver

import (
	"fmt"

	"github.com/chriso345/gspl/internal/brancher"
	"github.com/chriso345/gspl/internal/common"
	"github.com/chriso345/gspl/internal/simplex"
	"github.com/chriso345/gspl/lp"
	"gonum.org/v1/gonum/mat"
)

// Solve solves the given linear program and returns the optimal objective value
func Solve(prog *lp.LinearProgram, opts ...SolverOption) error {
	// Apply options
	options := NewSolverConfig(opts...)

	tol := options.Tolerance

	if hasIPConstraints(prog) {

		ip := newIP(prog)

		// Call the Integer Programming solver
		err := brancher.BranchAndBound(ip, options)
		if err != nil {
			fmt.Println("Error during Solving:", err)
		}

		prog.ObjectiveValue = ip.BestObj
		prog.PrimalSolution = mat.NewVecDense(ip.SCF.NumPrimals, nil)
		for i := range ip.SCF.NumPrimals {
			item := ip.BestSolution.AtVec(i)
			if item < tol && item > -tol {
				continue
			}
			prog.PrimalSolution.SetVec(i, item)
		}

		return nil
	}

	// Create the SCF instance
	scf := newSCF(prog)

	// Call the Simplex solver
	if err := simplex.Simplex(scf, options); err != nil {
		return fmt.Errorf("simplex failed: %w", err)
	}

	// Remove any artificials and copy the solution back
	prog.ObjectiveValue = *scf.ObjectiveValue
	prog.PrimalSolution = mat.NewVecDense(scf.NumPrimals, nil)
	for i := range scf.NumPrimals {
		item := scf.PrimalSolution.AtVec(i)
		if item < tol && item > -tol {
			continue
		}
		prog.PrimalSolution.SetVec(i, item)
	}

	return nil
}

// newSCF creates a new SCF instance for the linear program
func newSCF(prog *lp.LinearProgram) *common.StandardComputationalForm {
	slackIndices := make([]int, len(prog.Vars))
	numPrimals := 0
	for i, constr := range prog.Vars {
		if constr.IsSlack {
			slackIndices[i] = i
		} else {
			slackIndices[i] = -1
			numPrimals++
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
		NumPrimals:     numPrimals,
	}
}

// newIP creates a new IP instance for the linear program
func newIP(prog *lp.LinearProgram) *common.IntegerProgram {
	return &common.IntegerProgram{
		SCF: newSCF(prog),
	}
}
