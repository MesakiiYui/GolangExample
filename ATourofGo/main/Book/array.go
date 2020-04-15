package main

import "fmt"

func main() {
	// 数组的声明
	var arr1 [2]int // 声明一个有两个整形的数组，默认值为0，很少使用
	fmt.Println("arr1", arr1)
	//初始化并赋值
	arr2 := [...] float64{1.1, 2.2, 3.3}
	fmt.Println("arr2", arr2)
	arr3 := [3]int{1,2,3}
	fmt.Println("arr3", arr3)
	arr4 := [...]int{1,2,3}// 虽然不指定初始长度，但是由后面的初始化列表的长度来确定其长度
	fmt.Println("arr4", arr4)
	arr5 := [3]int{1:2,2:3}// 指定总长度，并通过索引值进行初始化，没有初始化元素值时使用类型默认值 {1的位置是2，2的位置是3 }
	fmt.Print("arr5", arr5)

	// 数组长度固定，不可追加
	// 数组是指类型的，数组赋值或者是函数赋值都是值拷贝
	// 数组长度是数组类型的组成部分，[10]int 和[20]int表示不同的类型
	// 可以根据数组创建切片
}
