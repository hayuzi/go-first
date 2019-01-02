// 这是memo包的测试文件
package memo

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

// 功能测试函数
func TestMemo_Get(t *testing.T) {
	// 主程序中初始化
	rand.Seed(time.Now().UnixNano())
	m := New(initCacheValue)
	key := "test"
	value, err := m.Get(key)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(value)
}

// 基准测试函数
func BenchmarkMemo_Get(b *testing.B) {

}

// 示例函数
func ExampleMemo_Get() {

}
