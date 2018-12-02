package main

import (
	"fmt"
	"math/cmplx"
)

func main() {

	// =================
	// 复数
	// Go具备两种大小的复数complex64和complex128，二者分别由float32和float64组成
	// 内置的complex函数根据戈丁的实部和虚部创建复数，而内置的 real函数和imag函数则分别提取复数的实部和虚部

	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))

	// 源码中方，如果在浮点数或者十进制整数后面紧接着写自负i，如 3.14i 或者 2i, 它就变成了一个虚数，表示一个实部为0的虚数
	// 复数声明简写
	// x := 1 + 2i
	// y := 3 + 4i

	// == 判断复数是否相等的时候，必须实部与虚部都相等才为true
	// math.cmplx 包提供了复数运算所需要的库函数

	fmt.Println(cmplx.Sqrt(-1))

}
