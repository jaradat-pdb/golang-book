package concurrency

import "fmt"

func Main() {
	for i := 0; i < 10; i++ {
		go F(i)
	}
	var input string
	fmt.Scanln(&input)
}
