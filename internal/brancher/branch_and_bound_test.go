package brancher

import (
	"bytes"
	"io"
	"os"
	// "testing"
	//
	// "github.com/chriso345/gspl/internal/common"
	// "github.com/chriso345/gspl/internal/matrix"
	//
	// "github.com/chriso345/gore/assert"
)

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

// func TestBranchAndBound_Basic(t *testing.T) {
// 	A := matrix.MatDenseFromArray([][]float64{{2, 1}, {1, 2}})
// 	b := matrix.VecDenseFromArray([][]float64{{4}, {4}})
// 	c := matrix.VecDenseFromArray([][]float64{{1}, {1}})
// 	m, n := A.Dims()
// 	opts := common.DefaultSolverConfig()
//
// 	z, x, _, exitflag := BranchAndBound(A, b, c, m, n, *opts)
//
// 	assert.Equal(t, exitflag, 0)
// 	assert.IsClose(t, z, 2.0, 1e-4)
//
// 	if x != nil {
// 		expected := []float64{1, 1}
// 		result := matrix.VecToSlice(x)
// 		for i := range expected {
// 			assert.IsClose(t, result[i], expected[i], 1e-4)
// 		}
// 	}
// }

// func TestBranchAndBound_Infeasible(t *testing.T) {
// 	A := matrix.MatDenseFromArray([][]float64{{1, 1}, {1, 0}, {0, 1}})
// 	b := matrix.VecDenseFromArray([][]float64{{-1}, {1}, {1}})
// 	c := matrix.VecDenseFromArray([][]float64{{1}, {1}})
// 	m, n := A.Dims()
// 	opts := common.DefaultSolverConfig()
//
// 	z, x, _, exitflag := BranchAndBound(A, b, c, m, n, *opts)
//
// 	assert.Equal(t, exitflag, 1)
// 	assert.Equal(t, x, nil)
// 	assert.Equal(t, z, 0.0)
// }

// func TestBranchAndBound_IntegerAlready(t *testing.T) {
// 	A := matrix.MatDenseFromArray([][]float64{{1, 0}, {0, 1}})
// 	b := matrix.VecDenseFromArray([][]float64{{2}, {3}})
// 	c := matrix.VecDenseFromArray([][]float64{{1}, {1}})
// 	m, n := A.Dims()
// 	opts := common.DefaultSolverConfig()
//
// 	z, x, _, exitflag := BranchAndBound(A, b, c, m, n, *opts)
//
// 	assert.Equal(t, exitflag, 0)
// 	assert.IsClose(t, z, 5.0, 1e-4)
//
// 	expected := []float64{2, 3}
// 	result := matrix.VecToSlice(x)
// 	for i := range expected {
// 		assert.IsClose(t, result[i], expected[i], 1e-4)
// 	}
// }

// func TestBranchAndBound_LoggingEnabled(t *testing.T) {
// 	A := matrix.MatDenseFromArray([][]float64{{2, 1}, {1, 2}})
// 	b := matrix.VecDenseFromArray([][]float64{{4}, {4}})
// 	c := matrix.VecDenseFromArray([][]float64{{1}, {1}})
// 	m, n := A.Dims()
// 	opts := common.DefaultSolverConfig()
// 	opts.Logging = true
//
// 	logOutput := captureOutput(func() {
// 		BranchAndBound(A, b, c, m, n, *opts)
// 	})
//
// 	assert.True(t, strings.Contains(logOutput, "Starting Branch-and-Bound"))
// 	assert.True(t, strings.Contains(logOutput, "Branching on variable"))
// }

// func TestBranchAndBound_LoggingDisabled(t *testing.T) {
// 	A := matrix.MatDenseFromArray([][]float64{{2, 1}, {1, 2}})
// 	b := matrix.VecDenseFromArray([][]float64{{4}, {4}})
// 	c := matrix.VecDenseFromArray([][]float64{{1}, {1}})
// 	m, n := A.Dims()
// 	opts := common.DefaultSolverConfig()
// 	opts.Logging = false
//
// 	logOutput := captureOutput(func() {
// 		BranchAndBound(A, b, c, m, n, *opts)
// 	})
//
// 	assert.True(t, !strings.Contains(logOutput, "Starting Branch-and-Bound"))
// 	assert.True(t, !strings.Contains(logOutput, "Branching on variable"))
// }
