package main

import (
	"fmt"
	"sync"
	"time"
)

func pinger(wg *sync.WaitGroup) {
	for i := 1; i <= 5; i++ {
		fmt.Println("ping", i)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("pinger ended\n")
	wg.Done()
}
func ponger(wg *sync.WaitGroup) {
	for i := 1; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("pong", i)
	}
	fmt.Printf("ponger ended\n")
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go pinger(&wg)
	wg.Add(1)
	go ponger(&wg)
	wg.Wait()
	fmt.Println("All subgoroutines finished")
}
