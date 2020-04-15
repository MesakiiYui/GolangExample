package main

import "fmt"

func main() {
	// 创建
	ma := map[string] int{"a":1, "b":2}
	fmt.Println(ma["b"])

	// 利用make创建
	//make(map[K]T)// map的容量使用默认值
	//make(map[K]T, len)// map的容量使用给定的len值
	mp1 := make(map[int]string)
	mp2 := make(map[int]string, 19)
	mp1[0] = "tom"
	mp2[0] = "bob"
	// map支持的操作
	// 1.单值访问
	fmt.Println(mp1[0])
	// 2.删除操作
	mp2[1] = "kali"
	fmt.Println(mp2)
	delete(mp2, 1)
	fmt.Println(mp2)
	// 3.返回键值对数量
	fmt.Println(len(mp2))

	// 注:Go内置的map并不是并发安全的，若需要并发安全可使用标准包sync中的map
	// 如果value是一个struct或者其他复合数据结构，需要修改其中一个value中的某个子项的值，需要对其进行整体赋值，而不是单个赋值


}
