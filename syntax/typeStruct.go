package main

import (
	"fmt"
	"time"
)

// 1. 结构体是将领个或者多个任意类型的命名变量组合在一起的聚合数据类型， 每个变量叫做结构体的成员
// 大写字母开头的结构体变量是可以导出的.*** 访问控制
// 结构的成员变量通常一行写一个，变量的名称在类型的前面，但是相同类型的连续成员变量可以写在一行
// 成员变量的顺序对于结构体同一性很重要。
// ** 如果我们将也是字符串类型的 Position,Name,Address组合在一起或互换了Name和Address的顺序,
// 那么我们就在定义一个不同的结构体 **

// 一般情况下我们不写匿名结构体
type Employee struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

type tree struct {
	value       int
	left, right *tree
}

func main() {
	var dilbert Employee

	// 2. 成员变量通过 . 访问, 或者通过获取成员变量的地址，然后通过指针来访问
	dilbert.Salary -= 5000
	position := &dilbert.Position
	*position = "Senior " + *position

	fmt.Println(dilbert.Salary)
	fmt.Println(dilbert.Position)

	// 3. . 点号同样可以作用在结构体指针上
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	// 后面这条等价于 (*employeeOfTheMonth).Position += " (proactive team player)"

	// 4. 一个聚合类型不可以包含它自己（ 适用于结构体/数组 ）
	// 命名类型结构体类型 s 不可以定义一个拥有相同结构体 s 的成员变量。
	// 但是 s 中可以定义一个 s 的指针类型， 即 *s, 这样我们就可以创建一些递归数据结构， 比如链表和树
	s1 := []int{5, 3, 123, 12, 123, 55, 512, 44, 58}
	fmt.Println(s1)
	Sort(s1)
	fmt.Println(s1)

}

// 利用二叉树来实现插入排序
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// 等价与返回 &tree{value: value}
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
