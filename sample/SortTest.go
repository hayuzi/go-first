package main

import (
	"fmt"
	"sort"
)

type Int64Slice []int64

// 排序需要的类型需要实现sort.Interface接口. 该接口有如下三个方法
func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	i := Int64Slice{23, 1, 15, 9, 7}
	fmt.Println(i)

	sort.Sort(i)
	fmt.Println(i)

	sort.Sort(sort.Reverse(i))
	fmt.Println(i)
}
