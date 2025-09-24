package matrix

import "gonum.org/v1/gonum/mat"

func VecDenseFromArray(arr [][]float64) *mat.VecDense {
	if len(arr) == 0 || len(arr[0]) != 1 {
		return mat.NewVecDense(0, nil)
	}

	data := make([]float64, len(arr))
	for i := range arr {
		data[i] = arr[i][0]
	}

	return mat.NewVecDense(len(arr), data)
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

// VecDenseAppend appends a new *mat.VecDense vector to an existing one,
func VecDenseAppend(v *mat.VecDense, newElem *mat.VecDense) *mat.VecDense {
	if newElem.Len() == 0 {
		return v
	}

	newLen := v.Len() + newElem.Len()
	newVec := ResizeVecDense(v, newLen)

	for i := range newElem.Len() {
		newVec.SetVec(v.Len()+i, newElem.AtVec(i))
	}

	return newVec
}

// VecToSlice converts a *mat.VecDense to a slice of float64.
func VecToSlice(v *mat.VecDense) []float64 {
	s := make([]float64, v.Len())
	for i := range v.Len() {
		s[i] = v.AtVec(i)
	}
	return s
}

// VecDenseStack vertically stacks two *mat.VecDense vectors.
func VecDenseStack(v1, v2 *mat.VecDense) *mat.VecDense {
	if v1 == nil || v1.Len() == 0 {
		return v2
	}
	if v2 == nil || v2.Len() == 0 {
		return v1
	}

	newLen := v1.Len() + v2.Len()
	newVec := mat.NewVecDense(newLen, nil)

	for i := range v1.Len() {
		newVec.SetVec(i, v1.AtVec(i))
	}
	for j := range v2.Len() {
		newVec.SetVec(v1.Len()+j, v2.AtVec(j))
	}

	return newVec
}
