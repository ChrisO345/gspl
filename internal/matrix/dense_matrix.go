package matrix

import "gonum.org/v1/gonum/mat"

// Eye creates a dense diagonal matrix of the given size with ones on the diagonal.
func Eye(size int) *mat.DiagDense {
	data := make([]float64, size)
	for i := range data {
		data[i] = 1.0
	}
	return mat.NewDiagDense(size, data)
}

// ResizeMatDense resizes a *mat.Dense matrix to newRows x newCols,
// copying over the existing data as much as possible.
func ResizeMatDense(m *mat.Dense, newRows, newCols int) *mat.Dense {
	newMat := mat.NewDense(newRows, newCols, nil)

	if m != nil {
		oldRows, oldCols := m.Dims()
		minRows := min(oldRows, newRows)
		minCols := min(oldCols, newCols)

		for i := range minRows {
			for j := range minCols {
				newMat.Set(i, j, m.At(i, j))
			}
		}
	}

	return newMat
}

// ExtractColumns extracts columns from a *mat.Dense matrix based on the provided indices.
func ExtractColumns(A *mat.Dense, indices *mat.VecDense) *mat.Dense {
	m, n := A.Dims()
	numIndices := indices.Len()

	if numIndices == 0 {
		return mat.NewDense(m, 0, nil) // Return empty matrix if no indices
	}

	extracted := mat.NewDense(m, numIndices, nil)
	for i := range m {
		for j := range numIndices {
			colIndex := int(indices.AtVec(j))
			if colIndex >= 0 && colIndex < n {
				extracted.Set(i, j, A.At(i, colIndex))
			} else {
				extracted.Set(i, j, 0.0) // Handle out-of-bounds indices
			}
		}
	}

	return extracted
}

// MatDenseFromArray converts a 2D slice of float64 to a *mat.Dense matrix.
func MatDenseFromArray(arr [][]float64) *mat.Dense {
	rows := len(arr)
	if rows == 0 {
		return mat.NewDense(0, 0, nil)
	}
	cols := len(arr[0])
	data := make([]float64, rows*cols)

	for i := range rows {
		for j := range cols {
			data[i*cols+j] = arr[i][j]
		}
	}

	return mat.NewDense(rows, cols, data)
}

// MatDenseToArray converts a *mat.Dense matrix to a 2D slice of float64.
func MatDenseToArray(A *mat.Dense) [][]float64 {
	if A == nil {
		return nil
	}

	m, n := A.Dims()
	arr := make([][]float64, m)
	for i := range m {
		arr[i] = make([]float64, n)
		for j := range n {
			arr[i][j] = A.At(i, j)
		}
	}

	return arr
}

// MatDenseAppendColumn appends a column vector to a *mat.Dense matrix.
func MatDenseAppendColumn(A *mat.Dense, col *mat.VecDense) *mat.Dense {
	if A == nil || col == nil {
		return A
	}

	m, n := A.Dims()
	if col.Len() != m {
		panic("Column length must match the number of rows in the matrix")
	}

	newMat := mat.NewDense(m, n+1, nil)
	newMat.Copy(A)

	for i := range m {
		newMat.Set(i, n, col.AtVec(i))
	}

	return newMat
}

func MatDenseAppendRow(A *mat.Dense, row *mat.Dense) *mat.Dense {
	if A == nil || row == nil {
		return A
	}

	m, n := A.Dims()
	rowM, rowN := row.Dims()
	if rowN != n {
		panic("Row length must match the number of columns in the matrix")
	}

	newMat := mat.NewDense(m+rowM, n, nil)
	newMat.Copy(A)

	for i := range rowM {
		for j := range n {
			newMat.Set(m+i, j, row.At(i, j))
		}
	}

	return newMat
}
