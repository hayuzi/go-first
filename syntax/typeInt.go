package main

import "fmt"

func main() {

	// =================
	// Go的数据类型分为四大类
	// 1. 基础类型 basic type
	// 数字 字符串 布尔

	// 2. 聚合类型 aggregate type
	// 数组和结构体

	// 3. 引用类型 reference type
	// 指针pointer， 切片slice， map， 函数, 通道 channel

	// 4. 接口类型 interface type
	//

	/*
		==================
		1.1 整数
		有符号整数  int int8 int16 int32 int64
		无符号整数  uint uint8 uint16 uint32 uint64
		int uint的大小取决于是32位或者64位的硬件平台，但也不是绝对，还依赖于编译器的位数

		rune类型是int32类型的同义词，常常用于指明一个值是 Unicode码点
		同样 byte类型谁  uint8类型的同义词，强调一个值是原始数据，而非量值
		无符号整型 uintptr，其大小并不明确，但是足以完整存放指针，uintptr仅用于底层编程

		int uint uintptr 都有别于其他大小明确的相似类型的类型
		int 和 int32 是不同类型，int值如果要当作 int32使用，必须要显式转换，反之亦然
		有符号整型以补码表示

		通常我们很少使用无符号整数，除非特殊用途。比如数组长度虽然理论上可以用uint但是我们实际采用的还是有符号整数
		在我们使用for循环赋值迭代时候，如果len得倒的是非负们，则下面的代码将会造成死循环

		medals := []string{"gold", "silver", "broze"}
		for i := len(medals) - 1; i >= 0; i-- {
			fmt.Println(medals[i])
		}

	*/

	/**
	-------------------
	二元操作符优先级降序排列, 同层左结合律

	* 	/	% 	<<	>>	&	&^
	+	-	|	^
	==	|=	<	<=	>	>=
	&&
	||

	注意点: 	整数相除，结果还是整数，会舍弃小数部分
			整数运算超出数值范围会溢出

			& 	位运算 AND
			| 	位运算 OR
			^ 	位运算 XOR 作为二元运算时表示按位异或， 若是一元前缀则表示按位取反或取补
			&^	位清空（AND NOT） 表达式 z=x&^y表示若y的某位是1，则z的对应位等于0，否则就等于x的对应位的值
			<< 	左移
			>> 	右移 注意无符号右移以0补位，而有符号的则以符号位值补位置。
	*/

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	// %b二进制形式输出 08表示前面不足的部位0填充
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x>>1)

	/**
	===================
	重要： 不同类型的整数不能做运算，编译时候便会报错，如果需要预算，必须转换成同一类型
	但要注意不同类型的数值转换时候精度丢失的问题


	===================
	进制
	八进制 以 0 开头
	16进制 以 0x 或者 0X 开头
	*/

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // 4338 666  0666
	x1 := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x1) // 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF
	//  fmt.Printf中 %后面的[1]表示重复使用第一个操作数， #表示输出相应的前缀 0 0x 0X

	// =================
	// %c输出文字富豪，如果希望输入带有单引号则使用%q
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)

}
