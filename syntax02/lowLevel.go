package main

import (
	"unsafe"
	"fmt"
)

func main() {
	// ==================
	// 1. unsafe
	// unsafe 不保证新版本的兼容
	// unsafe.Sizeof() // 报告传递给它的参数在内存中占用的字节长度, 这个参数可以是任何类型的表达式，不会计算表达式
	fmt.Println(unsafe.Sizeof(float64(0))) // "8"

	//类型									大小
	//bool									1个字节
	//intN, uintN, floatN, complexN			N/8个字节(例如float64是8个字节)
	//int, uint, uintptr					1个机器字
	//*T									1个机器字
	//string								2个机器字(data,len)
	//[]T									3个机器字(data,len,cap)
	//map									1个机器字
	//func									1个机器字
	//chan									1个机器字
	//interface								2个机器字(type,value)


	// unsafe.Alignof() // 报告它参数类型所要求的对齐方式, 它的参数可以是任意类型的表达式, 并且返回一个常量。
	// 典型地，布尔类型和数值类型对齐到她们的长度(最大8字节)


	// unsafe.Offsetof() // 计算成员f相对于结构体x其实地址的偏移值, 如果有内存空位，也计算在内，
	// 该函数的操作数必须是和一个成员选择器 x.f

	// 基于不同的平台(32位/64位)，值也不同
	var x struct{
		a bool
		b int16
		c []int
	}
	x.a = true
	x.b = 12

	fmt.Println("===============")

	fmt.Println(unsafe.Sizeof(x))
	fmt.Println(unsafe.Alignof(x))

	fmt.Println(unsafe.Sizeof(x.a))
	fmt.Println(unsafe.Alignof(x.a))
	fmt.Println(unsafe.Offsetof(x.a))

	fmt.Println(unsafe.Sizeof(x.b))
	fmt.Println(unsafe.Alignof(x.b))
	fmt.Println(unsafe.Offsetof(x.b))

	fmt.Println(unsafe.Sizeof(x.c))
	fmt.Println(unsafe.Alignof(x.c))
	fmt.Println(unsafe.Offsetof(x.c))


	// unsafe.Pointer
	// unsafe.Pointer 是一种特殊类型的指针，它可以存储任何变量的地址。
	// 当然，我们无法简介地通过一个unsafe.Pointer 变量来使用 *p , 因为我们不知道这个表达式的具体类型

	// unsafe.Pointer 类型的指针是可以比较的并且可以和 nil做比较， nil是指针类型的 零值

	// 普通指针与 unsafe.Pointer可以互相转换，而且从 unsafe.Pointer转回普通指针的时候，可以不必和原来的类型 *T相同






	// ==================
	// cgo 工具
}
