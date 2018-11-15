package main

import (
	"fmt"
	"sort"
)

func main() {
	// ==================
	// 1. 在Go语言中， map是散列表的引用.
	// map类型是map[k]v, 其中 k 和 v 是字典的键和值对应的数据类型
	// map中所有的键都有相同的数据类型, 同时所有的值也有相同的类型，但是键和值的类型不一定相同，
	// 键的类型k, 必须是可以通过操作符 == 来进行比较的数据类型, 所以map可以检测某一个键是否存在

	// 2. 内置函数make可以创建一个 map, 也可以使用字面量来创建一个带初始化键值对元素的字典
	ages := make(map[string]int) // 创建一个从 string 到 int 的 map
	ages["alice"] = 31
	ages["charlie"] = 34
	ages2 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	agesEmpty := map[string]int{}
	var agesNil map[string]int
	// agesNil["test"] = 12  // 此处直接赋值会宕机, var声明零值map之后必须初始化才能设置元素
	fmt.Println(ages, ages2, agesEmpty, agesNil)

	// 3. map使用给定的键来查找元素
	delete(ages2, "alice")
	delete(agesNil, "blank")  // 安全的, 即便不存在也不会报错
	fmt.Println(ages["june"]) // 直接访问不存在的元素则返回类型的零值
	ages["june"] += 1

	// 4. map元素不是一个变量，不能获取它的地址
	// 如 _ = &ages["bob"] // 编译错误, 无法获取map元素的地址

	// 5. map中元素的迭代顺序是不固定的，不同的实现方法会使用不同的散列算法，得到不同的元素顺序。
	// 实践中我们认为这种顺序是随机的
	// 如果是键是字符串类型，可以使用sort包中的 Strings 函数来进行键的排序， 这是一种常见的模式
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
	fmt.Println(names, len(names), cap(names))

	// 6. 通过下表访问 map中的元素总是有值 (未设置的会是零值)， 可以通过如下方式判断元素是否存在
	age, ok := ages["bob"]
	if !ok {
		fmt.Println("bob 不是字典中的额键 age=", age)
	}
	// 合并成一条语句如下
	if age, ok := ages["bob"]; !ok {
		fmt.Println("ok=", ok, age)
	}

	// 7. map不可比较，唯一合法的是和 nil比较， 为了判断两个map是否拥有相同的键值，必须要写一个循环去判断
	// 使用 ok 来区分 "元素不存在" 和 "元素存在但是值为零值"的情况

	// 8. Go没有提供集合类型, 但是既然 map的键是唯一的, 就可以用map来实现这个功能.
	// 以下测试例子所有的键就可以看成是集合的元素
	testSet := map[string]bool{
		"element1": true,
		"element2": true,
		"element3": true,
	}
	for ele := range testSet {
		fmt.Println(ele) // 可以看到多次执行 ele的取的顺序不一样
	}

	// 文件结束 EOF

}
