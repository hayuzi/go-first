package main

import "fmt"

var ( // 这种因式分解关键字的写法一般用于声明全局变量
	va int    // 默认0
	vb bool   // 默认false
	vc string // 默认空字符串
)

func main() {
	// ====================
	// == 变量声明 ==
	// 变量声明有三种方式 (第一种与第二种方法可以用于声明全局变量, 第三种只能用来声明局部变量, 但是比较常用语简略)
	// 第一种，指定变量类型，声明后若不赋值，使用默认值（ 零值 ）
	var a int32 = 1 // 格式 var v_name v_type = value
	var d uint      // 使用 默认值为 0
	var f float32 = 1.2

	// 第二种，根据值自行判定变量类型。
	var b = 10

	// 第三种, 省略var, 简短形式, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误。
	c := 100

	// ====================
	// == 多变量声明 ==
	// 第一种, 两步赋值. 格式如下
	// var vname1, vname2, vname3 type
	// vname1, vname2, vname3 = v1, v2, v3
	var h, i int
	_, i = 1, 3 // *** 空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。

	// 第二种. var vname1, vname2, vname3 = v1, v2, v3  //和python很像,不需要显示声明类型，自动推断
	var u, v, w = "String test", 1, 2

	// 第三种  这种因式分解关键字的写法一般用于声明全局变量, 参考开头方法外的声明
	// var (
	// 	  vname1 v_type1
	//	  vname2 v_type2
	// )

	// 第四种 不带声明格式
	x, y, z := 1, 2, false // 这种不带声明格式的只能在函数体中出现，且二次声明的时候必须至少要声明一个新变
	// x, y := 1, 2 	// 上面已经声明过 xy，此处不能通过
	// x, oy :=  1, 2 	// 必须有新变量才能通过，如 oy

	// ====================
	// 值类型和引用类型, 引用是引用变量的内存地址
	// 当使用等号 = 将一个变量的值赋值给另一个变量时，如  j = i，实际上是在内存中将 i 的值进行了拷贝
	// 而 j = &i   	这样的方式 j 获取的是 i 的内存地址
	var g = &a // 获取内存地址

	// 使用 new函数来创建变量，表达式new(T)创建一个未命名的变量。初始化化为T类型的零值
	j := new(int) // 此时j是一个指针，指向未命名的int变量
	*j = 2

	// ===================
	// 注意事项
	// 不允许重复定义变量, 变量定以后不能再使用初始化变量的方式来赋值, 只能直接使用  j = value 这样的格式给变量赋新值
	a = 32

	/**
	=====================
	关于变量名称，有如下要点：
	1. 字母或者下划线开头后面可以跟任意数量的自负，数字和下划线
	2. 以下程序关键字不可以使用作为变量名称：
		break 		default 	func 		interface 	select
		case		defer		go			map			struct
		chan		else		goto		package		switch
		const		fallthrough	if			range		type
		continue	for			import		return		var
	3. 另外还有三十几个内置的预声明的常量/类型/函数
		常量		true	false	iota	nil
		类型		int		int8	int16	int32	int64
				uint	uint8	uint16	uint32	uint64	uintptr
				float32	float64	complex128	complex64
				bool	byte	rune	string	error
		函数		make	len		cap		new		append
				copy	close	delete
				complex	real	imag
				panic	recover


	*/

	// 声明了一个局部变量必须在相同的代码块中使用它 ( 全局变量例外 )
	// 说明 g是一个指针，表示变量a的内存地址，如果要获取g所指向的值，请使用 *g来获取 g所指向的值
	fmt.Println(va, vb, vc, a, b, c, d, f, g, *g, h, i, j, u, v, w, x, y, z)

}
