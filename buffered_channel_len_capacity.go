package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 3)
	fmt.Println("capacity =", cap(ch))
	fmt.Println("length =", len(ch))
	ch <- "1st"
	fmt.Println("length =", len(ch))
	ch <- "2nd"
	fmt.Println("length =", len(ch))
	ch <- "3rd"
	fmt.Println("length =", len(ch))
	<-ch
	fmt.Println("length =", len(ch))
	<-ch
	fmt.Println("length =", len(ch))
	<-ch
	fmt.Println("length =", len(ch))
	fmt.Println("main ended")
}
