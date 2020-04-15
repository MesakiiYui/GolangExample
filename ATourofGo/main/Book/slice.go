package main

import (
	"fmt"
)

func main() {
	// len 切片的元素数量
	// cap
	// 切片的创建
	// 由数组创建
	var arr1 = [...]int{1,2,3,4,5,6}
	s1 := arr1[0:3]
	fmt.Println(s1)

	// 利用内置函数make创建切片
	// 注意，由make创建的切片各元素被默认初始化为切片元素类型的零值
	// len=10,cap=10
	a := make([]int, 10) //len = cap = 10
	fmt.Println(a)
	b := make([]int, 10, 15) //len=10,cap=15
	fmt.Println(b)

	// 注：直接声明切片类型变量是没有意义的
	// 例：var a []int

	// 总结
	// 切片支持的操作
	// 内置函数len()返回切片的长度
	// 内置函数cap()返回切片底层数组容量
	// 内置函数append()对切片追加元素，底层元素的数量，即cap会进行扩展
	// 内置函数copy()用于复制一个切片
	
}
