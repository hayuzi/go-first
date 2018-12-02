package main

import "fmt"

func main() {

	// =================
	// 数组是具有固定长度且拥有领个或者多个相同数据类型元素的序列，由于数组的长度固定，所以在Go里面很少直接使用，
	// 而slice的长度可以增长和缩短，在很多场合下使用的更多

	var a [3]int // 3个整数的数组
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	// 输出索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// 仅仅输出元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	// 使用数组字面量根据一组值来初始化一个数组
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Printf("%d\n", q[2])
	fmt.Printf("%d\n", r[2]) // 0

	// 数组字面量中如果省略号出现在 数组长度的位置，那么数组的长度由初始化数组的元素个数决定
	a1 := [...]int{1, 2, 3}
	fmt.Println(a1[1])

	// 数组的长度是数组类型的一部分，所以 [3]int 与 [4]int是两种不同的数组类型，
	// 数组的长度必须是常量表达式，也就说，这个表达式的值在程序编译的时候就可以确定
	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB])

	// 数组的索引有时候也可以0值填充，例如
	a2 := [...]int{99: -1}
	fmt.Println(a2[90])

	// 如果一个数组的元素类型是可比较的，那么这个数组也是可以比较的
	a3 := [2]int{1, 2}
	a4 := [...]int{1, 2}
	a5 := [2]int{1, 3}
	fmt.Println(a3 == a4, a3 == a5, a4 == a5) // true false false
	a6 := [3]int{1, 2, 3}
	fmt.Println(a6[1])
	// fmt.Println(a3 == a6) // 编译错误，无法比较

}
