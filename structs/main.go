package structs

import "fmt"

func Main() {
	m := MultiShape{S: []Shape{
		Circle{R: 2},
		Rectangle{L: 5, W: 6},
		Trapezoid{A: 1, B: 2, C: 3, D: 4, H: 5},
	}}
	fmt.Println(m.CalculateAreaAndPerimeter())

	android := Android{Person: Person{Name: "Willay"}, Model: "V10"}
	android.Talk()
	android.Person.Talk()
}
