package main

import (
	"fmt"
	"sync"
)

func counting(i int, wg *sync.WaitGroup) int {
	defer wg.Done()
	fmt.Println("Counting", i)
	return i
}

func main() {
	// concurrency in goroutine
	var wg sync.WaitGroup

	// for i :=0 ; i<10;i++ {
	// 	wg.Add(1)
	// 	go counting(i,&wg)
	// }

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("Counting in anonymous function", i,
				"value of wg is:",
				wg) // using a closure property here
		}(i)
	}

	// value of wg is: {{} {{} {} 42949672961} 0}

	// adding sleep to see the output of goroutine
	// time.Sleep(2 * time.Second)
	// waint here extead of sleep
	wg.Wait()
}

// This Go code demonstrates launching multiple goroutines using both a named function and an anonymous function. It prints numbers concurrently and uses time.Sleep to allow goroutines to finish before the program exits. The comments explain that proper synchronization should use sync.WaitGroup instead of sleep.
