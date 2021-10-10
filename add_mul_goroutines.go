package main

import (
	"fmt"
	"time"
)

func print(c <-chan string) {
	for {
		fmt.Println(<-c)
		time.Sleep(time.Millisecond * 100)
	}
}
func calcAdd(number int, add chan<- int) {
	addVal := 0
	addC := make(chan string)
	go print(addC)
	for i := 1; i <= number; i++ {
		addVal += i
		addC <- fmt.Sprintf("add %d", i)
	}
	add <- addVal
}
func calcMul(number int, mul chan<- int) {
	mulVal := 1
	mulC := make(chan string)
	go print(mulC)
	for i := 1; i <= number; i++ {
		mulVal *= i
		mulC <- fmt.Sprintf("mul %d", i)
	}
	mul <- mulVal
}
func main() {
	add := make(chan int)
	mul := make(chan int)
	go calcAdd(3, add)
	go calcMul(3, mul)
	addRet, mulRet := <-add, <-mul
	fmt.Println("Final result: ", addRet+mulRet)
}
