package main

import "fmt"

// 该处type即为类型声明，类型声明通常出现在包级别，这类命名的类型在整个包中可见 如果名字是导出的（开头使用大些字母），其他包也可以访问它
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {

	f := CToF(FreezingC)

	c := FToC(1000.0)

	fmt.Println(f, c)

}

func CToF(c Celsius) Fahrenheit {
	// 对于每个类型 T, 都有一个类型转换操作T(x)将 x的值转换为类型T，比如以下这个Fahrenheit()
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
