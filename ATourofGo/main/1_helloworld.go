package main
// Programs start running in package main.-> must be "main" package
// Go通过包来管理命名空间

import (
	"fmt"
	"math/rand"
	"math"
)
// fmt 是标准输入输出包

// build a function
// e.g
// (x,y int) <=> (x int, y int)
func add(x , y int)(int, int, int){
	return x+y, x, y
}

func nakedReturn(sum int)(x ,y int){
	x = sum * 3
	y = sum * 2
	return
}
func main(){
	fmt.Println("hello, world. num is", rand.Intn(10))

	// In Go, a name is exported if it begins with a capital letter
	// e.g: math.Pi
	fmt.Println(math.Pi,)

	// use := to receive data 利用 := 来接受函数返回的数据，此处为add()函数
	res, a, b := add(1,3)
	fmt.Println("addFunction:",  res, a, b)

	// naked return
	// Naked return statements should be used only in short functions,
	// as with the example shown here. They can harm readability in longer functions.
	nakea, nakeb := nakedReturn(4)
	fmt.Println("naked return:", nakea, nakeb)


}
