package main

import "fmt"

func main() {
	// bool 布尔型 如果不初始化，则默认为false
	var ok bool
	ok2 := false
	fmt.Println(ok, ok2)

	// 整型
	var a int = 1;
	fmt.Println(a)

	// 浮点型 浮点型字面量被自动类型推断为float64类型
	// 由于计算机难以存储，浮点型的比较应该用math库而不是 == 或 !=
	var b = 10.00
	fmt.Println(b)

	// 字符串
	var strA = "hello"
	fmt.Println(strA, strA[0], strA[1], strA[2:])
	var strB = "buildyourdreams"
	for i := 0; i < len(strB); i++{
		fmt.Print(strB[i])
	}
	for i, v := range strB{
		fmt.Println(i,v)
	}

}
