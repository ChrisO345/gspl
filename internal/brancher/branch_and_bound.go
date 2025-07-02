package brancher

import (
	"fmt"
	"math"
	"strings"

	"github.com/chriso345/gspl/internal/common"
	"github.com/chriso345/gspl/internal/matrix"
	"github.com/chriso345/gspl/internal/simplex"
	"gonum.org/v1/gonum/mat"
)

// BranchAndBound solves an integer linear programming problem using branch-and-bound.
func BranchAndBound(A *mat.Dense, b, c *mat.VecDense, m, n int, opts common.SolverConfig) (float64, *mat.VecDense, *mat.VecDense, int) {
	if opts.Logging {
		fmt.Println("Starting Branch-and-Bound")
	}

	root := &Branch{
		incumbentZ: math.Inf(-1),
	}

	branchRecurse(root, A, b, c, m, n, opts, 0)

	if root.incumbentX != nil {
		if opts.Logging {
			fmt.Println("Finished Branch-and-Bound: Found integer feasible solution.")
		}
		return root.incumbentZ, root.incumbentX, nil, 0
	}

	if opts.Logging {
		fmt.Println("Finished Branch-and-Bound: No integer feasible solution found.")
	}
	return 0, nil, nil, 1
}

// branchRecurse performs the recursive branching for the branch-and-bound algorithm.
func branchRecurse(root *Branch, A *mat.Dense, b, c *mat.VecDense, m, n int, opts common.SolverConfig, depth int) (float64, *mat.VecDense) {
	if depth > opts.MaxIterations {
		logf(opts, "%s!! Max depth reached at depth %d\n", strings.Repeat("  ", depth), depth)
		return root.incumbentZ, root.incumbentX
	}
	indent := strings.Repeat("  ", depth)
	logf(opts, "%s-> Solving LP relaxation at depth %d\n", indent, depth)

	// Solve the LP relaxation
	relaxZ, relaxX, relaxPi, relaxIndices, relaxFlag := simplex.Simplex(A, b, c, m, n)
	_, _ = relaxPi, relaxIndices // TODO: Unused variables for now, could these be used in a branching strategy?
	if relaxFlag != 0 {
		logf(opts, "%s!! Infeasible or unbounded at depth %d\n", indent, depth)
		return relaxZ, nil
	}

	logf(opts, "%s** LP relaxation z = %.4f, x = %v\n", indent, relaxZ, matrix.VecToSlice(relaxX))

	if isIntegerSolution(relaxX) {
		if relaxZ > root.incumbentZ {
			logf(opts, "%s>> Integer feasible solution found with z = %.4f (better than incumbent %.4f)\n", indent, relaxZ, root.incumbentZ)
			root.incumbentZ = relaxZ
			root.incumbentX = relaxX
		} else {
			logf(opts, "%s-- Integer feasible but not better than incumbent %.4f\n", indent, root.incumbentZ)
		}
		return relaxZ, relaxX
	}

	nextVar := getNextBranchingVariable(relaxX, opts.BranchingStrategy)
	if nextVar < 0 {
		logf(opts, "%s!! No fractional variables to branch on\n", indent)
		return relaxZ, relaxX
	}

	low, high := getBranchingValue(relaxX, nextVar)
	logf(opts, "%s>> Branching on variable x[%d] = %.4f → <= %d and >= %d\n", indent, nextVar, relaxX.AtVec(nextVar), low, high)

	leftConstraint := make([]float64, A.RawMatrix().Cols)
	leftConstraint[nextVar] = 1
	leftA := matrix.MatDenseAppendRow(A, mat.NewDense(1, len(leftConstraint), leftConstraint))
	leftB := matrix.VecDenseAppend(b, mat.NewVecDense(1, []float64{float64(low)}))
	logf(opts, "%s>> Going LEFT: x[%d] <= %d\n", indent, nextVar, low)
	root.left = &Branch{}
	branchRecurse(root, leftA, leftB, c, m+1, n, opts, depth+1)

	rightConstraint := make([]float64, A.RawMatrix().Cols)
	rightConstraint[nextVar] = -1
	rightA := matrix.MatDenseAppendRow(A, mat.NewDense(1, len(rightConstraint), rightConstraint))
	rightB := matrix.VecDenseAppend(b, mat.NewVecDense(1, []float64{-float64(high)}))
	logf(opts, "%s>> Going RIGHT: x[%d] >= %d\n", indent, nextVar, high)
	root.right = &Branch{}
	branchRecurse(root, rightA, rightB, c, m+1, n, opts, depth+1)

	return root.incumbentZ, root.incumbentX
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
func getNextBranchingVariable(x *mat.VecDense, method common.BranchingStrategy) int {
	if method != common.FirstFractional {
		panic(fmt.Sprintf("Branching strategy %s not implemented", method))
	}
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
