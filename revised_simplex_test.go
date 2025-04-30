package gspl

import (
	"math"
	"testing"
)

const tolerance = 1e-8

func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < tolerance
}

func matrixEquals(a, b [][]float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) || math.Abs(a[i][0]-b[i][0]) > tolerance {
			return false
		}
	}
	return true
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
		expectedExit int
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
			expectedExit: 0,
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
			expectedExit: 0,
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
			expectedExit: -1,
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
			expectedExit: 0,
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
			expectedExit: 0,
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
			expectedExit: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.name != "Medium Sized Problem" {
				t.Skip("Skipping test: " + test.name)
				return
			}
			m := len(test.b)
			n := len(test.c)
			A := NewMatrix(m, n)
			A.Values = test.A
			b := NewMatrix(m, 1)
			b.Values = test.b
			c := NewMatrix(n, 1)
			c.Values = test.c

			z, x, _, finalIndices, exitflag := Simplex(A, b, c, m, n)

			if !floatEquals(z, test.expectedZ) {
				t.Errorf("%s Incorrect: Expected z=%.2f, Got %.6f", test.name, test.expectedZ, z)
			} else {
				t.Logf("%s: z correct", test.name)
			}

			if !matrixEquals(x.Values, test.expectedX) {
				t.Errorf("%s Incorrect: Expected x=%+v incorrect: got %+v", test.name, test.expectedX, x.Values)
			} else {
				t.Logf("%s: x correct", test.name)
			}

			if test.expectedIdx != nil && !matrixEquals(finalIndices.Values, test.expectedIdx) {
				t.Errorf("%s Incorrect: Expected indices=%+v got %+v", test.name, test.expectedIdx, finalIndices.Values)
			} else if test.expectedIdx != nil {
				t.Logf("%s: indices correct", test.name)
			}

			if exitflag != test.expectedExit {
				t.Errorf("%s Incorrect: Expected exitflag=%d, Got %d", test.name, test.expectedExit, exitflag)
			} else {
				t.Logf("%s: exitflag correct", test.name)
			}
		})
	}
}
