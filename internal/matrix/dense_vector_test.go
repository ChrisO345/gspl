package matrix

import (
	"testing"

	"github.com/chriso345/gore/assert"
	"gonum.org/v1/gonum/mat"
)

func TestVecDenseFromArrayAndVecToSlice(t *testing.T) {
	arr := [][]float64{{1}, {2}, {3}}
	vec := VecDenseFromArray(arr)
	slice := VecToSlice(vec)
	assert.Equal(t, slice[0], 1.0)
	assert.Equal(t, slice[2], 3.0)
}

func TestResizeVecDense(t *testing.T) {
	vec := mat.NewVecDense(2, []float64{1, 2})
	resized := ResizeVecDense(vec, 3)
	assert.Equal(t, resized.Len(), 3)
	assert.Equal(t, resized.AtVec(0), 1.0)
	assert.Equal(t, resized.AtVec(2), 0.0)
}

func TestVecDenseAppend(t *testing.T) {
	v1 := mat.NewVecDense(2, []float64{1, 2})
	v2 := mat.NewVecDense(2, []float64{3, 4})
	result := VecDenseAppend(v1, v2)
	assert.Equal(t, result.Len(), 4)
	assert.Equal(t, result.AtVec(2), 3.0)
}

func TestVecDenseStack(t *testing.T) {
	v1 := mat.NewVecDense(2, []float64{1, 2})
	v2 := mat.NewVecDense(2, []float64{3, 4})
	stacked := VecDenseStack(v1, v2)
	assert.Equal(t, stacked.Len(), 4)
	assert.Equal(t, stacked.AtVec(3), 4.0)
}
