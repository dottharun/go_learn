package main

import (
	"fmt"
	"sync"
	"time"
)

/*
# GOROUTINE
- a goroutine is a lightweight thread managed by go runtime

*/

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func GoRoutineExample() {
	go say("world")
	say("hello")
}

/*
# CHANNELS
- Channels are a typed conduit(---???Array like object maybe) through which
    you can send and receive values with the channel operator, <-
- The data flows in the direction of the arrows
- By default, sends and receives block until the other side is ready
    This allows goroutines to synchronize without explicit locks or condition variables
    (---  =>all receives are blocked until all sends are completed)

- ---channel seems to be a pointer to the array, not the array itself,
    since we can pass it around function without reference

*/

// Find sum of all elems in a slice and send it to channel c
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum // send sum to c
}

func SomethingFunc(c chan int) {
	c <- 34
}

func ChannelExample() {
	// MySlice := []int{7, 2, 8, -9, 4, 0}
	MySlice := []int{0, 0, 0, 1, 1, 1}

	//making a channel of type int for the coming goroutine
	c := make(chan int)

	//find sum of half of slices separately in separate threads
	mid := len(MySlice) / 2
	go sum(MySlice[:mid], c)
	go sum(MySlice[mid:], c)

	// receive from c

	// this receive is blocked until a send is completed then the
	// runtime waits before next receive until the next receive is completed
	x := <-c
	fmt.Println(x)

	y := <-c
	fmt.Println(y)

	//send-to-channel operation can only be done inside a go-routine function
	//and a goroutine can only be invoked by running the go command on a function
	go SomethingFunc(c)
	// c <- 34 //throws error since not in go routine
	x = <-c

	fmt.Println(x, y, x+y)
}

/*
# Buffered Channel
- provide the buffer length as the second argument to make to initialize a buffered channel

Note:---maybe if i think a channel as an array available across goroutines, then buffers maybe just the size of the array

- Sends to a buffered channel block only when the buffer is full
- Receives block when the buffer is empty
*/

func bufferedChannel() {
	//channel of int type and 2 size(length of the channel array)
	ch := make(chan int, 2)

	//sends to ch
	ch <- 1
	// above is synchronous send(does'nt spawn groutine threads like below)
	// go SomethingFunc(ch)
	ch <- 3

	// ch <- 4 //will throw error since buffer is already full

	// go SomethingFunc(ch) //will not throw err since spawned thread will not do the send, garbage collected before doing anything

	//receives from ch
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch) //will err since buffer is empty
}

/* RANGE AND CLOSE
- A sender can close a channel to indicate that no more values will be sent

- A ranged for loop on a channel "receives values from the channel repeatedly until it is closed"


- recommendation from docs: do not use close from receiver,
    since after closing sending on a closed channel will cause panic

*/

func fibonacci(n int, ch chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		//sending fib nums to channel
		ch <- x

		x, y = y, x+y
	}

	//since we are done with all the sends to the channel we explicitly block it
	//from further receives
	//this close call signals the ranged for loop receive to stop the receive to exit the loop
	// close(ch)

	// close(ch) //will cause err close is allowed only once
}

func rangeAndClose() {
	//making a int channel with 10 max size
	ch := make(chan int, 10)

	//spawning a single goroutine thread to do all the sends and block the channel
	go fibonacci(cap(ch), ch)

	//receives from channel repeatedly(wait and receive) until it is closed(by a close call)
	//Note: ---there is no arrow here to indicate a receive from a channel
	// -- for each receiv we will wait until a send is done
	for i := range ch {
		fmt.Println(i)
	}
}

/*
SELECT:
  - The select statement lets a goroutine wait on multiple communication operations
     (multiple channel receives and send interchangably)

  - A select blocks until one of its cases can run, then it executes that case
    It chooses one at random if multiple are ready
*/

func fibonacci2(ch chan int, quit chan int) {
	x, y := 0, 1

	//does sends to ch channel until quit channel has a send
	//  (--only then we can do a receive from quit channel)
	//Note: this function is ran syncronously
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

/* Default case in chanells
- The default case in a select is run if no other case is ready
-  Use a default case to try a send or receive without blocking

---Use for not blocking the runtime when channels are checked
*/

func selectExample() {
	//making two channels - ch and quit
	ch := make(chan int)
	quit := make(chan int)

	//spawing a goroutine for receives from ch??? and a exit send to quit
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		quit <- 0
	}()

	fibonacci2(ch, quit)
}

func defaultSelection() {
	//receives from inbuilt time.Time channel or these itself are channel not sure???
	//???maybe its a channel where only receives can be done(maybe an interface for a channel)
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println(".")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

/* SYNC MUTEX

-- to prevent data races when using multiple channels that access same data

-- ----a mutex object just seems to be a integer with state: on, off / locked, unlocked
            so may here when the mutex is locked only one goroutine thread can access the variable
            -but normally any goroutine thread can access any variable
*/

// SafeCounter is safe to use concurrently.
// ---custom type strct
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
// ---but same key is given for all spawned goroutines to introduce data race problems
func (c *SafeCounter) Inc(key string) {
	//safeCounter object c is now locked
	c.mu.Lock()

	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++

	//c is now unlocked for other goroutines
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()

	return c.v[key]
}

func syncMutexExample() {
	c := SafeCounter{v: make(map[string]int)}

	//spawing too many goroutines to show the data race problem
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)

	//trying to access the data used for writing in the spawned go routines
	fmt.Println(c.Value("somekey"))
}

func main() {
	// GoRoutineExample()
	// say("nice") //only executed after completing prev function

	// ChannelExample()
	// bufferedChannel()
	// rangeAndClose()
	// selectExample()
	defaultSelection()
	syncMutexExample()
}
