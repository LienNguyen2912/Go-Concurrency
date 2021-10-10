# Go concurrency
## Goroutines
>Goroutines are functions or methods that run concurrently with other functions or methods. Goroutines can be thought of as lightweight threads. The cost of creating a Goroutine is tiny when compared to a thread. Hence it's common for Go applications to have thousands of Goroutines running concurrently.</br>

To create a goroutine we prefix the keyword _go_ before the function invocation:
```sh
package main

import "fmt"

func f() {
	fmt.Println("f goroutine")
}
func main() {
	// Create other routine
	go f()
	fmt.Println("main gorountine")
}
```
This program consists of two goroutines. The first goroutine is implicit and is the main function itself.</br>
The output is</br>
![mainRoutineOnly1](https://user-images.githubusercontent.com/73010204/136644466-fd311588-91b0-44bc-898d-a0486eb6be90.PNG)</br>
**What a suprise!** There is only text printed from the main routine. Where is f goroutine?</br>
Unlike functions, the control does not wait for the Goroutine to finish executing. The control returns immediately to the next line of code after the Goroutine call and any return values from the Goroutine are ignored.</br>
By that way, after printing _"main gorountine"_ the main goroutines is completed hence the program will be terminated, and all other Goroutines did not get a chance to run.</br>

Let's ask the main goroutine to wait for a second.
```sh
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
```
![printBoth2](https://user-images.githubusercontent.com/73010204/136644467-aa9fc90d-da15-47ee-8415-f5898a63d516.PNG)</br>
You see, we got text both!</br>
We can create as many goroutines as we want. We will create 5 routines and see they run simultaneously.
```sh
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
```
Our result may different each time executing because we randomed a waiting time after printing out each number.</br>
![routine3](https://user-images.githubusercontent.com/73010204/136644869-538c62ef-8324-4560-92d8-13d5e611833a.PNG)</br>
We can visualize how this program worked as the following image</br>
![routine3Visualize](https://user-images.githubusercontent.com/73010204/136654568-6a3781ba-bab4-42f7-8164-5a5e840c0a7f.PNG)

## Channels
Channels provide a way for two goroutines to communicate with one another by sending and receiving data between them. Channel can be used for synchronizing two goroutines.
What we have to remember when using channel are:
-  Each channel has a type associated with it. Other type is not allowed to be passed.
-  A channel has send direction and receive direction specified by the arrow.  A bi-directional channel can be casted to send-only or receive-only channel, but the reverse is impossible.
-  Sends and receives are blocking by default

See how the control is blocked in the send statement till there is other goroutine read out that data.
```sh
package main

import (
	"fmt"
	"time"
)

func f(chanchan chan string) {
	fmt.Println("f routine")
	chanchan <- "blabla"  // send data to channel
	fmt.Println("Send data done")
}
func main() {
	myChan := make(chan string)
	go f(myChan)
	time.Sleep(2 * time.Second) // make sure that f routine has enough time to execute
	fmt.Println("Main goroutine's finished")
}
```
Here is the output</br>
![send_blocking](https://user-images.githubusercontent.com/73010204/136654363-1c35f332-4c4a-4a4f-9cf2-804ec2d6a31b.png)</br>
We can visualize how it worked as below.</br>
![send_blocking_visual](https://user-images.githubusercontent.com/73010204/136654364-9c6131e4-d412-4d5a-b1ea-410a3a36336d.png)</br>
When there is **other** goroutine receives that data, the send statement is unblocked.
```sh
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
```
![send_blocking2](https://user-images.githubusercontent.com/73010204/136654367-a84052c9-3eaa-4e77-af7e-d93a7503ccd9.png)</br>
Visualize it (gap distances mean nothing)<br/>
![send_blocking_visual2](https://user-images.githubusercontent.com/73010204/136654365-cd12f2a5-a581-48a4-af78-8023cf897471.png)<br/>
As you see, the main goroutine is blocked in read statement also for waiting data from the channel. That's why without Sleep vocation, _"Send data done"_ text was printed out.</br>
Let 's see how the control is blocked in the read statement by another example.
```sh
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

```
The output is dealock cause main goroutine kept waiting for a routine to send data to the channel but no one did.</br>
![deadlock](https://user-images.githubusercontent.com/73010204/136654991-49afdbf9-1875-46cb-a15e-8c9866004c6d.PNG)</br>
## Another channel example
Let's write a program having a goroutine for adding calculation, another for mulitplying calculation then finally sum them all in the main goroutine.
```sh
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
```
As you see, we casted bi-directional channels to either send-only or receive-only channels when pass to other functions.
The output may different each time excuting.</br>
![add_mul](https://user-images.githubusercontent.com/73010204/136657209-be191a44-0cf8-4442-be5e-8a2d1613e599.PNG)</br>
There are 4 goroutines run simultaneously like this.</br>
![add_mul_visual](https://user-images.githubusercontent.com/73010204/136657452-dfa4c33b-431c-4b02-964c-77816d37f872.PNG)</br>

## Buffered Channels
Channels can be buffered. Provide the buffer capacity as the second argument to initialize a buffered channel.
By default, channel have capacity as 0. In other words, without specifying capacity parameter, capacity is 0 and channels are unbuffered.
- Sends to a buffered channel block when the buffer is full. 
- Receives block when the buffer is empty.
- Different with the capacity, the length of the buffered channel is element numbers currently queued in.
```sh
package main

import (
	"fmt"
)

func main() {
	myChan := make(chan string, 2) // initialize a buffered channel
	myChan <- "first"
	myChan <- "second"
	fmt.Println(<-myChan)
	fmt.Println(<-myChan)
	fmt.Println("Main goroutine ended")
}
```
![bufferedChannel1](https://user-images.githubusercontent.com/73010204/136684894-2f5d7e7a-1b23-4408-893f-7f7824208bb5.PNG)</br>
But when the buffered channel keeps being full without any other routines read/remove data from it like this
```sh
package main

import (
	"fmt"
)

func main() {
	myChan := make(chan string, 2) // initialize a buffered channel
	myChan <- "first"
	myChan <- "second"
	myChan <- "third"
	fmt.Println(<-myChan)
	fmt.Println(<-myChan)
	fmt.Println(<-myChan)
	fmt.Println("Main goroutine ended")
}
```
We are getting a deadlock certainly :D.</br>
![bufferedChannel2](https://user-images.githubusercontent.com/73010204/136684897-d90fbdb1-90c0-47f6-b70a-d1178f4158fa.PNG)</br>
But the below works.
```sh
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
```
![bufferedChannel3](https://user-images.githubusercontent.com/73010204/136684898-9d883366-e427-45ac-8cbd-e27b3f3948b9.PNG)</br>
Finally for buffered channels, let's see the difference between _len_ and _cap_
```sh
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
```
![bufferedChannel_len_capacity](https://user-images.githubusercontent.com/73010204/136684980-5d5c8f1b-741d-4fe7-b0ad-e7e9529c75be.PNG)</br>

## WaitGroup
WaitGroup is a way to inform a routine that all of its subroutines are completed so it can further proceed.
Without WaitGroup, we can use Sleep or receive/send statement for blocking, but sometimes it may not make sense. WaitGroup can be a solution.
```sh
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
```
We used WaitGroup to let main Goroutine wait for pinger goroutine and ponger goroutine till they are both finished. Your output may be different in executation order.</br>
![waitGroup](https://user-images.githubusercontent.com/73010204/136686115-243a4b92-40ef-4803-bf36-55b25011a201.PNG)</br>
Here are some notes about WaitGroup:
- WaitGroup works by using a counter.
- _.Add(number int)_ increases the WaitGroup's counter by the value passed to.
- _.Done()_ decreases the WaitGroup's counter by 1.
- _Wait()_ method will blocks the Goroutine in which it is called until the WaitGroup's counter becomes 0.

## Worker Pools
By using buffered channels and WaitGroup, we can create thread pools in golang.</br>
Let's create a worker pool like the image below</br>
![workerPool](https://user-images.githubusercontent.com/73010204/136691973-50af60c4-9c0f-43c5-abc9-92eb639b3f12.PNG)</br>
-  Create a pool of Goroutines having 3 subgoroutines( workers).
-  Workers get data from the input buffered channel, implement tasks then ouput results to the output buffered channel
-  Task implementation is to square a number, very simple, just for simulation.
-  In order not to get confused, I intentionally chose the total number of tasks , the number of input/output buffered channels and the number of subroutines( workers) to be different.
```sh
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
```
Your result may vary.</br>
![workerPool2](https://user-images.githubusercontent.com/73010204/136691818-05163e0f-f24f-4f26-9935-732f338fa430.PNG)</br>
Try to increase the buffered channel capacity, and increase the number of workers in pools, you will see the total time are reduced!

## Select

## Mutex
