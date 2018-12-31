package main

import (
	"sync"
	"fmt"
	"time"
)

var (
	mu sync.Mutex
	rwMu sync.RWMutex
	loadOnce sync.Once
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

	// =================
	// 2. 读写互斥锁 sync.RWMutex
	// 2.1 多读单写锁 顾名思义，允许多个goroutine同时读取，但同时只允许一个goroutine写入
	go func() {
		rwMu.RLock()			// 读锁, 该读锁是共享的。
		defer rwMu.RUnlock()
	}()

	// ================
	// 3. 内存同步
	// 现在计算机系统的多个处理器，会维护一个内存的本地缓存，为了提高效率，对内存的写入是缓存在每个处理器中的，
	// 只在必要时才刷回内存，甚至刷会内存的顺序都可能与goroutine的写入顺序不一致。
	// 像通道通信或者互斥锁操作这样的同步原语都会导致处理器把积累的写操作刷回内存并提交

	// 所以在可能的情况下， 把变量限制在单个goroutine中. 其他情况下使用互斥锁


	// ==============
	// 4. 延迟初始化 sync.Once
	// 从概念上来讲， Once包含一个布尔变量和一个互斥量，布尔变量记录初始化是否已经完成，
	// 互斥量则负责保护这个布尔变量和客户端的数据结构，Once的唯一方法Do以初始化函数作为它的参数
	var b int
	go func() {
		loadOnce.Do(func() {
			b = 1
		})
	}()
	// 每次调用 Do()时候，会先锁定互斥量并检查里面的布尔变量。第一次调用时候，布尔变量为假，
	// Do调用初始化函数然后把变量设置为真，后续的低矮用相当于空操作。


	// ======================
	// 5. 竞态检测器
	// 简答的把 -race 参数加入到  go build、 go run、go test 命令李即可使用该功能


	// ======================
	// 6.示例：并发非阻塞缓存 参考 sample/memo.go






}