package main

import (
	"fmt"
	"net"
)

//gorountine
func process(conn net.Conn) {
	//读取用户发送信息
	defer conn.Close()

	buf := make([]byte, 8096)
	//1 接受len
	for {

		fmt.Println("wait for the message from the sender")
		_, err := conn.Read(buf[:4])
		if err != nil {
			fmt.Println("len sent error:", err)
			return
		}
		fmt.Println("read :", buf[:4])

	}
}

func main() {
	fmt.Println("服务器再8889端口监听")
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("listen err")
		return
	}
	defer listen.Close()
	//listen succ
	for {
		fmt.Println("wait for the connection")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("conn error")
		}

		go process(conn)

	}
}
