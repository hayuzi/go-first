package main

import (
	"fmt"
	"os"
)

func main() {

	// go的命令行参数
	//os包提供一些函数和变量，以与平台无关的方式和平台和操作系统打交道
	// os.Args 是一个字符串slice（切片）
	// slice是go中的基础概念
	var s, s2, sep string

	// for是go中唯一的循环语句
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	// form循环的三个组成部分每一个都可以省略
	// for initialization; condition; post {
	// 		xxxxxx.
	// }

	// for {} 是一个无限循环
	var condition = true
	for condition {
		s2 = "test"
		condition = false
	}

	// for循环的另一种形式，在 字符串或者 slice数据上迭代
	var s3, sep3 string
	for _, arg := range os.Args[1:] {
		s3 += sep3 + arg
		sep3 = "_"
	}

	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(s3)

}
