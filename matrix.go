package gspl

import (
	"fmt"
	"math"
)

type Numeric interface {
	~int | ~int64 | ~float32 | ~float64
}

type Matrix struct {
	Rows, Columns int

	Values [][]float64
}

// NewMatrix returns a new matrix with 0-cells.
func NewMatrix(rows, columns int) *Matrix {
	m := &Matrix{
		Rows:    rows,
		Columns: columns,
		Values:  make([][]float64, rows),
	}
	for i := range m.Values {
		m.Values[i] = make([]float64, columns)
	}
	return m
}

// FromArray creates a matrix from a Numeric array
func FromArray[T Numeric](rows, columns int, matrix [][]T) *Matrix {
	// Create a new Matrix of float64 type
	m := NewMatrix(rows, columns)

	// Convert each element from T to float64
	for i := range rows {
		for j := range columns {
			m.Values[i][j] = float64(matrix[i][j])
		}
	}

	return m
}

// String returns the formatted version of the Matrix.
func (m *Matrix) String() string {
	s := fmt.Sprintf("Matrix %d x %d\n", m.Rows, m.Columns)
	for i := range m.Values {
		s += fmt.Sprintf("%v\n", m.Values[i])
	}
	return s
}

// Size returns the size of the Matrix
func (m *Matrix) Size() (int, int) {
	return m.Rows, m.Columns
}

// Length returns the largest dimension of the Matrix
func (m *Matrix) Length() int {
	return max(m.Rows, m.Columns)
}

// Get returns the value at the specified row and column of the Matrix
func (m *Matrix) Get(i, j int) float64 {
	return m.Values[i][j]
}

// Set sets the value at the specified row and column of the Matrix
func (m *Matrix) Set(i, j int, value float64) {
	if i < 0 || i >= m.Rows || j < 0 || j >= m.Columns {
		panic("Index out of bounds")
	}
	m.Values[i][j] = value
}

// Eye generates an identity matrix of the specified size
func Eye(size int) *Matrix {
	m := NewMatrix(size, size)
	for i := range size {
		m.Values[i][i] = 1.0
	}
	return m
}

// Add returns the sum of two matrices
func (m *Matrix) Add(m2 *Matrix) *Matrix {
	if m.Rows != m2.Rows || m.Columns != m2.Columns {
		panic("Matrix dimensions do not match")
	}

	s := NewMatrix(m.Rows, m.Columns)
	for i := range m.Rows {
		for j := range m.Columns {
			s.Values[i][j] = m.Values[i][j] + m2.Values[i][j]
		}
	}
	return s
}

// Sub returns the difference of two matrices
func (m *Matrix) Sub(m2 *Matrix) *Matrix {
	if m.Rows != m2.Rows || m.Columns != m2.Columns {
		panic("Matrix dimensions do not match")
	}

	s := NewMatrix(m.Rows, m.Columns)
	for i := range m.Rows {
		for j := range m.Columns {
			s.Values[i][j] = m.Values[i][j] - m2.Values[i][j]
		}
	}
	return s
}

// Mul returns the product of two matrices
func (m *Matrix) Mul(m2 *Matrix) *Matrix {
	if m.Columns != m2.Rows {
		panic("Matrix dimensions do not match")
	}

	s := NewMatrix(m.Rows, m2.Columns)
	for i := range m.Rows {
		for j := range m2.Columns {
			for k := range m.Columns {
				s.Values[i][j] += m.Values[i][k] * m2.Values[k][j]
			}
		}
	}

	return s
}

// Transpose returns the transpose of the Matrix
func (m *Matrix) Transpose() *Matrix {
	if m.Rows == 0 || m.Columns == 0 {
		return nil
	}

	s := NewMatrix(m.Columns, m.Rows)
	for i := range m.Rows {
		for j := range m.Columns {
			s.Values[j][i] = m.Values[i][j]
		}
	}
	return s
}

// Determinate returns the determinant of the Matrix via Laplace expansion
func (m *Matrix) Determinate() float64 {
	if m.Rows != m.Columns {
		panic("Matrix is not square")
	}

	// Base case for 2 x 2 Matrix
	if m.Rows == 2 && m.Columns == 2 {
		return m.Values[0][0]*m.Values[1][1] - m.Values[0][1]*m.Values[1][0]
	}

	// Recursive case for larger matrices
	det := 0.0
	for i := range m.Rows {
		subMatrix := NewMatrix(m.Rows-1, m.Columns-1)
		for j := 1; j < m.Rows; j++ {
			for k := range m.Columns {
				if k < i {
					subMatrix.Values[j-1][k] = m.Values[j][k]
				} else if k > i {
					subMatrix.Values[j-1][k-1] = m.Values[j][k]
				}
			}
		}
	}

	return det
}

// Inv returns the inverse of the Matrix via Gauss-Jordan elimination
func (m *Matrix) Inv() *Matrix {
	rows, cols := m.Size()
	if rows != cols {
		panic("Matrix is not square")
	}

	// Make a deep copy of the original matrix
	copyMatrix := NewMatrix(rows, cols)
	for i := range rows {
		for j := range cols {
			copyMatrix.Values[i][j] = m.Values[i][j]
		}
	}

	// Create an identity matrix
	inv := Eye(rows)

	for i := range rows {
		pivot := copyMatrix.Values[i][i]
		if math.Abs(pivot) < 1e-10 {
			// Try to swap with a lower row
			swapped := false
			for j := i + 1; j < rows; j++ {
				if math.Abs(copyMatrix.Values[j][i]) > 1e-10 {
					copyMatrix.Values[i], copyMatrix.Values[j] = copyMatrix.Values[j], copyMatrix.Values[i]
					inv.Values[i], inv.Values[j] = inv.Values[j], inv.Values[i]
					pivot = copyMatrix.Values[i][i]
					swapped = true
					break
				}
			}
			if !swapped {
				panic("Matrix is singular and cannot be inverted")
			}
		}

		// Normalize the pivot row
		for j := range cols {
			copyMatrix.Values[i][j] /= pivot
			inv.Values[i][j] /= pivot
		}

		// Eliminate other rows
		for k := range rows {
			if k == i {
				continue
			}
			factor := copyMatrix.Values[k][i]
			for j := range cols {
				copyMatrix.Values[k][j] -= factor * copyMatrix.Values[i][j]
				inv.Values[k][j] -= factor * inv.Values[i][j]
			}
		}
	}

	return inv
}
