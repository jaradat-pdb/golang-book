package structs

import "fmt"

// Shape : an interface for common methods implemented by structs representing shapes
type Shape interface {
	Area() float64
	Perimeter() float64
}

// MultiShape : defines a struct representing a collection of shape structs
type MultiShape struct {
	S []Shape
}

// Area : calculates the total area of all the shapes contained in the multishape struct
func (m MultiShape) Area() (totalArea float64) {
	for _, value := range m.S {
		totalArea += value.Area()
	}
	return
}

// Perimeter : calculates the total perimeter of all the shapes contained in the multishape struct
func (m MultiShape) Perimeter() (totalPerimeter float64) {
	for _, value := range m.S {
		totalPerimeter += value.Perimeter()
	}
	return
}

// CalculateAreaAndPerimeter : calls the Area() and Perimeter() methods of the multishape struct and builds a string containing the total area and perimeter of the shapes in the collection
func (m MultiShape) CalculateAreaAndPerimeter() string {
	return "Total Area: " + fmt.Sprintf("%g", m.Area()) + ", Total Perimeter: " + fmt.Sprintf("%g", m.Perimeter())
}
