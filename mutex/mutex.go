package main

import (
	"fmt"
	"sync"
)

func addtion(a int ,b int,wg *sync.WaitGroup){
	wg.Done()
	fmt.Println("Addition is ",a+b)
}

func main(){

	// var wg sync.WaitGroup
	var mu sync.Mutex // mutex variable
	// making go channels 
	var cha = make(chan int)
	var chb = make(chan int)

	// wg.Add(1)
	// go addtion(2,3,&wg);
	cha <- 2
	chb <- 3
	for i:=0;i<2;i++{
		 func(){
			mu.Lock() // lock the resource
			defer mu.Unlock() // unlock the resource
			a:= <- cha
			b:= <- chb
			fmt.Println("Multiplication is ",a*b)
		}()

	}
	close(cha)
	close(chb)
	

	// wg.Wait()

	fmt.Println("Welcome to mutex in golang")
}

// race condition in golang
// mutex is used to avoid
// race condition , two goroutine trying to access the same resource at the same time
// mutex is a lock
// mutex is a mutual exclusion
// mutex is a binary semaphore
// mutex is used to lock a resource
// mutex is used to unlock a resource