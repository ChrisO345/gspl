package matrix

import (
	"math"
	"testing"
)

// Epsilon tolerance function
func epsilon(a, b float64) bool {
	return math.Abs(a-b) <= 1e-9
}

func TestNewMatrix(t *testing.T) {
	m := NewMatrix(3, 3)

	if m.Rows != 3 || m.Columns != 3 {
		t.Errorf("Expected matrix of size 3x3, but got %dx%d", m.Rows, m.Columns)
	}

	// Check that all initial values are 0
	for i := range m.Rows {
		for j := range m.Columns {
			if !epsilon(m.Values[i][j], 0) {
				t.Errorf("Expected 0 at position (%d, %d), but got %f", i, j, m.Values[i][j])
			}
		}
	}
}

func TestFromArray(t *testing.T) {
	matrix := [][]int{
		{1, 2},
		{3, 4},
	}
	m := FromArray(2, 2, matrix)

	expectedValues := [][]float64{
		{1.0, 2.0},
		{3.0, 4.0},
	}

	// Check if values match
	for i := range m.Rows {
		for j := range m.Columns {
			if !epsilon(m.Values[i][j], expectedValues[i][j]) {
				t.Errorf("Expected value %f at (%d, %d), but got %f", expectedValues[i][j], i, j, m.Values[i][j])
			}
		}
	}
}

func TestSize(t *testing.T) {
	m := NewMatrix(4, 5)
	rows, cols := m.Size()

	if rows != 4 || cols != 5 {
		t.Errorf("Expected size 4x5, but got %dx%d", rows, cols)
	}
}

func TestLength(t *testing.T) {
	m := NewMatrix(3, 4)
	if m.Length() != 4 {
		t.Errorf("Expected length 4, but got %d", m.Length())
	}
}

func TestGetSet(t *testing.T) {
	m := NewMatrix(2, 2)
	m.Set(0, 0, 5.0)
	m.Set(1, 1, 10.0)

	if !epsilon(m.Get(0, 0), 5.0) {
		t.Errorf("Expected value 5.0 at (0, 0), but got %f", m.Get(0, 0))
	}

	if !epsilon(m.Get(1, 1), 10.0) {
		t.Errorf("Expected value 10.0 at (1, 1), but got %f", m.Get(1, 1))
	}
}

func TestSetOutOfBounds(t *testing.T) {
	m := NewMatrix(2, 2)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when setting out-of-bounds value")
		}
	}()
	m.Set(3, 3, 5.0) // This should panic
}

func TestAdd(t *testing.T) {
	m1 := NewMatrix(2, 2)
	m1.Set(0, 0, 1.0)
	m1.Set(0, 1, 2.0)
	m1.Set(1, 0, 3.0)
	m1.Set(1, 1, 4.0)

	m2 := NewMatrix(2, 2)
	m2.Set(0, 0, 5.0)
	m2.Set(0, 1, 6.0)
	m2.Set(1, 0, 7.0)
	m2.Set(1, 1, 8.0)

	result := m1.Add(m2)

	if !epsilon(result.Get(0, 0), 6.0) ||
		!epsilon(result.Get(0, 1), 8.0) ||
		!epsilon(result.Get(1, 0), 10.0) ||
		!epsilon(result.Get(1, 1), 12.0) {
		t.Errorf("Matrix addition failed, got: \n%s", result.String())
	}
}

func TestSub(t *testing.T) {
	m1 := NewMatrix(2, 2)
	m1.Set(0, 0, 5.0)
	m1.Set(0, 1, 6.0)
	m1.Set(1, 0, 7.0)
	m1.Set(1, 1, 8.0)

	m2 := NewMatrix(2, 2)
	m2.Set(0, 0, 1.0)
	m2.Set(0, 1, 2.0)
	m2.Set(1, 0, 3.0)
	m2.Set(1, 1, 4.0)

	result := m1.Sub(m2)

	if !epsilon(result.Get(0, 0), 4.0) ||
		!epsilon(result.Get(0, 1), 4.0) ||
		!epsilon(result.Get(1, 0), 4.0) ||
		!epsilon(result.Get(1, 1), 4.0) {
		t.Errorf("Matrix subtraction failed, got: \n%s", result.String())
	}
}

func TestMul(t *testing.T) {
	m1 := NewMatrix(2, 3)
	m1.Set(0, 0, 1.0)
	m1.Set(0, 1, 2.0)
	m1.Set(0, 2, 3.0)
	m1.Set(1, 0, 4.0)
	m1.Set(1, 1, 5.0)
	m1.Set(1, 2, 6.0)

	m2 := NewMatrix(3, 2)
	m2.Set(0, 0, 7.0)
	m2.Set(0, 1, 8.0)
	m2.Set(1, 0, 9.0)
	m2.Set(1, 1, 10.0)
	m2.Set(2, 0, 11.0)
	m2.Set(2, 1, 12.0)

	result := m1.Mul(m2)

	if !epsilon(result.Get(0, 0), 58.0) ||
		!epsilon(result.Get(0, 1), 64.0) ||
		!epsilon(result.Get(1, 0), 139.0) ||
		!epsilon(result.Get(1, 1), 154.0) {
		t.Errorf("Matrix multiplication failed, got: \n%s", result.String())
	}
}

func TestTranspose(t *testing.T) {
	m := NewMatrix(2, 3)
	m.Set(0, 0, 1.0)
	m.Set(0, 1, 2.0)
	m.Set(0, 2, 3.0)
	m.Set(1, 0, 4.0)
	m.Set(1, 1, 5.0)
	m.Set(1, 2, 6.0)

	result := m.Transpose()

	if !epsilon(result.Get(0, 0), 1.0) ||
		!epsilon(result.Get(0, 1), 4.0) ||
		!epsilon(result.Get(1, 0), 2.0) ||
		!epsilon(result.Get(1, 1), 5.0) ||
		!epsilon(result.Get(2, 0), 3.0) ||
		!epsilon(result.Get(2, 1), 6.0) {
		t.Errorf("Matrix transpose failed, got: \n%s", result.String())
	}
}

func TestDeterminant2x2(t *testing.T) {
	m := NewMatrix(2, 2)
	m.Set(0, 0, 1.0)
	m.Set(0, 1, 2.0)
	m.Set(1, 0, 3.0)
	m.Set(1, 1, 4.0)

	det := m.Determinate()

	if !epsilon(det, -2.0) {
		t.Errorf("Expected determinant -2.0, but got %f", det)
	}
}

func TestDeterminantNonSquare(t *testing.T) {
	m := NewMatrix(2, 3)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for non-square matrix determinant")
		}
	}()
	m.Determinate() // This should panic since the matrix is not square
}

func TestInv(t *testing.T) {
	m := NewMatrix(2, 2)
	m.Set(0, 0, 4.0)
	m.Set(0, 1, 7.0)
	m.Set(1, 0, 2.0)
	m.Set(1, 1, 6.0)

	result := m.Inv()

	if !epsilon(result.Get(0, 0), 0.6) ||
		!epsilon(result.Get(0, 1), -0.7) ||
		!epsilon(result.Get(1, 0), -0.2) ||
		!epsilon(result.Get(1, 1), 0.4) {
		t.Errorf("Matrix inversion failed, got: \n%s", result.String())
	}
}

func TestInvSingularMatrix(t *testing.T) {
	m := NewMatrix(2, 2)
	m.Set(0, 0, 1.0)
	m.Set(0, 1, 2.0)
	m.Set(1, 0, 2.0)
	m.Set(1, 1, 4.0)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for singular matrix")
		}
	}()
	m.Inv() // This should panic since the matrix is singular
}
