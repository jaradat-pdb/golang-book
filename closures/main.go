package closures

import "fmt"

func Main() {
	nextEven := MakeEvenGenerator()
	fmt.Print("NextEven closure function: ")
	for i := 0; i < 10; i++ {
		fmt.Print(nextEven(), ", ")
	}
	fmt.Println()

	nextOdd := MakeOddGenerator()
	fmt.Print("NextOdd closure function: ")
	for i := 0; i < 10; i++ {
		fmt.Print(nextOdd(), ", ")
	}
	fmt.Println()
}
