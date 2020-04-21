package structs

import "math"

// Circle : defines a struct representing a circular shape
type Circle struct {
	R float64 //radius
}

// Area : calculates the area of the circle struct type
func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

// Perimeter : calculates the perimeter of the circle struct type
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}
