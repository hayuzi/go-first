package main

import (
	"fmt"
	"log"
	"net"
	"bufio"
)

func main() {

	// 启动
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}

}

// 广播器
type client chan<- string // 对外发送消息的通道
var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // 所有接收的客户消息
)

func broadcaster() {
	clients := make(map[client]bool) // 所有连接的客户端
	for {
		select {
		case msg := <-messages:
			// 把所有接收的消息广播给所有的客户
			// 发送消息通道
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}

}

func handleConn(conn net.Conn) {
	fmt.Println(conn)
	ch := make(chan string) // 对外发送客户消息的通道
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	ch <- "Your are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// 注意, 忽略input.Err中可能的错误

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)	// 注意，忽略网络层面的错误
	}
}
