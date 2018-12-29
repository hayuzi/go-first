package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"path/filepath"
	"flag"
	"sync"
	"time"
)

var verbos = flag.Bool("v", false, "show verbos progress messages")

func main() {
	// 确定目录
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fmt.Println(roots)

	// 并行遍历每一个文件树
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// 检测到输入时候取消遍历
	go func() {
		os.Stdin.Read(make([]byte, 1)) // 读取一个字节
		close(done)
	}()

	// 定期输出结果
	var tick <-chan time.Time
	fmt.Println(*verbos)
	if *verbos {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			for range fileSizes {
				// 不执行任何操作, 耗尽fileSizes以允许已有的goroutine结束
			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes 关闭
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

// walkDir 递归地遍历以dir为根目录的整个文件树
// 并在 fileSize 上发送每个已经找到的文件的大小
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// sema 是一个用于限制并发数的计数信号量
var sema = make(chan struct{}, 20)

// dirents 返回dir目录中的条目
func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // 获取令牌
	case <-done:
		return nil //取消
	}
	defer func() { <-sema }()		// 释放令牌
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}


// 增加取消机制
var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
	return true
	default:
		return false
	}
}
