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

type Point struct{ X, Y int }

// 定义一个结构体
type Circle struct {
	X, Y, Radius int
}

// 定义一个嵌套
type Wheel struct {
	X, Y, Radius, Spokes int
}

// Circle可以改造成如下结构体
type Circle2 struct {
	Center Point
	Radius int
}

// Wheel可以改造成这样
type Wheel2 struct {
	Circle Circle2
	Spokes int
}

type Circle3 struct {
	Point  // 匿名嵌入
	Radius int
}

type Wheel3 struct {
	Circle3
	Spokes int
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
	fmt.Println(employeeOfTheMonth)

	// =====================
	// 5. 结构体字面量
	// 格式一： 该格式必须按照正确的顺序给全部数据赋值， 一般用于有明显成员变量顺序约定的小结构体重
	p1 := Point{1, 2}

	// 格式二： 指定附文后者全部成员变量的名称和值来初始化结构体变量， 可以不考虑顺序
	// 这种方式如果有成员变量没有指定，那么该成员变量的值会是其类型的 零值
	p2 := Point{X: 1, Y: 2}

	// 以上两种初始化方式不可以混合使用，也不可以用第一种初始化方式来绕过不可导出变量无法在其他包中使用的规则
	fmt.Println(p1, p2)

	// ============
	// 6. 结构体类型的值可以作为参数传递给函数或者作为函数的返回值
	fmt.Println(Scale(Point{1, 2}, 5))

	// 出于效率的考虑，大型结构体通常使用结构体指针的方式直接传递给函数或者从函数中返回
	// 这种方式在函数需要修改结构体内容的时候也是必须的，
	// ** 在Go这种按照值调用的语言中，调用的函数接收到的是实参的一个副本，并不是实参的引用
	employeeOfTheMonth.Salary = 10000
	fmt.Println(dilbert)
	Bonus(employeeOfTheMonth, 200)
	fmt.Println(dilbert)

	// 7. 由于通常结构体都通过指针的方式使用， 因此可以使用一种简单的方式来创建/初始化一个struct类型的变量并获取他的地址
	pp := &Point{1, 2}
	fmt.Println(pp)

	pp2 := new(Point)
	*pp2 = Point{1, 2}
	fmt.Println(pp2)

	// 8. 结构体比较
	// 结构体可以使用  == 以及 != 比较。
	// 和其他可以比较的类型一样，可比较的结构体类型都可以作为 map的键类型

	p5 := Point{1, 2}
	p6 := Point{1, 2}
	fmt.Println(p5, p6, p5 == p6)

	// 9. 结构体嵌套和匿名成员
	// 以上创建了好几个对象， Circle Wheel 等
	// 创建一个Wheel类型的对象
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
	// 如上的访问方式访问对象的成员，但是创建结构体的时候比较麻烦，结构体可以嵌套，所以我们优化 Wheel 为 Wheel2
	var w2 Wheel2
	w2.Circle.Center.X = 8
	w2.Circle.Center.Y = 8
	w2.Circle.Radius = 5
	w2.Spokes = 20
	// 如此以来结构体的对象成员访问太麻烦了
	// Go允许我们定义不带名称的结构体成员，只需要制定类型即可；
	// 这种结构体成员叫做匿名成员， 这个结构体成员的类型必须是一个命名类型或者指向命名类型的指针
	// 由于有匿名嵌套结构体的功能，我们可以直接访问到我们需要的变量而不是指定一大串中间变量
	var w3 Wheel3
	w3.X = 8       // 等价于 w3.Circle3.Point.X = 8
	w3.Y = 8       // 等价于 w3.Circle3.Point.Y = 8
	w3.Radius = 5  // 等价于 w3.Circle3.Radius = 5
	w3.Spokes = 20 // 等价于 w3.Spokes = 20
	// 上面注释里面的方式也是正确的，但是使用匿名成员的说法或许不合适
	// 上面的结构体成员Circle3 和 Point是有名字的，就是对应类型的名字，只是浙西额名字在点好访问变量时候是可选的
	// 当我们访问最终需要的变量的时候可以省略中间所有的匿名成员
	// 遗憾的是， 结构体字面量并没有什么快捷方式来初始化结构体
	w4 := Wheel3{Circle3{Point{8, 8}, 5}, 20}
	w5 := Wheel3{
		Circle3: Circle3{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Printf("%v\n", w4)  // 只输出值
	fmt.Printf("%#v\n", w5) // 类似Go愈发的方式输出对象，这个方式里面包含了成员变量的名字

	// .号访问匿名成员内部变量的语法糖，可以跨过不可导出结构体而访问其中的可导出成员
	// 假如 w3内部的 circle与point是不可倒出的。
	// w3.X依然可以访问 point中的X, 但是   w3.circle.point.X在包外部值是不可访问的

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

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}
