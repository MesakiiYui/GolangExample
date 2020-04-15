package main

import "fmt"

func main() {
	 var a  = 11
	 p := &a
	 fmt.Println(p)

	 // 利用指针访问结构体
	 type User struct {
	 	name string
	 	age int
	 }
	 bob := User{
	 	name : "bob",
	 	age : 16,
	 }
	 pBob := &bob
	 fmt.Println(pBob.name)
	 fmt.Println(pBob.age)

	 // Go 禁止指针运算
	 // 不存在 p++

	 var mySum = sum(1,2)
	 fmt.Println(mySum)
}

// Go编译器使用“栈逃逸”机制将这种局部变量的空间分配在堆上
func sum (a,b int) *int {
	sum := a + b
	return &sum

}
