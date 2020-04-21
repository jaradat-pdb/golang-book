package slices

// AppendSlice : make a new slice and append entries to new slice
func AppendSlice() (y []int) {
	x := []int{1, 2, 3, 4, 5}
	y = append(x, 6, 7)
	return
}

// CopySlice : make a new slice by copying a previous slice
func CopySlice() (x []float64, y []float64) {
	x = []float64{1, 2, 3}
	y = make([]float64, 4)
	copy(y, x)
	return
}
