package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err = ", err)
		return
	}
	fmt.Println("登陆成功, conn:=", conn)

	//功能一：客户端可以发送单行数据，然后就退出
	reader := bufio.NewReader(os.Stdin) //os.Stdin 代表标准输入
	//从终端读取一行用户的输入。并准备发送到服务器
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err = ", err)
			return
		}
		//将line发送给服务器
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("客户端退出...")
			break
		}

		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write err = ", err)

		}
	}

	fmt.Println("客户端发送了数据，并推出")
}
