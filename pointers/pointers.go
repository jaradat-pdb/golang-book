package pointers

// ZeroPoint : updates the value of the variable referenced by the pointer to 0
func ZeroPoint(xPtr *int) {
	*xPtr = 0
}

// OnePoint : updates the value of the variable referenced by the pointer to 1
func OnePoint(xPtr *int) {
	*xPtr = 1
}

// TwoPoint : updates the value of the variable referenced by the pointer to 2
func TwoPoint(xPtr *int) {
	*xPtr = 2
}

// SwapValues : swaps the value of the two parameter passed integer pointers
func SwapValues(xPtr *int, yPtr *int) {
	iniX := *xPtr
	*xPtr = *yPtr
	*yPtr = iniX
}
