package structs

// Trapezoid : defines a struct representing a trapezoidal shape
type Trapezoid struct {
	A float64 //top base
	B float64 //bottom base
	C float64 //left side
	D float64 //right side
	H float64 //height
}

// Area : calculates the area of the trapezoid struct type
func (t Trapezoid) Area() float64 {
	return ((t.A + t.B) / 2) * t.H
}

// Perimeter : calculates the perimeter of the trapezoid struct type
func (t Trapezoid) Perimeter() float64 {
	return t.A + t.B + t.C + t.D
}
