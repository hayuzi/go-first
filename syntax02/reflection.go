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




}
