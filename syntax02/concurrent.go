package main

import (
	"sync"
	"fmt"
	"time"
)

var (
	mu sync.Mutex
	rwMu sync.RWMutex
)

func main() {

	// =========================
	// 并发情况下形成竞态, 会对业务造成冲击, 所以引入互斥机制

	// ==============
	// 1. 互斥锁 sync.Mutex
	// 当一个goroutine已经取得了互斥锁之后，其他goroutine再次调用该互斥量的 Lock方法会阻塞，
	// 直到那个获取到锁的 goroutine调用 Unlock来释放锁之后，程序才能继续往下执行
	// 如果一直调用而不释放会造成死锁.
	a := 1
	go func() {
		mu.Lock()
		defer mu.Unlock()
		time.Sleep(1 * time.Second)
		a = 5
	}()

	go func() {
		mu.Lock()
		defer mu.Unlock()
		time.Sleep(1 * time.Second)
		a = 6
	}()

	for i := 0; i < 5 ; i++ {
		time.Sleep(1*time.Second)
		fmt.Println(a)
	}

	// 2. 读写互斥锁
	// 2.1 多读单写锁 顾名思义，允许多个goroutine同时读取，但同时只允许一个goroutine写入
	go func() {
		rwMu.RLock()			// 读锁, 该读锁是共享的。
		defer rwMu.RUnlock()
	}()




}