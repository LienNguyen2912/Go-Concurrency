package main

import (
	"fmt"
)

func f(chanchan chan string) {
	fmt.Println("f routine")
}
func main() {
	myChan := make(chan string)
	go f(myChan)
	<-myChan // read data from channel
	fmt.Println("Main goroutine's finished")
}
