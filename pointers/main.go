package pointers

import "fmt"

func Main() {
	point := 5
	fmt.Print("Original value of point=", point)
	// the & operator returns the address (in memory) of variable "point"
	ZeroPoint(&point)
	fmt.Print(", modified value of point=", point)
	OnePoint(&point)
	fmt.Print(", newly modified value of point=", point, "\n")

	pointPtr := new(int)
	TwoPoint(pointPtr)
	fmt.Println("Value of pointPtr=", *pointPtr)

	a := 111
	b := 212
	fmt.Println("Initial values: a=", a, ", b=", b)
	SwapValues(&a, &b)
	fmt.Println("Modified values: a=", a, ", b=", b)
}
