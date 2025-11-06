package brancher

import (
	"github.com/chriso345/gspl/internal/common"
	"gonum.org/v1/gonum/mat"
)

// minimal SCF for testing
func newTestSCF(primal []float64) *common.StandardComputationalForm {
	objVal := 0.0
	status := common.SolverStatusOptimal
	vec := mat.NewVecDense(len(primal), primal)
	return &common.StandardComputationalForm{
		PrimalSolution: vec,
		ObjectiveValue: &objVal,
		Status:         &status,
		NumPrimals:     len(primal),
		SlackIndices:   make([]int, len(primal)),
	}
}
