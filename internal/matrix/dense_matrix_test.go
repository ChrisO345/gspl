package matrix

import (
	"testing"

	"github.com/chriso345/gore/assert"
	"gonum.org/v1/gonum/mat"
)

func TestEye(t *testing.T) {
	m := Eye(3)
	r, c := m.Dims()
	assert.Equal(t, r, 3)
	assert.Equal(t, c, 3)
	for i := range 3 {
		for j := range 3 {
			if i == j {
				assert.Equal(t, m.At(i, j), 1.0)
			} else {
				assert.Equal(t, m.At(i, j), 0.0)
			}
		}
	}
}

func TestResizeMatDense(t *testing.T) {
	orig := mat.NewDense(2, 2, []float64{1, 2, 3, 4})
	resized := ResizeMatDense(orig, 3, 3)

	r, c := resized.Dims()
	assert.Equal(t, r, 3)
	assert.Equal(t, c, 3)
	assert.Equal(t, resized.At(0, 0), 1.0)
	assert.Equal(t, resized.At(1, 1), 4.0)
	assert.Equal(t, resized.At(2, 2), 0.0)
}

func TestExtractColumns(t *testing.T) {
	A := mat.NewDense(2, 3, []float64{1, 2, 3, 4, 5, 6})
	indices := mat.NewVecDense(2, []float64{0, 2})
	extracted := ExtractColumns(A, indices)

	r, c := extracted.Dims()
	assert.Equal(t, r, 2)
	assert.Equal(t, c, 2)
	assert.Equal(t, extracted.At(0, 0), 1.0)
	assert.Equal(t, extracted.At(0, 1), 3.0)
	assert.Equal(t, extracted.At(1, 1), 6.0)
}

func TestMatDenseFromArrayAndToArray(t *testing.T) {
	arr := [][]float64{{1, 2}, {3, 4}}
	matDense := MatDenseFromArray(arr)
	out := MatDenseToArray(matDense)
	assert.Equal(t, out[0][0], 1.0)
	assert.Equal(t, out[1][1], 4.0)
}

func TestMatDenseAppendColumn(t *testing.T) {
	A := mat.NewDense(2, 2, []float64{1, 2, 3, 4})
	col := mat.NewVecDense(2, []float64{5, 6})
	result := MatDenseAppendColumn(A, col)
	assert.Equal(t, result.At(0, 2), 5.0)
	assert.Equal(t, result.At(1, 2), 6.0)
}

func TestMatDenseAppendRow(t *testing.T) {
	A := mat.NewDense(2, 2, []float64{1, 2, 3, 4})
	row := mat.NewDense(1, 2, []float64{5, 6})
	result := MatDenseAppendRow(A, row)

	r, c := result.Dims()
	assert.Equal(t, r, 3)
	assert.Equal(t, c, 2)
	assert.Equal(t, result.At(2, 0), 5.0)
	assert.Equal(t, result.At(2, 1), 6.0)
}

func TestMatDenseStack(t *testing.T) {
	A := mat.NewDense(2, 2, []float64{1, 2, 3, 4})
	B := mat.NewDense(1, 2, []float64{5, 6})
	result := MatDenseStack(A, B)

	r, c := result.Dims()
	assert.Equal(t, r, 3)
	assert.Equal(t, c, 2)
	assert.Equal(t, result.At(2, 0), 5.0)
	assert.Equal(t, result.At(2, 1), 6.0)
}
