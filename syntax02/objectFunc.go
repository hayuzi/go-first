package main

import (
	"bytes"
	"fmt"
	"image/color"
	"math"
	"sync"
)

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // 函数调用
	fmt.Println(p.Distance(q))  // 方法调用

	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	length := r.Distance(q) // 合法,会隐式转换
	fmt.Println(length)
	p.ScaleBy(2) // 合法,会隐式转换
	fmt.Println(p)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var pc1 = ColoredPoint{Point{1, 1}, red}
	var pc2 = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(pc1.Distance(pc2.Point)) // 省略Point直接访问到 Point的 Distance方法
	pc1.ScaleBy(2)
	pc2.ScaleBy(2)
	fmt.Println(pc1.Distance(pc2.Point))

	fmt.Println(lookup("red"))
	fmt.Printf("%T\n", lookup("red"))

	funcVar() // 直接调用这个函数

	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(256)
	fmt.Println(x)
	fmt.Println(x.String())
	fmt.Println(&x) // *Inset的方法String被 fmt默认调用，此处产生意外。实际上 IntSet的值并不含有String方法
	//fmt.Println(y)

}

// =============
// 1. 方法和函数声明类似，只是在函数名字面前多了一个参数，这个参数把这个方法绑定到这个参数对应的类型上
type Point struct{ X, Y float64 }

// 普通函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Point类型的方法, 附加的参数P称为方法的接收者， 它源自早先的面向对象语言， 用来描述主调方法， 就像像对象发送消息
// Go 语言中接收者不使用特殊名（比如this或者self），通常选择类型名称的首字母作为方法的接收者名称， 如下：p Point
// Point结构体类型中X， Y已经使用，如果方法声明为X则会冲突，并且编译报错
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Go语言和许多其他面向面向对象语言不同，它可以将方法绑定到任何类型上
// 可以很方便的为简单类型（如数字，字符串 slice， map， 甚至函数等）定义附加行为
// 同一个包下的任何类型都可以声明发发， 只要它的类型既不是指针类型也不是接口类型

// 类型拥有的所有方法名都必须唯一，但是不同的类型可以使用相同的方法名
// 这就能看出使用方法的第一个好处，命名可以比函数更简单，
// 在包外部调用的时候，方法能够使用更加间断的名字且省略包的名字

// ==============
// 2. 指针接收者的方法
// 由于主调函数会复制每一个参实参变量，如果函数需要更新一个变量，或者一个实参太大我们希望避免复制整个实参
// 我们就可以使用指针来传递变量的地址，这也同样适用于更新接收者，我们将它绑定到指针类型
// 下面这个方法的名字是 (*Point).ScaleBy 圆括号是必须的，否则会被解析为 *(Point.ScaleBy)
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// ================
// 3. nil是一个合法的接收者
// 当定义一个类型允许nil作为接收者时，应当在文档注释中显式地表明，如下面的例子所示

// IntList是一个整型链表
// *IntList的类型nil代表空列表
type IntList struct {
	Value int
	Tail  *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

// ===================
// 4. 通过结构体内嵌组成类型, 我们可以可以直接调用内嵌类型 Point的方法

type ColoredPoint struct {
	Point
	Color color.RGBA
}

// 这样使用指针类型也是可以的
type ColoredPoint2 struct {
	*Point
	color.RGBA
}

// 下面是一个很好的缓存实现
// 匿名结构体直接生成对象并赋值给cache
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: map[string]string{
		"red": "redValue",
	},
}

func lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

// =============
// 5. 方法变量与表达式
func funcVar() {
	p := Point{1, 2}
	q := Point{4, 6}

	// 这样得到的是方法变量
	distanceFromP := p.Distance

	// 这样的声明方式是 方法表达式,  方法表达式是一种函数变量
	// 把原来方法的接收者替换成函数的第一个形参，因此它可以像平常函数一样调用
	distance := Point.Distance
	fmt.Println(distanceFromP(q))
	fmt.Println(distance(p, q))
}

// ==================
// 6. 位向量
// Go语言中的集合通常使用 map[T]bool 来实现， 其中 T 是元素类型
// 位向量是一个无符号整型值的 slice， 每一位代表集合中的一个元素。
// 如果设置第i位的元素，则认为集合包含i

type IntSet struct {
	words []uint64
}

// Has方法的范湖遗址表示是否存在非负数x
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add添加非负数x到集合中
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//=================
// 7. 封装
// 如果变量或者方法不能通过对象访问到，这称为封装的变量或者方法。
// 封装（有时候称作数据隐藏）是面向对象编程中重要的一方面
// Go语言只有一种方式控制命名的可见性： 定义的时候，首字母大写的标识符是可以从包中导出的，而首字母没有大写的则不导出
// 同样的机制也同样作用于结构体内的字段和类型中的方法
// 7.1  结论就是要封装一个对象，必须使用结构体
// 7.2  在Go语言中封装的单元是包而不是类型
// 7.3  封装提供了三个优点
//			使用方不能直接修改对象的变量，所以不需要更多的语句来检查变量的值
// 			隐藏实现细节可以防止使用方以来的属性发生改变，使得设计者可以更加灵活的改变API的实现而不破坏兼容性
//			防止使用者肆意的改变对象内的变量
