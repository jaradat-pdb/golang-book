package recursion

import "fmt"

func Main() {
	for i := 0; i <= 15; i++ {
		fmt.Println("Factorial of", i, "is", Factorial(uint(i)))
	}
}
