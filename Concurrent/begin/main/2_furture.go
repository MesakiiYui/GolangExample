package main

import "fmt"

// 构造一个查询结构体

type query struct{
	sql chan string
	result chan string
}

// 执行query
func execQuery(q query){
	// 启动协程
	go func(){
		// 获取输入
		sql:= <-q.sql
		// 访问数据库，输出结果通道
		q.result <- "get" + sql
	}()
}

func main(){
	// 初始化query
	q := query{make(chan string, 1), make(chan string, 2)}
	// 执行query， 注意执行的时候无需准备参数
	execQuery(q)

	// 准备参数
	q.sql <- "select * from table;"
	fmt.Println(<-q.result)
}

// 调用一个函数的时候，往往是参数已经准备好了。调用协程的时候也同样如此。
// 但是如果我们将传入的参数设为通道，这样我们就可以在不准备好参数的情况下调用函数。
// 这样的设计可以提供很大的自由度和并发度。函数调用和函数参数准备这两个过程可以完全解耦

// Furture和生成器的区别在于，Furture返回一个结果，而生成器可以重复调用。
// 还有一个值得注意的地方，就是将参数Channel和结果Channel定义在一个结构体里面作为参数，而不是返回结果Channel。这样做可以增加聚合度，好处就是可以和多路复用技术结合起来使用。

// Furture技术可以和各个其他技术组合起来用。可以通过多路复用技术，监听多个结果Channel，当有结果后，自动返回。
// 也可以和生成器组合使用，生成器不断生产数据，Furture技术逐个处理数据。Furture技术自身还可以首尾相连，形成一个并发的pipe filter。
// 这个pipe filter可以用于读写数据流，操作数据流。

// Future是一个非常强大的技术手段。可以在调用的时候不关心数据是否准备好，返回值是否计算好的问题。让程序中的组件在准备好数据的时候自动跑起来。