package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var taskQueue = make(chan int, 5)
var results = make(chan int, 5)

func implementTask(wg *sync.WaitGroup) {
	for number := range taskQueue {
		results <- number * number
	}
	wg.Done()
}
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go implementTask(&wg)
	}
	wg.Wait()
	close(results)
}
func queueTasks(totalTaskCount int) {
	for i := 0; i < totalTaskCount; i++ {
		randNumber := rand.Intn(100)
		taskQueue <- randNumber
	}
	close(taskQueue)
}
func outputResult(wg *sync.WaitGroup) {
	for result := range results {
		fmt.Println("square result: ", result)
	}
	wg.Done()
}
func main() {
	startTime := time.Now()
	totalTaskCount := 50
	go queueTasks(totalTaskCount)

	var wg sync.WaitGroup
	wg.Add(1)
	go outputResult(&wg)

	noOfWorkers := 3
	createWorkerPool(noOfWorkers)
	wg.Wait()

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time: ", diff.Seconds(), "seconds")
}
