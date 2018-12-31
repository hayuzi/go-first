package main

import "sync"

// Memo 缓存了调用func的结果
type Memo struct {
	f     Func
	mu    sync.Mutex // 保护cache
	cache map[string]*entry
}

// func是用于记忆的函数类型
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // 当res准备好后关闭该通道
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// 对key的第一次访问， 这个goroutine负责计算数据和广播数据
		// 已经准备完毕的消息
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready) // 广播数据已准备完毕的消息
	} else {
		memo.mu.Unlock()

		<-e.ready // 等待数据准备完毕
	}
	return e.res.value, e.res.err
}
