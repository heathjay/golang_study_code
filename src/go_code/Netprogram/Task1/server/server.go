package main

import (
	"fmt"
	"io"
	"net" //做网络开发时，包含了我们需要所有的方法和函数
)

//listen 返回一个连接， net：tcp？
//返回一个Listen 借口--- Acept  简介
//用完之后close

func process(conn net.Conn) {
	//这里我们循环的介绍客户端发送的数据
	defer conn.Close() //关闭conn

	for {
		//创建一个新的切片
		buf := make([]byte, 1024)

		//fmt.Println("服务器在等待客户端：", conn.RemoteAddr().String())
		n, err := conn.Read(buf) //	对方没有信息一致阻塞
		if err == io.EOF {
			fmt.Println("客户端退出")
			return
		}

		//3.显示客户端发送大的内容到服务器端
		//有限的个数，不然可能会出现问题
		//Read函数n := conn.Read(buf)
		fmt.Print(string(buf[:n])) //后面的所有

	}

}

func main() {
	fmt.Println("I am waiting for u....")

	listen, err := net.Listen("tcp", "127.0.0.1:8888")

	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close()

	//循环等待客户端连接我

	//read , remote, local . close
	//net.dial("tcp","google.com:80")
	//bufio.NewReader(conn).ReadString('\')
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept() err=", err)
			return
		} else {
			fmt.Printf("Accept() suc con = %v\n ip:=%v", conn, conn.RemoteAddr().String())

		}

		//这里准备其一个协程，为客户端服务
		go process(conn)
	}

	fmt.Printf("listen = %V\n", listen)
}
