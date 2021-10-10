package main

import (
	"fmt"
	"time"
)

func f(chanchan chan string) {
	fmt.Println("f routine")
	chanchan <- "blabla"
	fmt.Println("Send data done")
}
func main() {
	myChan := make(chan string) // define a channel of string type
	go f(myChan)
	time.Sleep(2 * time.Second) // make sure that f routine has enough time to execute
	fmt.Println("Main goroutine's finished")
}
