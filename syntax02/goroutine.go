package main

import (
	"fmt"
	"time"
)

func main() {
	// ===============
	// 1. Go中每一个并发执行的活动称为goroutine
	// 当程序执行的时候，一个go语句是在普通的函数或者方法调用前加上 go关键字前缀。
	// go语句使用函数在一个新创建的goroutine中调用。go语句本身的执行立即完成

	// 这段程序在主 goroutine 中创建了一个新的goroutine,
	// 在main函数执行完成时，所有的goroutine都暴力的直接终止
	go spinner(100 * time.Millisecond)
	const n = 5
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)

	// ================
	// 2. 通道 channel
	// 通道是可以让一个goroutine发送特定值到另一个goroutine的通信机制
	// 每一个通道是一个具体类型的导管叫做通道的元素类型
	// 一个有 int类型元素的通道写为 chan int
	ch := make(chan int) // ch 的类型是 chan int

	// 通道是一个使用make 创建的数据结构的引用.  当赋值或者作为参数传递到一个函数时, 复制的是引用.
	// 通道零值是  nil,  可以使用 == 比较, 当两者是同一个通道数据的引用时候, 通道相等
	// 通道有两个主要操作 发送和接收，两者统称为通信
	x := 1
	ch <- x // 发送语句
	// 此处程序被阻塞了。需要在另一个goroutine中操作接收
	x = <-ch // 赋值语句中的接收者表达式
	<-ch     // 接收语句, 丢弃结果

	// 通道可以关闭. 它设置一个标志位来指示值当前已经发送完毕，这个通道后面没有值了; 关闭后的发送操作将导致宕机.
	// 在一个已经关闭的通道上进行接收操作，将获取所有已经发送的值，直到通道为空;
	// 这时, 任何接收操作会立即完成, 同时获取到一个通道元素类型对应的零值.
	// 调用内置的close函数来关闭通道
	close(ch)

	// 使用简单make调用创建的通道叫做无缓冲通道, 但是 make还可以接收第二个可选参数，一个表示通道容量的整数
	// make(chan int) // 无缓冲通道
	// make(chan int, 0) // 无缓冲通道
	// make(chan int, 3) // 容量为3的缓冲通道

	// ================
	// 2.1 无缓冲通道 (同步通道)
	// 无缓冲通道上的发送会阻塞, 直到另外一个goroutine在另外的在对应的通道上执行接收操作,
	// 这时值传送完成, 两个 goroutine都可以继续执行.
	// 相反, 如果接收操作先执行, 接收方goroutine将阻塞，直到另一个goroutine在同一个通道上发送值
	// 参考 sample/netcat.go

	// 单向通道类型
	// 当一个通道被用做函数的行参时候，他几乎总是被有意的限制为只能接收或者只能发送
	// func(out chan<- int)
	// func(in <-chan int)

	// =================
	// 2.2 缓冲通道 (异步通道)
	// 缓冲通道有一个元素队列, 队列的最大长度在创建的时候被设置
	// 在队列未满的时候, 通道可以无阻塞的发送, 但如果队列已经满了,
	// 发送操作会阻塞所在的goroutine直到另一个goroutine对它戒行接收操作来留出可用空间.
	// 反过来，如果通道是空的, 执行接收操作的goroutine阻塞, 直到另一个goroutine在通道上发送数据
	ch2 := make(chan int, 5)
	m := cap(ch2) // 查看容量
	y := len(ch2) // 查看当前通道中有多少元素
	fmt.Println(m)
	fmt.Println(y)
	close(ch2)

	//  =============
	// 2.3 并行循环
	// 需要用到一个特殊类型的计数器 sync.WaitGroup, 它可以被多个goroutine安全的调用
	// 可以使用range迭代接收 通道的值

	// ================
	// 3. select 多路复用

}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
