package main

import (
	"reflect"
	"fmt"
	"io"
	"os"
)

func main() {

	// ========================
	// 反射.

	// 1. reflect.Type 和 reflect.Value
	// 反射功能由 reflect包提供。
	// Type表示Go语言的一个类型， 它是一个有很多方法的接口，这些方法可以用来识别类型以及透视类型的组成部分，
	// 比如一个结构的哥哥字段活着一个函数的各个参数。
	// reflect.Type 接口只有一个实现，即类型描述符，接口址值的动态类型也是类型描述符

	t := reflect.TypeOf(3)  // 一份 reflect.Type
	fmt.Println(t.String()) // "int"
	fmt.Println(t)          // "int"

	// 因为 reflect.TypeOf 返回一个接口值对应的动态类型，所以它返回的总是具体类型（ 而不是接口类型 ）

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // 是 *os.Files 而不是 io.Writer

	// reflect.Value 可以包含一个任意类型的值
	// reflect.ValueOf函数接收任意的interface{} 并将接口的动态值以reflect.Value的形式返回
	// 与reflect.TypeOf类似, reflect.ValueOf的返回值也都是具体值，不过 reflect.Value也可以包含一个接口值
	v := reflect.ValueOf(3) // 一个reflect.Value
	fmt.Println(v)          // "3"
	fmt.Printf("%v\n", v)   // 3
	fmt.Println(v.String()) // 注意： "<int Value>"
	// reflect.Value也满足fmt.Stringer, 但除非Value包含的是一个字符串，否则 String方法的结果仅仅暴露了类，
	// 通常需要使用 fmt包的%v 功能，它对 reflect.Value会进行特殊处理

	// 调用 value的Type方法会把它的类型以 reflect.Type方式返回
	t2 := v.Type()
	fmt.Println(t2.String()) // "int"

	// reflect.ValueOf的逆操作是 reflect.Value.Interface 方法。它返回一个 interface{} 接口值，
	// 与reflect.Value包含同一个具体值

	v2 := reflect.ValueOf(3) // a reflect.Value
	x := v2.Interface() // an interface{}
	i := x.(int)		// an int
	fmt.Printf("%d\n", i)


	// Kind 方法可以区分不通的类型。 类型分支只有少数几种：
	// 基础类型 Bool,string以及各种数字类型；
	// 聚合类型Array和Struct ;
	// 引用类型 Chan、func、Ptr、Slice和 Map、接口类型Interface；
	// 最后还有Invalid类型，表示还没有任何值(reflect.Value 的零值就属于 Invalid类型)
	k := v2.Kind() // reflect.Int
	fmt.Println(k == reflect.Int)



	// ====================
	// 使用reflect.Value来设置值
	// reflect.ValueOf(&x).Elem() 可以获取任意变量 x 可寻址的 Value 值。
	// 可以通过变量 CanAddr() 方法来询问 reflect.Value 变量是否可以寻址
	x3 := 2
	a3 := reflect.ValueOf(2)
	b3 := reflect.ValueOf(x)
	c3 := reflect.ValueOf(&x)
	d3 := c3.Elem()
	// a3 里面的值不可以寻址. b3 也是如此.
	// c3 也不可以寻址。但是 d3是通过 对c3中的指针提领得来的，所以它可以寻址
	fmt.Println(x3)
	fmt.Println(a3.CanAddr())
	fmt.Println(b3.CanAddr())
	fmt.Println(c3.CanAddr())
	fmt.Println(d3.CanAddr())

	// ===============
	// 从一个可以寻址的 reflect.Value()获取变量需要三步
	// 第一 调用 Addr(), 返回一个value, 其中包含一个指向变量的指针
	// 第二 在这个Value 上调用Interface(), 会返回一个包含这个指针的interface{}值
	// 第三 类型断言把接口内容转换为普通指针
	// 之后可以通过这个指针来更新变量
	x4 := 2
	d4 := reflect.ValueOf(&x4).Elem()
	px4 := d4.Addr().Interface().(*int) // px := &x ( 类型断言将接口内容转换为指针 )
	*px4 = 3
	fmt.Println(x4)

	// 或者可以直接通过可以寻址的 reflect.Value来更新变量，不用通过指针，而是直接调用 reflect.Value.Set 方法
	d4.Set(reflect.ValueOf(4))	// 如果类型对不上，则崩溃
	fmt.Println(x4)

	// 由于类型的限制，reflect.Value提供了一些兼容设置函数
	// SetString SetInt()
	// 但是 在指向 interface{} 变量的 reflect.Value 上调用 SetInt会崩溃 (尽管使用Set就没有问题)


	// =================
	// reflect.Type 和 reflect.Value 都有一个叫做 Method 的方法。
	// 每个 t.Method(i) (从reflect.Type 调用) 都会返回一个reflect.Method类型的实例,
	// 这个结构类型描述了这个方法的名称和类型












}
