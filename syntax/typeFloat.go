package main

import (
	"fmt"
	"math"
)

func main() {

	// =================
	/**
	浮点数分为 float32 和float64
	最大的浮点值大约为 3.4e38 和 1.8e308
	最小的正浮点值大约为 1.4e-45 和 4.9e-324

	十进制下，float32的有效数字大约是6位， float64的有效数字大约是15位 ，绝大多数情况下应该优先选用float64
	float32的运算会迅速积累误差值， 另外，float32能精确表示的正整数范围有限
	*/
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	var f float32 = 16777216
	fmt.Println(f)
	fmt.Println(f + 1)
	fmt.Println(f == f+1)

	// 在源码中浮点数可以写成小数，
	const e = 2.71828 // 近似值

	// 如果是 0.123可以写成 .123
	// 如果是 123.00 也可以写成  123.
	// 非常大的数值最好写成科学计数法
	const Avogadro = 6.02214129e23
	const Planck = 6.62606957e-34

	fmt.Printf("%g\n", e)
	fmt.Printf("%g\n", Avogadro)
	fmt.Printf("%g\n", Planck)
	fmt.Printf("%e\n", Planck)

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e的x次方 = %8.3f\n", x, math.Exp(float64(x)))
	}

	// 正无穷 负无穷 以及 NaN
	// math.isNaN判断是否非数值， 每一个NaN总是不相等，并且大小比较也都是 false，无任何意义
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // 0 -0 +Inf -Inf NaN

}
