package main

import "fmt"

// 常量的声明方式
func main(){
	// 类似枚举的iota
	// iota -> 预声明标识符iota用在常量声明中，初始值为0.一组多个常量同时声明时，其值逐行增加，其可看作自增的初始变量，专门用来初始化变量
	const (
		c0 = iota // 0
		c1 = iota // 1
		c2 = iota // 2
	)
	fmt.Println(c0)
	fmt.Println(c1)
	fmt.Println(c2)

	// 简写模式
	const (
		d0 = iota // 0
		d1	// 1
		d2	// 2
	)
	fmt.Println(d0)
	fmt.Println(d1)
	fmt.Println(d2)

	// 加入位移运算
	const (
		a = 1 << iota // iota = 0, a = 1*2^0 = 1
		b = 1 << iota // iota = 1, b = 1*2^1 = 2
		c = 3
		d = 1 << iota // iota = 3, d = 1*2^3 = 8
	)
	fmt.Println(a, b, c, d)

	const x = iota
	const y = iota
	fmt.Println( x, y)
	// 分开的const，iota不会叠加


}