package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 3; i++ {
		fmt.Println("goroutine no.", n, " :", i)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(250)))
	}
}
func main() {
	for i := 0; i < 5; i++ {
		go f(i)
	}
	time.Sleep(2 * time.Second)
}
