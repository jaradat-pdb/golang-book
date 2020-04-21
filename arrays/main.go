package arrays

import (
	"fmt"
)

func Main() {
	// passes an array to a variadic function
	fmt.Println(IterateArray([]float64{33.333333334, 44, 55, 66.667, 77, 88, 99.123}...))
	// the trailing comma is required in the multi-line syntax below
	x, y := AverageArrays([]float64{78, 69, 15, 49, 93}, []float64{
		87,
		66,
		51,
		94,
		92,
	})
	fmt.Println("Avg of x[]:", x, "|| Avg of y[]:", y)
	fmt.Println("Avg of float64 arr[]:", AverageFloat64Array([]float64{78.2, 69.9, 15.3, 49.7, 93.4}))
	fmt.Println("Avg of float64 arr[]:", AverageFloat64Array([]float64{11.1, 22.2, 33.3, 44.4, 55.5, 66.6, 77.7, 88.8, 99.9}))
	// the trailing comma is required in the multi-line syntax below
	fmt.Println("Smallest number in list:", SmallestNumberInArray([]int{
		48, 96, 86, 68,
		57, 22, 63, 70,
		37, 34, 83, 17,
		19, 97, 9, 17,
	}))
}
