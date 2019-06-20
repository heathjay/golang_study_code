/*
客户端： who are u 服务器端：我是小冰
客户端： 你的性别 服务器端：你猜猜看
客户端： 你会什么， 服务器端：我讲故事

*/

/*
监听端口 8888
接收客户端tcp，建立客户端和服务器端的链接
创建goroutine,处理该链接的请求
*/

package main

import (
	"fmt"
	"io"
	"net" //做网络socket开发的时候，要使用
)

func process(conn net.Conn) {
	//循环接受客户端发送的数据
	//循环读取
	//实现点对点会话，
	//小心！！
	defer conn.Close() //及时关闭
	for {
		info := make([]byte, 512)
		fmt.Println("waits")
		//1.等待客户端通过conn发送信息给我
		//2.如果客户端没有write，那么就一直阻塞在这里,协程就阻塞在这里
		n, err := conn.Read(info) //链接阻塞
		if err == io.EOF {
			fmt.Printf("remore host out")
			return
		}
		//服务器端显示
		//数据显示客户端发送的内容到服务器的终端
		//客户端是有戴\n

		fmt.Print(string(info[:n]))
	}
}

func main() {

	fmt.Println("server is listening.....")
	//net.Listen("tcp","127.0.0.1:8888")//只能支持ipv4
	listen, err := net.Listen("tcp", "0.0.0.0:8888") //ipv4/ipv6
	if err != nil {
		fmt.Println("listen err =", err)
		return
	}
	defer listen.Close()
	//循环等待客户端来连接我
	for {
		fmt.Println("waiting for connecting")
		con, AcceptErr := listen.Accept()
		if AcceptErr != nil {
			fmt.Println("accept error=", AcceptErr)
			return
		} else {
			fmt.Printf("connect=%v, 客户端地址:=%v \n", con, con.RemoteAddr().String())
		}
		go process(con)
		//准备起一个协程为客户端服务
		//接受一个工作不要在主线程上写，
	}

	//fmt.Printf("listen=%v\n", listen) //listen 就是一个接口
	//listen 下一个accept
}
