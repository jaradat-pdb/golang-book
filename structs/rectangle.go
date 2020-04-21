package structs

// Rectangle : defines a struct representing a rectangular shape
type Rectangle struct {
	L float64 //length
	W float64 //width
}

// Area : calculates the area of the rectangle struct type
func (r Rectangle) Area() float64 {
	return r.L * r.W
}

// Perimeter : calculates the perimeter of the rectangle struct type
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.L + r.W)
}
