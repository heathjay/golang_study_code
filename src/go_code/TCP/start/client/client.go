package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	//fmt.Println("conn 成功=", conn)
	//功能1:客户端可以发送单行数据，然后退出
	//stdin stdout:标准输入输出，错误
	//本质是文件读取，所有都是文件，

	for {
		reader := bufio.NewReader(os.Stdin) //标准输入
		//终端读取一行用户输入，并准备发送到服务器
		content, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
		}
		//再将content 发送给服务器
		n, err := conn.Write([]byte(content))
		if err != nil {
			fmt.Println("con.write err=", err)
		}
		if content == "exit\n" {
			break
		}
		fmt.Println("sendre done:", n)
	}

}
