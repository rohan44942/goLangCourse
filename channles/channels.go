package main

import (
	"fmt"
	"sync"
	"time"
)

func printnum(numchan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range numchan {
		// print this num here
		fmt.Println("value from num chan is: ", num)
		time.Sleep(time.Second)
	}
}

func sum(sumchan chan int, a int, b int) {
	// defer wg.Done()
	sumchan <- (a + b)
}

func main() {
	// var numchan = make(chan int)
	var sumchan = make(chan int)
	// var wg sync.WaitGroup

	// call a functin
	// wg.Add(1)
	// go printnum(numchan, &wg)
	go sum(sumchan, 5, 10)
	fmt.Println("Sum of 5 and 10 is:", <-sumchan)
	// close(sumchan)

	// for i := 0; i < 20; i++ {
	// 	numchan <- i
	// }
	close(sumchan)
	// close(numchan)
	// wg.Wait()
	// fmt.Println("Welcome to channels in golang")
}

// how channels work in golang

// channels are used to communicate between goroutines
// channels are typed, you can create a channel of any type
// channels are created using the make function
// channels are blocking, when you send a value to a channel, the sender is blocked until the receiver receives the value
// channels can be buffered or unbuffered
// unbuffered channels block the sender until the receiver receives the value
// buffered channels allow the sender to send values without blocking until the buffer is full
