package solver

import (
	"fmt"

	"github.com/chriso345/gspl/internal/matrix"
	"github.com/chriso345/gspl/internal/simplex"
	"github.com/chriso345/gspl/lp"
	"gonum.org/v1/gonum/mat"
)

func Solve(prog *lp.LinearProgram, config ...SolverOption) *lp.LinearProgram {
	var opts SolverOption
	if len(config) == 0 {
		opts = DefaultSolverOption()
	} else {
		opts = config[0]
		opts.fillDefaults()
	}

	// Validate the solver options
	err := opts.validateSolverOption()
	if err != nil {
		panic(fmt.Sprintf("Invalid Linear Program: %s", err.Error()))
	}

	// Add slacks for non-equality constraints
	for i, constraintType := range prog.ConstraintVector {
		if constraintType != lp.LpConstraintEQ {
			slack := lp.NewVariable(fmt.Sprintf("s%d", i))
			prog.VariablesMap = append(prog.VariablesMap, slack.Name)
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

	m := prog.Constraints.RawMatrix().Rows
	n := len(prog.VariablesMap)

	z, x, _, idx, flag := simplex.Simplex(prog.Constraints, prog.RHS, prog.ObjectiveFunc, m, n)

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

	if prog.Sense == lp.LpMaximise {
		prog.Solution *= -1
	}

	return prog
}
