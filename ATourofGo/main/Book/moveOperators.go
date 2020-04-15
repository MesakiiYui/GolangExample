package main

import "fmt"
// 位移运算
func main() {
	//左移
	var a int = 1
	a = a << 10
	fmt.Print("左移后的结果为:")// 1 * 2^10
	fmt.Println(a)

	//右移
	var b int = 1024
	b = b >> 10
	fmt.Print("右移后的结果为:") // 1024 / 2^10
	fmt.Println(b)
}