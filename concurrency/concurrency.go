package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

func F(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(250)))
	}
}

func Pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping " + fmt.Sprintf("%d", i)
	}
}

func Ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "pong " + fmt.Sprintf("%d", i)
	}
}

func Printer(c <-chan string) {
	for {
		fmt.Println(<-c)
		time.Sleep(time.Second * 1)
	}
}
