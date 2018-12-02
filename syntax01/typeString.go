package main

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {

	// =================
	// 字符串
	// 字符串有编码相关问题，一般情况下，字符串被解读成按照 utf-8编码的 Unicode码点序列
	// 内置的 len函数返回的是字符串的字节数， 下标访问操作s[i]则取得第i个字符，其中 0<=i<len(s)
	// len返回的
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])

	// 访问超出范围的下标会宕机

	// 子串生成操作 s[i:j]产生一个新的字符串，内容自i开始(包含i) 到 j截止(不含j)的字符串
	// 操作数i和j的默认值分别为0和len(s),若省略，则取默认值
	fmt.Println(s[0:5])
	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])

	// + 连接两个字符串生成一个新的字符串
	// 字符串值无法改变，但是字符串变量可以被重新赋值
	s1 := "left foot"
	t1 := s1
	s1 += ", right foot"
	fmt.Println(s1)
	fmt.Println(t1)

	//=========================
	// 字符串字面量

	// 转意符 转意序列以\开始
	/**
	\a	"警告"或者响铃
	\b	退格符
	\f 	换页符
	\n	换行符
	\r	回车符（指返回行首）
	\t	制表符
	\v	垂直制表符
	\'	单引号
	\"	双引号
	\\	反斜杠
	*/

	// 源码字符串中可以包含 十六进制或者八进制的任意字节
	// \xhh(h是十六进制的数字，且必须是2位)
	// \ooo必须使用三位八进制数字，且不能超过\377
	// 这两者都表示单个字节，内容是给定值

	// ----
	// 原生的字符串字面量的书写形式是 `...`
	const c1 = "I just want to see primordial string literal \n \t \xfa \345"
	const c2 = `I just want to see
primordial string literal \n \t \xfa \345
and so on`
	fmt.Println(c1+"\n", c2)

	// ======================
	// UTF-8
	// utf-8 码点
	fmt.Println("世界")
	fmt.Println("\xe4\xb8\x96\xe7\x95\x8c")
	fmt.Println("\u4e16\u754c")
	fmt.Println("\U00004e16\U0000754c")
	fmt.Println("世界" == "\xe4\xb8\x96\xe7\x95\x8c") // true
	fmt.Println()

	// unicode/utf8包用来处理utf8编码的字符串
	s3 := "hello, 世界"
	fmt.Println(len(s3))
	fmt.Println(len(s3))
	fmt.Println(utf8.DecodeRuneInString(s3))

	for i := 0; i < len(s3); {
		r, size := utf8.DecodeRuneInString(s3[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	n := 0
	for _, _ = range s3 {
		n++
	}
	fmt.Println(n)
	fmt.Println(utf8.RuneCountInString(s3))

	// 注意：如果把文字符号类型的slice转换车给你一个字符串，他会输出各个文字符号utf-8编码拼接的结果
	s4 := "呵呵哒"
	fmt.Printf("% x\n", s4)
	r := []rune(s4)
	fmt.Printf("%x\n", r)

	// 注意： 若将一个整数转换成字符串，其值按照文字符号类型解读，并且产生代表该文字符号值的urf-8码
	fmt.Println(string(65))     // A 而不是 65
	fmt.Println(string(0x4eac)) // 京

	// ====================
	// 字符串和字节slice
	// 4个标准包对字符串操作特别重要 ： bytes、strings、strconv、和unicode

	// strings包提供了许多函数，用于搜索、替换、比较、修整、切割与连接字符串

	// bytes也有类似函数，用于操作字节slice
	fmt.Println(intsToString([]int{1, 2, 3}))

	// strconv 包具备的函数，主要用来转换布尔/整数/浮点数位与之对应的字符串形式，或者把字符串转换为布尔/整数/浮点数。
	// 	另外还有为字符串添加/去除引号的函数
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x)) // strconv.itoa ( integer to ASCII )
	// formatInt 和 formatUint 可以按照不同的进制位格式化数字
	fmt.Println(strconv.FormatInt(int64(x), 2))
	fmt.Println(strconv.FormatInt(int64(x), 16))

	x1, err := strconv.Atoi("123")
	y1, err := strconv.ParseInt("123", 10, 64) // 十进制，最长为64位
	fmt.Println(x1, y1, err)

	// unicode 包备有判别文字符号值特性的函数

}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
