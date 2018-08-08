package main

import "fmt"

func simpleIf(x int)(){
	if x < 5{
		fmt.Print(x, " less than 5")
		x++
	}
}

func main(){
	simpleIf(4)
}