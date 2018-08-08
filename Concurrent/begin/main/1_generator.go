package main

import (
	"math/rand"
	"fmt"
)

// 这是非并发的场景，调用函数，返回结果，如果函数的内部处理非常耗时，该函数的调用者就会挂起
func rand_generator_1() int{ // 返回的是Int
	return rand.Int()
}

// 创建一个协程，专门用于执行rand.Int
func rand_generator_2() chan int{ //返回的是通道
	// 创建通道
	out := make(chan int)
	// 创建协程
	go func(){ // go func表示func这个函数会是以协程的方式运行。这样就可以提供程序的并发处理能力
		for{
			// 向通道内写入数据，如果无人读取就会等待
			out <- rand.Int()
			fmt.Println("out", out)
		}
	}()
	return out

}

// 上面的这段函数就可以并发执行了rand.Int()。有一点值得注意到函数的返回可以理解为一个“服务”。
// 但我们需要获取随机数据 时候，可以随时向这个服务取用，他已经为我们准备好了相应的数据，无需等待，随要随到。
// 如果我们调用这个服务不是很频繁，一个协程足够满足我们的需求了。但如果我们需要大量访问，怎么办？
//
// 我们可以用下面介绍的多路复用技术，启动若干生成器，再将其整合成一个大的服务。
// 调用生成器，可以返回一个“服务”。可以用在持续获取数据的场合。用途很广泛，读取数据，生成ID，甚至定时器。这是一种非常简洁的思路，将程序并发化。

// 多路复用技术， 多路复用是让一次处理多个队列的技术。
func rand_generator_3() chan int{
	// 创建两个随机数生成服务
	rand_generator_1 := rand_generator_2()
	rand_generator_2 := rand_generator_2()

	// 创建通道
	out := make(chan int)

	// 创建协程
	go func(){
		for{
			//读取生成器1中的数据，整合
			out <- <- rand_generator_1
		}
	}()
	go func(){
		for{
			//读取生成器2中的数据，整合
			out <- <- rand_generator_2
		}
	}()
	return out
}

// 上面是使用了多路复用技术的高并发版的随机数生成器。通过整合两个随机数生成器，这个版本的能力是刚才的两倍。
// 虽然协程可以大量创建，但是众多协程还是会争抢输出的通道。Go语言提供了Select关键字来解决，各家也有各家窍门。
// 加大输出通道的缓冲大小是个通用的解决方法。

// 多路复用技术可以用来整合多个通道。提升性能和操作的便捷。配合其他的模式使用有很大的威力。


func main(){
	// 生成一个随机数作为一个服务
	rand_service_handler := rand_generator_2()
	rand_service_handler2 := rand_generator_3()
	// 从服务中读取随机数并打印
	fmt.Println("rand_service_handler", rand_service_handler)
	fmt.Println("rand_service_handler2", rand_service_handler2)


}
