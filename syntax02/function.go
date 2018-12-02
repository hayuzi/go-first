package main

import (
	"fmt"
	"image"
	"io"
	"log"
	"math"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	// =====================
	// 1. Go 中， 实参是按值传递的

	fmt.Println(hypot(3, 4))
	fmt.Println(add(2, 2))
	fmt.Println(sub(15, 8))
	fmt.Printf("%T\n", hypot)
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", first)
	fmt.Printf("%T\n", f1)
	fmt.Printf("%T\n", f2)

	// ======================
	// 2.函数的递归 recursion, 在函数中调用自身
	fmt.Println(recursion(1))

	// ======================
	// 3.多返回值, 多返回值的函数，给返回值一个良好的名称很有意义
	// 一个函数如果有命名的返回值，可以神略return 语句的操作数，这称为裸返回
	fmt.Printf("%T\n", test)
	y, yes := test(1)
	fmt.Println(y, yes)

	// ======================
	// 4. 错误与错误处理策略
	// Go语言使用通常的控制流机制（if / return）应对错误 ，想要详细了解可以寻求相关信息，这里不做详细描述
	fmt.Println(io.EOF)

	// 5. 函数变量， 把函数类型赋值给一个变量, 函数变量不可比较， 函数类型的零值是 nil
	f := add
	fmt.Printf("%T\n", f)
	fmt.Println(f(1, 3))

	// 6. 匿名函数
	// 匿名函数只能在包级别的作用域声明，
	// 不过匿名函数可以调用外层函数的局部变量，此时隐藏着变量引用，变量x在squares调用后依然存在
	// 一般来说Go程序员通常称呼函数变量为闭包
	f3 := squares()
	fmt.Println(f3())
	fmt.Println(f3())
	fmt.Println(f3())
	fmt.Println(f3())

	// 7. 变长函数，如下的声明方式 ... 展开符号表示多个参数
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3, 4, 5))
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))

	// =====================
	// 8.延迟函数调用
	// defer 语句可以延迟函数的执行， defer的函数执行顺序为 defer调用顺序的倒序
	// defer语句经常拥有成对的操作，如 关闭打开资源，或者解锁一个互斥锁
	var mu sync.Mutex
	var m = make(map[string]int)
	m["test"] = 1
	var lookup = func(key string) int {
		mu.Lock()
		defer mu.Unlock()
		return m[key]
	}
	fmt.Println(lookup("test"))
	// 使用defer做trace
	bigSlowOperation()

	// -------
	// 	延迟函数在return之后调用, 因此defer可以改变外层函数反馈给调用者的返回值, 示例如下
	fmt.Println(deferTest(2))

	// 延迟函数不到当前函数最后一刻是不会执行的， 因此尤其需要注意循环中调用 defer生效的问题
	// 避免在循环中使用 defer

	// 不过本地文件的io如果打开并操作，由于写错误往往是在 资源Close()的时候返回的
	// 这个时候就需要避免使用defer，而是直接处理获取结果，或者解析失败原因

	// 9. 宕机 runtime包提供了转储栈的方法使程序员可以诊断错误
	var printStack = func() {
		var buf [4096]byte
		n := runtime.Stack(buf[:], false)
		os.Stdout.Write(buf[:n])
	}
	defer printStack()

	var n2 = func(x int) int {
		return 100 / x
	}
	// 此处的参数如果为 0 会导致宕机， 但是栈追踪信息会被上面的代码记录打印
	fmt.Println(n2(1))

	// ========================
	// 10. 恢复 recover() 和 panic()

	// 如果内置的recover函数在延迟函数内部调用， 而且这个包含defer语句的函数发生宕机,
	// recover会种植当前的宕机状态并返回宕机的值。函数不会从之前宕机的地方继续运行而是正常返回，
	// 如果recover在其他任何情况下运行，它不会有任何效果
	y, err := recoverTestOne(0)
	fmt.Println(y, err)

	// 可以通过使用一个明确的非导出类型作为宕机的值，之后recover检测饭不一致是否是这个类型
	// 如果是，可以像普通error那样处理宕机； 如果不是，使用同一个参数调用 panic 以继续触发宕机、
	y2, err := recoverTestTwo(0)
	fmt.Println(y2)
	fmt.Println(err) // 此处得到预期 err = {}， 可以进行单独处理

}

// 1.  函数的基本结构如下： 函数定义的时候，参数叫做形参， 实际调用传入的叫做实参，
// 当函数存在返回值时候，必须显式的使用 return 语句结束，并将值返回
// func name(parameter-list) result-list {
//	  body
// }

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func add(x int, y int) int {
	return x + y
}

// 将 z 返回
func sub(x, y int) (z int) {
	z = x - y
	return
}

// _ 表示形参未被使用
func first(x int, _ int) int {
	return x
}

// 使用几个实例, 以下两种方法申命是一样的效果
func f1(i, j, k int, s, t string)                {}
func f2(i int, j int, k int, s string, t string) {}

// ===============
func recursion(x int) int {
	if x < 5 {
		return recursion(x + 1)
	} else {
		return x
	}
}

// =============
func test(x int) (int, bool) {
	if x > 0 {
		return x, true
	} else {
		return x, false
	}
}

func Size(rec image.Rectangle) (width, height int) {
	point := rec.Size()
	width = point.X
	height = point.Y
	return
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

// ===============
// 变长函数
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func bigSlowOperation() {
	defer trace("bingSlowOperation")()
	// 处理区间
	time.Sleep(1 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func deferTest(x int) (result int) {
	defer func() { result += x }()
	return x + 1
}

func recoverTestOne(x int) (y int, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error : %v", p)
		}
	}()
	y = 100 / x
	return
}

func recoverTestTwo(x int) (y int, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			// 没有宕机 不做任何处理
		case bailout{}:
			err = fmt.Errorf("internal error : %v", p)
		default:
			panic(p)
		}
	}()

	// 此处主动做预期宕机处理
	if x == 0 {
		panic(bailout{})
	}
	y = 100 / x
	return
}
