package arrays

import "fmt"

// IterateArray : iterates through a parameter passed float64 array and returns its entries as a string
func IterateArray(args ...float64) (str string) {
	for _, value := range args {
		str += fmt.Sprintf("%g", value) + ", "
	}
	return
}

// AverageArrays : iterates through the entries of an array and calculates the average
func AverageArrays(x []float64, y []float64) (avgX float64, avgY float64) {
	sumX := 0.0
	sumY := 0.0

	/*
		The Go compiler throws an error when variables are unused,
		thus when using the syntax below where an array (or other iterable)
		object is being cycled through without the need for an iterator
		variable (similar to the Java & JS forEach approach) use an
		underscore to tell the compiler the variable is not needed
	*/
	for _, value := range x {
		sumX += value
	}
	avgX = sumX / float64(len(x))

	for _, value := range y {
		sumY += value
	}
	avgY = sumY / float64(len(y))

	return avgX, avgY
}

// AverageFloat64Array : calculates the average of the entries of a parameter passed float64 array
func AverageFloat64Array(arr []float64) float64 {
	total := 0.0
	for _, value := range arr {
		total += value
	}
	return total / float64(len(arr))
}

// SmallestNumberInArray : determines the smallest number in an (int) array (or slice)
func SmallestNumberInArray(arr []int) (low int) {
	low = arr[0]
	for _, value := range arr {
		if value < low {
			low = value
		}
	}
	// can use just return when the name of the return type is provided in the signature
	return
}
