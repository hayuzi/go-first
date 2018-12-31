package memo2

import (
	"math/rand"
	"time"
	"fmt"
	"log"
	"sync"
)

func TestMemoGet() {
	// 主程序中初始化
	rand.Seed(time.Now().UnixNano())
	m := New(initCacheValue)
	var n sync.WaitGroup
	var keys = []string{"test", "test1", "test", "test", "test3", "test5"}
	for _, s := range keys {
		n.Add(1)
		go func(s string) {
			value, err := m.Get(s)
			if err != nil {
				log.Print(err)
			}
			fmt.Println(value)
			n.Done()
		}(s)
	}
	n.Wait()
	m.Close()
}

func initCacheValue(key string) (interface{}, error) {
	x := rand.Intn(10000)
	return fmt.Sprintf("key:%s: %d", key, x), nil
}