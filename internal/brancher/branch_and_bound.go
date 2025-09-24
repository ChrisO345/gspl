package brancher

import (
	"fmt"
	"math"

	"github.com/chriso345/gspl/internal/common"
	"github.com/chriso345/gspl/internal/matrix"
	"github.com/chriso345/gspl/internal/simplex"
	"gonum.org/v1/gonum/mat"
)

// BranchAndBound solves an integer linear programming problem using branch-and-bound.
func BranchAndBound(A *mat.Dense, b, c *mat.VecDense, m, n int, sense BranchSense, opts common.SolverConfig) (float64, *mat.VecDense, *mat.VecDense, simplex.ExitFlag) {
	if opts.Logging {
		fmt.Println("Starting Branch-and-Bound")
	}

	z, x, piValues, _, exitflag := simplex.Simplex(A, b, c, m, n)
	logf(opts, "Initial LP relaxation: z=%.4f, x=%v.T\n", z, x.RawVector().Data)

	if exitflag != simplex.ExitOptimal {
		logf(opts, "LP relaxation not optimal, exit flag: %d\n", exitflag)
		return math.Inf(1), nil, nil, exitflag
	}

	if isIntegerSolution(x) {
		logf(opts, "Found integer solution at root: z=%.4f, x=%v.T\n", z, x.RawVector().Data)
		return z, x, piValues, exitflag
	}

	bestZ := math.Inf(-1)
	var bestX *mat.VecDense
	bestZ, bestX, exitflag = branch(A, b, c, m, n, sense, opts, bestZ, bestX)
	if bestX != nil {
		logf(opts, "Best integer solution found: z=%.4f, x=%v.T\n", bestZ, bestX.RawVector().Data)
		if sense == BranchSenseMinimize {
			bestZ = -bestZ
		}
	} else {
		logf(opts, "No integer solution found.\n")
		bestZ = math.Inf(1)
	}
	return bestZ, bestX, piValues, exitflag
}

func branch(A *mat.Dense, b, c *mat.VecDense, m, n int, sense BranchSense, opts common.SolverConfig, bestZ float64, bestX *mat.VecDense) (float64, *mat.VecDense, simplex.ExitFlag) {
	z, x, _, _, exitflag := simplex.Simplex(A, b, c, m, n)
	logf(opts, "Solving LP at depth %d: z=%.4f, x=%v.T, exitflag=%d\n", m, z, x.RawVector().Data, exitflag)
	if exitflag == simplex.ExitUnbounded || exitflag == simplex.ExitInfeasible {
		return bestZ, bestX, exitflag
	}
	if isIntegerSolution(x) {
		if z > bestZ {
			logf(opts, "Found new best integer solution: z=%.4f, x=%v.T\n", z, x.RawVector().Data)
			return z, x, exitflag
		}
		return bestZ, bestX, exitflag
	}

	index := getNextBranchingVariable(x)
	if index == -1 {
		return bestZ, bestX, exitflag
	}
	lowerVal, upperVal := getBranchingValue(x, index)

	// Branch on the lower bound
	A1 := mat.DenseCopyOf(A)
	b1 := mat.VecDenseCopyOf(b)
	newRow1 := make([]float64, n)
	newRow1[index] = 1
	A1 = matrix.MatDenseStack(A1, mat.NewDense(1, n, newRow1))
	b1 = matrix.VecDenseStack(b1, mat.NewVecDense(1, []float64{float64(lowerVal)}))
	logf(opts, "Branching on x[%d] >= %d\n", index, lowerVal)
	lowerZ, lowerX, lowerflag := branch(A1, b1, c, m+1, n, sense, opts, bestZ, bestX)

	// Branch on the upper bound
	A2 := mat.DenseCopyOf(A)
	b2 := mat.VecDenseCopyOf(b)
	newRow2 := make([]float64, n)
	newRow2[index] = -1
	A2 = matrix.MatDenseStack(A2, mat.NewDense(1, n, newRow2))
	b2 = matrix.VecDenseStack(b2, mat.NewVecDense(1, []float64{-float64(upperVal)}))
	logf(opts, "Branching on x[%d] <= %d\n", index, upperVal)
	upperZ, upperX, upperflag := branch(A2, b2, c, m+1, n, sense, opts, bestZ, bestX)

	logf(opts, "Depth %d results: lowerZ=%.4f (flag=%d), upperZ=%.4f (flag=%d), bestZ=%.4f\n", m, lowerZ, lowerflag, upperZ, upperflag, bestZ)

	if sense == BranchSenseMinimize {
		lowerZ = -lowerZ
		upperZ = -upperZ
	}

	if lowerflag == simplex.ExitOptimal && lowerZ > bestZ {
		bestZ = lowerZ
		bestX = lowerX
	}

	if upperflag == simplex.ExitOptimal && upperZ > bestZ {
		bestZ = upperZ
		bestX = upperX
	}

	if lowerflag != simplex.ExitOptimal && upperflag != simplex.ExitOptimal {
		exitflag = simplex.ExitInfeasible
	}

	return bestZ, bestX, exitflag
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
