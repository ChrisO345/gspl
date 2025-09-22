package brancher

import (
	"fmt"
	"math"

	"github.com/chriso345/gspl/internal/common"
	"github.com/chriso345/gspl/internal/simplex"
	"gonum.org/v1/gonum/mat"
)

// BranchAndBound solves an integer linear programming problem using branch-and-bound.
func BranchAndBound(A *mat.Dense, b, c *mat.VecDense, m, n int, opts common.SolverConfig) (float64, *mat.VecDense, *mat.VecDense, simplex.ExitFlag) {
	if opts.Logging {
		fmt.Println("Starting Branch-and-Bound")
	}

	z, x, piValues, _, exitflag := simplex.Simplex(A, b, c, m, n)
	return z, x, piValues, exitflag
}

// isIntegerSolution checks if the solution vector x contains only integer values.
func isIntegerSolution(x *mat.VecDense) bool {
	const epsilon = 1e-5
	for i := range x.Len() {
		if math.Abs(x.AtVec(i)-math.Round(x.AtVec(i))) > epsilon {
			return false
		}
	}
	return true
}

// getNextBranchingVariable selects the next variable to branch on.
func getNextBranchingVariable(x *mat.VecDense) int {
	const epsilon = 1e-5
	for i := range x.Len() {
		if math.Abs(x.AtVec(i)-math.Round(x.AtVec(i))) > epsilon {
			return i
		}
	}
	return -1
}

// getBranchingValue returns the two values to split the branch on.
func getBranchingValue(x *mat.VecDense, index int) (int, int) {
	val := x.AtVec(index)
	return int(math.Floor(val)), int(math.Ceil(val))
}

// logf only prints when logging is enabled in the config
func logf(opts common.SolverConfig, format string, args ...any) {
	if opts.Logging {
		fmt.Printf(format, args...)
	}
}
