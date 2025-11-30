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

type enteringVariable struct {
	A  *mat.Dense    // Pointer to the simpleMethod.A
	pi *mat.VecDense // Pointer to the rsmResult.pi
	c  *mat.VecDense // Pointer to the simpleMethod.c

	isbasic *mat.VecDense

	epsilon float64

	// Results
	as *mat.VecDense
	cs float64
	s  int
}

type leavingVariable struct {
	B       *mat.Dense    // Pointer to the simpleMethod.B
	indices *mat.VecDense // Pointer to the rsmResult.indices
	as      *mat.VecDense
	xb      *mat.VecDense
	phase   int
	n       int

	r int
}

type basisUpdate struct {
	BMat    *mat.Dense
	indices *mat.VecDense
	cb      *mat.VecDense
	as      *mat.VecDense
	s       int
	r       int
	cs      float64
}
