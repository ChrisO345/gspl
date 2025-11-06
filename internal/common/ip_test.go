package common

import (
	"testing"

	"github.com/chriso345/gore/assert"
	"gonum.org/v1/gonum/mat"
)

func TestIntegerProgramFields(t *testing.T) {
	scf := &StandardComputationalForm{
		Objective:      mat.NewVecDense(2, []float64{1, 2}),
		Constraints:    mat.NewDense(1, 2, []float64{3, 4}),
		RHS:            mat.NewVecDense(1, []float64{5}),
		PrimalSolution: mat.NewVecDense(2, []float64{6, 7}),
	}

	ip := &IntegerProgram{
		SCF:           scf,
		ConstraintDir: []string{"LE"},
		BestSolution:  mat.NewVecDense(2, []float64{0, 1}),
		BestObj:       10.0,
	}

	assert.NotNil(t, ip.SCF)
	assert.Equal(t, ip.BestObj, 10.0)
	assert.Equal(t, ip.BestSolution.AtVec(1), 1.0)
	assert.Equal(t, ip.ConstraintDir[0], "LE")
}
