package concurrency

import "fmt"

func Main() {
	for i := 0; i < 10; i++ {
		go F(i)
	}
	var input string
	fmt.Scanln(&input)

	var c chan string = make(chan string)
	go Pinger(c)
	go Ponger(c)
	go Printer(c)
	fmt.Scanln(&input)
}
