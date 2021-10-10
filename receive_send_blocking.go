package main

import (
	"fmt"
)

func f(chanchan chan string) {
	fmt.Println("f routine")
	chanchan <- "blabla" // send data to channel
	fmt.Println("Send data done")
}
func main() {
	myChan := make(chan string)
	go f(myChan)
	<-myChan // receive data from channel
	fmt.Println("Main goroutine's finished")
}
