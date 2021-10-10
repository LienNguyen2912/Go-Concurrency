package main

import (
	"fmt"
	"time"
)

func f() {
	fmt.Println("f goroutine")
}
func main() {
	// Create other routine
	go f()
	time.Sleep(1 * time.Second)
	fmt.Println("main gorountine")
}
