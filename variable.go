package main

import "fmt"

func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func calculatepower(b int, powerchan chan int64) int64 {
	// calculate power recursion
	if b == 0 {
		return 1
	}
	a := <-powerchan
	powerchan <- a
	return a * calculatepower(b-1, powerchan)
}


func main() {
	// Declare a variable with a specific type
	// var age int = 30
	// newage := 25
	// Declare a variable without specifying the type
	// var name = "rohan"
	// fmt.Println("age and newage with name are :", age, newage, name)

	// switch condition := age; condition {
	// case 30:
	// 	fmt.Println("age is 30, condition is:", condition)
	// default:
	// 	fmt.Println("age is not 30, condition is:", condition)
	// }

	// increament := counter()

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Counter value is:", increament())
	// }


	// testing routines and passing data through channels

	// taking data from console 
	var a int64
	var b int
	// make channels 

	var powerchan = make(chan int64)
	fmt.Println("enter the values: ")
	fmt.Scan(&a)
	fmt.Println("enter the power which you want to calculate (a^b): ")
	fmt.Scan(&b)

	// pass these value in make this function a go routine
	powerchan<-a
	// Send the base value to the channel
	powerchan <- a

	// Calculate the power
	go calculatepower(b, powerchan)
	// fmt.Println("value of a^b is: ", result)
	close(powerchan)

}

