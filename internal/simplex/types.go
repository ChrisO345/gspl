package simplex

import (
	"github.com/chriso345/gspl/internal/common"
	"gonum.org/v1/gonum/mat"
)

type simplexMethod struct {
	A *mat.Dense
	b *mat.VecDense
	c *mat.VecDense

	m int
	n int

	B  *mat.Dense
	cb *mat.VecDense

	rsmResult
}

type rsmResult struct {
	value   float64
	x       *mat.VecDense
	pi      *mat.VecDense
	indices *mat.VecDense
	flag    common.SolverStatus
}
