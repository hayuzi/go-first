package main

import (
	"reflect"
	"fmt"
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






}
