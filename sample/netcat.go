// netcat 是一个只读的TCP客户端程序
package main

import (
	"io"
	"net"
	"os"
	"log"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn) // 注意：忽略错误
		log.Println("done")
		done <- struct{}{} // 指示主 goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // 等待后台 goroutine 完成
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
