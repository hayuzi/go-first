package main

import (
	"fmt"
	"io"
	"os"
	//"bytes"
)

// ===================
// 接口即约定，该处暂不做详细解释

// ===================
// 1. 接口类型
// 一个接口类型定义了一套方法，如果一个具体类型要实现该接口，那么必须实现接口类型定义中的所有方法
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// 通过组合已有接口可以得到新接口，这样的语法称为嵌入式接口
// 但是也可以不用嵌入式来声明
type ReadWriter interface {
	Reader
	Writer
}

// ======================
// 2. 实现接口
// 如果一个类型实现了一个接口要求的所有方法，那么这个类型实现了这个接口
// 通过下面的代码我们可以知道，interface可以被任意的对象实现，
// 同理，一个对象可以实现任意多个interface
// 最后，任意的类型都实现了空interface(我们这样定义：interface{})，也就是包含0个method的interface。

// 以下是一个示例
type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名字段Human
	school string
	loan   float32
}

type Employee struct {
	Human   //匿名字段Human
	company string
	money   float32
}

//Human对象实现Sayhi方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human对象实现Sing方法， (指针类型的方法不能被student的匿名Human传递)
// func (h *Human) Sing(lyrics string) 不能通过编译
func (h Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

// Employee重载Human的Sayhi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //此句可以分成多行
}

//Student实现BorrowMoney方法
func (s Student) BorrowMoney(amount float32) {
	s.loan += amount // (again and again and...)
}

//Employee实现SpendSalary方法
func (e Employee) SpendSalary(amount float32) {
	e.money -= amount // More vodka please!!! Get me through the day!
}

// 定义interface
type Men interface {
	SayHi()
	Sing(lyrics string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

//----------------
type Element interface{}
type List []Element

func main() {

	// ===========================
	// 3.可以使用 flag.Value来解析参数，参考sample文件夹中的 testFlag.

	// ===========================
	// 4. 接口值以及接口类型的赋值规则
	// 一个接口类型的值（简称接口值）起始有两个部分，一个具体类型和该类型的一个值。二者称为接口的动态类型和动态值
	// 那么interface里面到底能存什么值呢？如果我们定义了一个interface的变量，那么这个变量里面可以存实现这个interface的任意类型的对象。
	// 例如上面例子中，我们定义了一个Men interface类型的变量m，那么m里面可以存Human、Student或者Employee值。
	// 因为m能够持有这三种类型的对象，所以我们可以定义一个包含Men类型元素的slice，这个slice可以被赋予实现了Men接口的任意结构的对象，这个和我们传统意义上面的slice有所不同。

	// 接口值比可以使用  == 以及 != 来做比较，如果两个接口值都是 nil或者二者动态类型完全一致且二者动态值相等，那么两个接口值相等。
	// 如过两个接口值的动态值不可以比较，则比较会崩溃。 这一点需要注意。

	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	Tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	//定义Men类型的变量i
	var i Men

	//i能存储Student ()
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	//i也能存储Employee
	i = Tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义了slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	//这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}

	// =====================
	// 5. 空interface
	// 空interface(interface{})不包含任何的method，正因为如此，所有的类型都实现了空interface。
	// 空interface对于描述起不到任何的作用(因为它不包含任何的method）
	// 但是空interface在我们需要存储任意类型的数值的时候相当有用，因为它可以存储任意类型的数值。
	// 它有点类似于C语言的void*类型。
	// 一个函数把interface{}作为参数，那么他可以接受任意类型的值作为参数，如果一个函数返回interface{},那么也就可以返回任意类型的值。

	// =====================
	// 6. 含有空指针的非空接口
	// 空的接口值(其中不包含任何信息)与仅仅动态值为 nil的接口值是不一样的。

	// ====================
	// 7. 使用订sort.Interface 来排序
	// sort包专门提供了对于 []int、 []string、[]float64自然排序的函数和相关类型，对于其他类型，比如[]int64与或者 uint则需要自己写
	// 参考 sample 文件夹中的 SortTest

	// =====================
	// 5.类型 断言   (常用来识别错误或者查询特性)
	// Go语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok = element.(T)，
	// 这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。
	// 如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false。
	//
	// 5.1 如果断言类型T是一个具体类型，那么断言会检查element的会价差x的动态类型是否就是T，如果检查成功，类型断言的结果就是 element 的动态值
	// 换句话说，类型断言就是用来从它的操作数中把具体类型的值提取出来的操作。

	// 5.2 如果断言类型T是一个接口类型，那么类型断言检查element的动态类型是否满足T，如果检查成功，动态值并没有提取出来，结果仍然是一个接口值.
	// 从一个接口类型变为拥有另外一套方法的接口类型（通常发发数量是增多），但保留了接口值中的动态类型和动态值部分
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File) 	// 成功 f == os.Stdout
	// c := w.(*bytes.Buffer)	// 崩溃：接口持有的是 *os.File, 不是 *bytes.Buffer
	fmt.Println(f)
	// fmt.Println(c)

	rw := w.(io.ReadWriter)
	fmt.Println(rw)
	fmt.Printf("%T", rw)

	// 如果使用两个结果的赋值表达式，断言不会在失败的时候崩溃， 而是返回一个布尔类型的表达式
	list := make(List, 3)
	list[0] = 1       // an int
	list[1] = "Hello" // a string
	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else {
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}

	// 6. 类型s分支
	var m int
	m = 1
	switch m.(type) {
	case nil:
		fmt.Println("nil")
		break
	case int:
		fmt.Println("int")
		break
	default:
		fmt.Println("default")
	}


	// 重新命名的变量，在语法块中会替换
	switch m := m.(type) {
	case nil:
		fmt.Println("nil")
		break
	case int:
		fmt.Println(m)
		fmt.Println("int")
		break
	default:
		fmt.Println("default")
	}




}
