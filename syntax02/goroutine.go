package main

import (
	"time"
	"fmt"
)

func main(){
	// ===============
	// 1. Go中每一个并发执行的活动称为goroutine
	// 当程序执行的时候，一个go语句是在普通的函数或者方法调用前加上 go关键字前缀。
	// go语句使用函数在一个新创建的goroutine中调用。go语句本身的执行立即完成

	// 这段程序在主 goroutine 中创建了一个新的goroutine,
	// 在main函数执行完成时，所有的goroutine都暴力的直接终止
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner (delay time.Duration) {
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