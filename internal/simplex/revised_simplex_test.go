package simplex_test

import (
	"math"
	"testing"

	"github.com/chriso345/gore/assert"
	"github.com/chriso345/gspl/internal/simplex"
	"gonum.org/v1/gonum/mat"
)

const tolerance = 1e-8

func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < tolerance
}

func vectorEquals(vec mat.Vector, slice [][]float64) bool {
	if vec.Len() != len(slice) {
		return false
	}
	for i := range vec.Len() {
		if len(slice[i]) != 1 {
			return false
		}
		if !floatEquals(vec.AtVec(i), slice[i][0]) {
			return false
		}
	}
	return true
}

func flattenArray(arr [][]float64) []float64 {
	flat := make([]float64, 0, len(arr)*len(arr[0]))
	for _, row := range arr {
		flat = append(flat, row...)
	}
	return flat
}

func TestSimplexCases(t *testing.T) {
	tests := []struct {
		name         string
		A            [][]float64
		b            [][]float64
		c            [][]float64
		expectedZ    float64
		expectedX    [][]float64
		expectedIdx  [][]float64
		expectedExit simplex.ExitFlag
	}{
		{
			name: "Optimal BFS",
			A: [][]float64{
				{1, -1, 1, 1, 0, 0},
				{1, 1, 0, 0, -1, 0},
				{0, 0, 1, 0, 0, 1},
			},
			b: [][]float64{
				{4}, {0}, {6},
			},
			c: [][]float64{
				{1}, {2}, {3}, {0}, {0}, {0},
			},
			expectedZ: 0,
			expectedX: [][]float64{
				{0}, {0}, {0}, {4}, {0}, {6},
			},
			expectedIdx: [][]float64{
				{3}, {5}, {4},
			},
			expectedExit: simplex.ExitOptimal,
		},
		{
			name: "Already Optimal",
			A: [][]float64{
				{1, 0, 0, 1},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
			},
			b: [][]float64{
				{3}, {2}, {1},
			},
			c: [][]float64{
				{0}, {0}, {0}, {1},
			},
			expectedZ: 0,
			expectedX: [][]float64{
				{3}, {2}, {1}, {0},
			},
			expectedIdx:  nil, // Indexes not validated in this test
			expectedExit: simplex.ExitOptimal,
		},
		{
			name: "Unbounded",
			A: [][]float64{
				{0.08, 0.06, -1, 0, 0},
				{1, 0, 0, -1, 0},
				{0, 1, 0, 0, -1},
			},
			b:         [][]float64{{12}, {60}, {60}},
			c:         [][]float64{{-2}, {-1.25}, {0}, {0}, {0}},
			expectedZ: -285,
			expectedX: [][]float64{
				{105}, {60}, {0}, {45}, {0},
			},
			expectedIdx: [][]float64{
				{3}, {0}, {1},
			},
			expectedExit: simplex.ExitUnbounded,
		},
		{
			name: "Infeasible",
			A: [][]float64{
				{1, 0},
				{1, 0},
			},
			b:            [][]float64{{1}, {-2}},
			c:            [][]float64{{1}, {1}},
			expectedZ:    3,
			expectedX:    [][]float64{{-2}, {0}, {3}, {0}},
			expectedIdx:  nil,
			expectedExit: simplex.ExitInfeasible,
		},
		{
			name: "Medium Sized Problem",
			A: [][]float64{
				{1, 1, 1, 1, 0},
				{0, -1, 0, -1, 1},
				{1, 0, 1, 0, 1},
			},
			b:            [][]float64{{4}, {4}, {8}},
			c:            [][]float64{{15}, {10}, {-10}, {1}, {2}},
			expectedZ:    -32,
			expectedX:    [][]float64{{0}, {0}, {4}, {0}, {4}},
			expectedIdx:  [][]float64{{2}, {4}, {7}},
			expectedExit: simplex.ExitOptimal,
		},
		{
			name: "All Zeros",
			A: [][]float64{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			b: [][]float64{
				{0}, {0},
			},
			c: [][]float64{
				{0}, {0}, {0}, {0},
			},
			expectedZ: 0,
			expectedX: [][]float64{
				{0}, {0}, {0}, {0},
			},
			expectedIdx: [][]float64{
				{4}, {5},
			},
			expectedExit: simplex.ExitOptimal,
		},
		{
			name: "Big Problem (35 variables, 11 constaints",
			A: [][]float64{
				{1, 2, -1, 1, 1, 3, -7, 1, 1, -1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{-1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{2, -1, 1, -1, 2, 0, 0, 0, 0, 0, 1, -1, 1, 1, -1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 1, -1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{-1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, -1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 1, 1, 1, -1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, -1, 1, 0, 0},
				{1, -1, 0, 0, 0, 0, 0, 1, 1, -1, 1, 1, -1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -1, 1, 1, 1, -1, 1, 1, 1, 1, -1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, -1, 1, 1, 1, 1, 1, 1, -1, 1, 0, 0},
				{1, 1, -1, 1, 1, -1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, -1},
				{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			},
			b: [][]float64{
				{50}, {45}, {60}, {70}, {40}, {55}, {65}, {50}, {75}, {60}, {500}},
			c: [][]float64{
				{3}, {5}, {-2}, {7}, {-1}, {4}, {6}, {8}, {9}, {2}, {1}, {2}, {0}, {-4}, {5}, {3}, {-6}, {2}, {-1}, {4}, {5}, {-2}, {1}, {3}, {-1}, {2}, {1}, {3}, {-5}, {1}, {4}, {6}, {-3}, {2}, {1}},
			expectedZ: -130.5,
			expectedX: [][]float64{
				{2.5}, {0}, {0}, {0}, {8}, {0}, {0}, {39.5}, {0}, {0}, {13}, {0}, {0}, {26}, {0}, {0}, {15.5}, {0}, {46}, {0}, {0}, {14.5}, {0}, {0}, {0}, {0}, {28}, {0}, {47}, {0}, {0}, {0}, {0}, {10}, {0}},
			expectedIdx: [][]float64{
				{4}, {33}, {26}, {13}, {16}, {21}, {10}, {7}, {28}, {0}, {18},
			},
			expectedExit: simplex.ExitOptimal,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := len(test.b)
			n := len(test.c)
			A := mat.NewDense(m, n, flattenArray(test.A))
			b := mat.NewVecDense(m, flattenArray(test.b))
			c := mat.NewVecDense(n, flattenArray(test.c))

			z, x, _, finalIndices, exitflag := simplex.Simplex(A, b, c, m, n)

			assert.IsClose(t, z, test.expectedZ, tolerance)

			if !vectorEquals(x, test.expectedX) {
				t.Errorf("%s Incorrect: Expected x=%+v incorrect: got %+v", test.name, test.expectedX, x.RawVector().Data)
			}

			if test.expectedIdx != nil && !vectorEquals(finalIndices, test.expectedIdx) {
				t.Errorf("%s Incorrect: Expected indices=%+v got %+v", test.name, test.expectedIdx, finalIndices.RawVector().Data)
			}

			assert.Equal(t, exitflag, test.expectedExit)
		})
	}
}
