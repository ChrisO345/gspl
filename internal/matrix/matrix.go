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

// ResizeVecDense resizes a *mat.VecDense vector to newLen,
// copying over the existing data as much as possible.
func ResizeVecDense(v *mat.VecDense, newLen int) *mat.VecDense {
	newVec := mat.NewVecDense(newLen, nil)

	if v != nil {
		oldLen := v.Len()
		minLen := min(oldLen, newLen)

		for i := range minLen {
			newVec.SetVec(i, v.AtVec(i))
		}
	}

	return newVec
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
