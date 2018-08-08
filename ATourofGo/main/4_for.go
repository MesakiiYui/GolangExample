package main

import "fmt"

func simpleLoop(){
	var sum int = 10
	// The init statement will often be a short variable declaration => var i int = 1  will not work
	for i := 0; i < sum; i++{
		fmt.Print(i, ",")

	}
	fmt.Println()
	for ; sum < 12;{
		sum ++
		fmt.Print(sum, ",")

	}
	fmt.Println()
}

func main(){
	simpleLoop()
}