package solver

import (
	"fmt"

	"github.com/chriso345/gspl/internal/brancher"
	"github.com/chriso345/gspl/internal/common"
	"github.com/chriso345/gspl/internal/matrix"
	"github.com/chriso345/gspl/internal/simplex"
	"github.com/chriso345/gspl/lp"
	"gonum.org/v1/gonum/mat"
)

// Solve takes a linear program and an optional configuration, and attempts to solve it using the revised simplex method.
func Solve(prog *lp.LinearProgram, opts ...SolverOption) *lp.LinearProgram {
	// Build the full solver options by applying defaults and options
	config := NewSolverConfig(opts...)

	// Validate config — implement validation method on solverConfig if needed
	if err := common.ValidateSolverConfig(config); err != nil {
		panic(fmt.Sprintf("Invalid solver configuration: %s", err.Error()))
	}

	prog.AddIPConstraints()
	if config.Logging {
		fmt.Printf("Linear Program: %v\n", prog)
	}

	// Add slacks for non-equality constraints
	for i, constraintType := range prog.ConstraintVector {
		if constraintType != lp.LpConstraintEQ {
			slack := lp.NewVariable(fmt.Sprintf("s%d", i))
			slack.IsSlack = true
			prog.VariablesMap = append(prog.VariablesMap, slack)
			unitVector := mat.NewDense(prog.Constraints.RawMatrix().Rows, 1, nil)
			one := 1.0
			if constraintType == lp.LpConstraintGE {
				one = -1.0
			}
			unitVector.Set(i, 0, one)
			prog.Constraints = matrix.ResizeMatDense(prog.Constraints, prog.Constraints.RawMatrix().Rows, len(prog.VariablesMap))

			prog.ObjectiveFunc = matrix.ResizeVecDense(prog.ObjectiveFunc, prog.ObjectiveFunc.RawVector().N+1)
			prog.ObjectiveFunc.SetVec(prog.ObjectiveFunc.RawVector().N-1, 0)
		}
	}

	if prog.Sense == lp.LpMinimise {
		for i := range prog.ObjectiveFunc.RawVector().N {
			prog.ObjectiveFunc.SetVec(i, -prog.ObjectiveFunc.At(i, 0))
		}
	}

	z, x, _, idx, flag := solveFormulation(prog, config)

	prog.Solution = z
	prog.Variables = x
	prog.Indices = idx

	switch flag {
	case 0:
		prog.Status = lp.LpStatusOptimal
	case 1:
		prog.Status = lp.LpStatusInfeasible
	case -1:
		prog.Status = lp.LpStatusUnbounded
	}

	if prog.Sense == lp.LpMinimise {
		prog.Solution *= -1
	}
	return prog
}

// solveFormulation solves the linear program via an appropriate method based on the presence of integer programming constraints.
func solveFormulation(prog *lp.LinearProgram, opts *common.SolverConfig) (float64, *mat.VecDense, *mat.VecDense, *mat.VecDense, simplex.ExitFlag) {
	m := prog.Constraints.RawMatrix().Rows
	n := len(prog.VariablesMap)

	if hasIPConstraints(prog) {
		if opts.Logging {
			fmt.Println("IP constraints detected, using branch-and-bound method")
		}

		z, x, idx, flag := brancher.BranchAndBound(prog.Constraints, prog.RHS, prog.ObjectiveFunc, m, n, brancher.BranchSense(prog.Sense), *opts)
		return z, x, nil, idx, flag
	}

	// If there are no IP constraints, we can solve the linear program simply.
	if opts.SolverMethod != SimplexMethod {
		panic(fmt.Sprintf("Solver method %s not implemented", opts.SolverMethod))
	}

	if opts.Logging {
		fmt.Println("No IP constraints detected, using simplex method")
	}
	return simplex.Simplex(prog.Constraints, prog.RHS, prog.ObjectiveFunc, m, n)
}
