package main

import "fmt"

const constNum1 = 1 // 声明常量，全局， 常量声明以后不能再次声明，并且不能被重新赋值

func main() {

	// ===================
	// 常量声明
	const constNum2 = 2 // 局部常量
	// 也可以使用如下方式声明
	//const (
	//	constNum3	= 3
	//	constNum4	= 4
	//)

	// 声明了一个局部变量必须在相同的代码块中使用它 ( 全局变量例外 )
	// 说明 g是一个指针，表示变量a的内存地址，如果要获取g所指向的值，请使用 *g来获取 g所指向的值
	fmt.Println(constNum1, constNum2)

}
