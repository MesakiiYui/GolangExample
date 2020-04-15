package main

import "fmt"

func main() {
	type User struct {
		Name string
		Age int
	}
	type Student struct{
		*User
		Number int
	}

	// 并不推荐下面这个方式，因为struct增减字段，上面这个a的赋值就会报错
	a := User{"Tom", 18}
	fmt.Println(a)

	// 推荐使用下面的方式
	b := &User{
		Name:"Rook",
		Age:30,
	}
	fmt.Println(b)
	c := Student{
		User: b,
		Number: 1,
	}
	fmt.Println(c)

}
