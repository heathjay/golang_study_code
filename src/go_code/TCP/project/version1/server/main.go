package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024*4)
	for {
		//循环读取客户端发消息

		fmt.Println("读取客户端发送的数据")
		_, err := conn.Read(buf[:4])
		if err != nil {
			fmt.Println("connect.read is error:", err)
			return
		}
		fmt.Println("good job,buf:= ", buf[0:4])
	}
}

func main() {
	fmt.Println("server is listening by port 8889")
	Listerner, err := net.Listen("tcp", "0.0.0.0:8889")
	defer Listerner.Close()
	if err != nil {
		fmt.Println("listener is err:", err)
		return
	}
	fmt.Println("listener is working")
	for {
		fmt.Println("waiting for connection")
		conn, err := Listerner.Accept()
		if err != nil {
			fmt.Println("connection err:", err)
			break
		}
		//一旦链接成功，则启动一个协程和客户端保持数据的交互
		go process(conn)
	}

}
