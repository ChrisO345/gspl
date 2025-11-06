package brancher

import (
	"testing"

	"github.com/chriso345/gore/assert"
)

func TestIsIntegerFeasible(t *testing.T) {
	scf := newTestSCF([]float64{1, 2, 3})
	for i := range scf.SlackIndices {
		scf.SlackIndices[i] = -1
	}
	assert.True(t, isIntegerFeasible(scf))

	scf.PrimalSolution.SetVec(1, 2.5)
	assert.False(t, isIntegerFeasible(scf))
}
