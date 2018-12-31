package main

import (
	_ "net"            // 空导入	(空倒入主要是为了替代不在该文件中使用，但需要用到初始化特性的包)
	mtRand "math/rand" // 别名倒入(替代名字只影响当前文件)
	"time"

	// "../sample2/memo"	// 本地目录包导入(强烈不建议使用)
)

func random() int {
	mtRand.Seed(time.Now().UnixNano())
	return mtRand.Intn(1000)
}
