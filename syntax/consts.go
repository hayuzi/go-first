package main

import (
	"fmt"
	"math"
)

const constNum1 = 1 // 声明常量，全局， 常量声明以后不能再次声明，并且不能被重新赋值

func main() {

	// ===================
	// 常量声明
	// 所有常量本质上都属于基本类型 布尔型，字符串或数字

	const constNum2 = 2 // 局部常量
	const constNum3 int64 = 3
	// 也可以使用如下方式声明
	//const (
	//	constNum3	= 3
	//	constNum4	= 4
	//)

	// 声明了一个局部变量必须在相同的代码块中使用它 ( 全局变量例外 )
	// 说明 g是一个指针，表示变量a的内存地址，如果要获取g所指向的值，请使用 *g来获取 g所指向的值
	fmt.Println(constNum1, constNum2, constNum3)

	// 声明一组常量的时候，等号右侧的表达式可以省略，这意味着复用前一项的表达式及其类型
	const (
		constNum6 = 1
		constNum7
		constNum8 = 2
		constNum9
	)
	fmt.Println(constNum6, constNum7, constNum8, constNum9)

	// ===============
	// 常量生成器 iota， 它创建一系列相关值， 而不是逐个值显式写出 常量声明中，iota从0开始取值，逐项加 1
	type Weekday = int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	type Flags uint
	const (
		FlagUp           Flags = 1 << iota // 向上
		FlagBroadcast                      // 支持广播访问
		FlagLoopback                       // 是回环接口
		FlagPointToPoint                   //属于点对点链路
		FlagMulticast                      // 支持多路广播访问
	)
	fmt.Println(FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)

	// ==================
	// 无类型常量 , 只有常量才可以是无类型的
	// 从属类型待定的常量有6种，分别是
	// 无类型布尔 无类型整数 无类型文字符号 无类型浮点数 无类型复数 无类型字符串

	// 无类型常量不仅能暂时维持更高的精度，更可以写进更多的表达式而无需转换类型
	// 例如浮点型常量 math.Pi 可用于任何需要浮点值或者复数的地方

	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi
	fmt.Println(x, y, z)

}
