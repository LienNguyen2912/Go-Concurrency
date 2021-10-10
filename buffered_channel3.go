package main

import (
	"fmt"
)

func main() {
	myChan := make(chan string, 2) // initialize a buffered channel
	myChan <- "first"
	myChan <- "second"
	go func() { myChan <- "third" }()
	fmt.Println(<-myChan)
	fmt.Println(<-myChan)
	fmt.Println(<-myChan)
	fmt.Println("Main goroutine ended")
}
