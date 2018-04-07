package main

import (
	"fmt"
	"net"
)

// conn, err := net.Dial("tcp", "127.0.0.1:5000")
// if err != nil {
// 	// handle error
// }
// fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")

func main() {
	currency := 20 //并发数,记住，一个连接数是打开一个端口号，window和linux的端口号都是有限制的
	//每条连接发送多少次连接
	for i := 0; i < currency; i++ {
		go func() {

			sendMessage()

		}()
	}
	select {}
}

func sendMessage() {
	conn, err := net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		panic("error")
	}
	header := "GET / HTTP/1.0\r\n\r\n"
	fmt.Fprintf(conn, header)
}
