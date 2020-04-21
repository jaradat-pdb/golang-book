package problems

import (
	"fmt"
)

func Main() {
	fmt.Println("The greatest is:", GreatestInList([]int{1, 2, 10, 113, 4, 66, 13, 69, 7, 8, 9}...))

	for i := 0; i <= 20; i++ {
		var isEvenStr string
		j, isEven := Halfer(i)
		if isEven {
			isEvenStr = ", it is even"
		} else {
			isEvenStr = ", it is odd"
		}
		fmt.Println("Half of", i, " is", j, isEvenStr)
	}

	for i := 10; i <= 28; i++ {
		fmt.Println("Fibonacci sequence of", i, ":", FibonacciSequencer(uint(i)))
	}
}
