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
