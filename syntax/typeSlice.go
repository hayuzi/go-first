package main

import "fmt"

func main() {

	// =================
	// 1. slice表示一个拥有相同类型元素的可变长度的序列，slice通常写成 []T, 其中元素的类型都是T，它看上去像是没有长度的数组类型。
	// 数组和slice紧密关联。 slice是一种轻量级的数据结构，可以用来访问数组的部分或者全部元素，而这个数组称为slice的底层数组。

	// 2. slice有三个属性，指针长度和容量
	// 指针指向数组的第一个可以从 slice中访问的元素，这个元素不一定是数组的第一个元素。
	// 长度是指slice中的元素个数，它不能超过 slice的容量。
	// 容量的大料通常是从slice的起始元素到底层数组的最后哦一个元素间的个数。 Go内置的函数 len和 cap用来发挥 slice的长度和容量

	// 3. 一个底层数组可以对应多个 slice，这些slice可以引用那个数组的任何位置。彼此间的元素还可重叠。
	// 如果slice应用超过了被引用的对象容量, 即cap(s)那么会导致宕机,
	// 但是如果slice的引用超出了被引用对象的长度
	// 即 len(s), 那么最终slice会比原来slice长

	a := [...]int{0, 1, 2, 3, 4, 5} // 初始化数组
	reverse(a[:])
	fmt.Println(a)

	// 4. 初始化slice使用 [],注意与初始化数组的区别
	s := []int{0, 1, 2, 3, 4, 5} // 初始化slice
	fmt.Println(len(s), cap(s))

	// 5. 和数组不同的是, slice无法做比较，因此不能 == 来测试两个slice是否有相同的元素
	// 标准库提供了高度优化的 函数 bytes.Equal 来比较两个字节slice ([]byte), 但是对于其他类型的 slice我们必须自己写函数比较
	ss1 := []string{"test", "hello"}
	ss2 := []string{"test", "hello"}
	fmt.Println(equal(ss1, ss2))

	// 6. slice类型的零值是 nil
	// 值为 nil的slice没有对应的底层数组
	// 值为 nil的 slice长度和容量都是零， 但是也有非 nil的slice长度和容量是 0
	var s1 []int // len(s) == 0, s == nil
	fmt.Println(len(s1), cap(s1), s1 == nil)
	s1 = nil // len(s) == 0, s == nil
	fmt.Println(len(s1), cap(s1), s1 == nil)
	s1 = []int(nil) // len(s) == 0, s == nil
	fmt.Println(len(s1), cap(s1), s1 == nil)
	s1 = []int{} // len(s) == 0, s != nil
	fmt.Println(len(s1), cap(s1), s1 == nil)

	// 7. 内置函数 make 可以创建一个具有制定给元素类型/长度和容量的 slice
	// make([]T, len)
	// make([]T, len, cap)
	// make([]T, cap)[:len]
	s2 := make([]int, 2)
	s3 := make([]int, 3, 6)
	s4 := make([]int, 6)[:3]
	fmt.Println(len(s2), cap(s2), s2)
	fmt.Println(len(s3), cap(s3), s3)
	fmt.Println(len(s4), cap(s4), s4)

	// 8. append 函数
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	var x2 []int
	x2 = append(x2, 1)
	x2 = append(x2, 2, 3)
	x2 = append(x2, 4, 5, 6)
	x2 = append(x2, x2...)
	fmt.Println(x2)

	// 9. slice就地修改
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)           //  `["one" "three" "three"]`

	// slice可以用来实现栈
	var stack = []int{1, 2, 3}
	stack = append(stack, 4)     // push
	top := stack[len(stack)-1]   // 栈顶
	stack = stack[:len(stack)-1] // pop
	fmt.Println(top, stack)

	// 从中间移除元素
	fmt.Println(remove(stack, 1))

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// slice有空间 扩展slice内容
		z = x[:zlen]
	} else {
		// slice 已经没没有空间，为他分配一个新的底层数组
		// 为了达到分摊线性复杂性，容量扩展一倍数
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) //内置copy函数
	}
	z[len(x)] = y
	return z
}

func nonempty(stringSlice []string) []string {
	// 该方法中 输入的 slice和输出的 slice拥有相同的底层数组，这样就避免在函数重新分配一个数组。
	// 当然，在这种情况下， 底层数组的元素只是部分被修改
	i := 0
	for _, s := range stringSlice {
		if s != "" {
			stringSlice[i] = s
			i++
		}
	}
	return stringSlice[:i]
}

func remove(slice []int, i int) []int {
	// 从中间移除一个元素， 并保留剩余元素
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
