package main

import (
	"./memo2"	// 测试临时导入
	"./memo"
	"fmt"
)

func main() {
	memo2.TestMemoGet()
	fmt.Println("================")
	memo.TestMemoGet()
}
