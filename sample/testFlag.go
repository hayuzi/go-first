package main

import (
	"flag"
	"fmt"
	"strconv"
)

// *testFlag 满足 flag.value 接口
type SingleNum int64

type testFlag struct {
	SingleNum
}

func (t *testFlag) String() string {
	return fmt.Sprintf("%d", t.SingleNum)
}

func (t *testFlag) Set(s string) error {
	switch s {
	case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
		if num, err := strconv.ParseInt(s, 10, 64); err == nil {
			t.SingleNum = SingleNum(num)
		}
		return nil
	}
	return fmt.Errorf("invalid SingleNum %q", s)
}

func TestFlag(name string, value SingleNum, usage string) *SingleNum {
	f := testFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.SingleNum
}

var num = TestFlag("num", 0, "the single num")

func main() {
	flag.Parse()
	fmt.Println(*num)
}
