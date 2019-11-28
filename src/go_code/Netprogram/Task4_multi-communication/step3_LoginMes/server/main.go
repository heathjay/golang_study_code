package main

/*
	step4:
		程序结构需要强化
		1. 先改进服务器端,先画出框架图，然后写代码
			1.1 main.go---1.监听2.获取conn 3.初始化
			1.2 process.go --- 总处理器
				1.2.1 根据客户端请求调用对应的处理器完成相应的任务
				1.2.2 userProcess.go 处理跟用户相关的请求（登陆，注册，注销、用户列表
				1.2.3 smsProcess.go 处理短消息相关的请求, 群聊，点对点聊天,
			1.3 utils.go 工具类
				把一些常用的工具函数和结构体抽取，统一提供常用的方法和函数
			1.4 common/message
				服务器端和客户端公用的文件


*/

import (
	"encoding/json"
	"fmt"
	"go_code/Netprogram/Task4_multi-communication/step4_structure/common/message"
	"io"
	"net"
)

//gorountine
func process(conn net.Conn) {
	//读取用户发送信息
	defer conn.Close()

	//1 接受len
	for {

		//conn.Read 只有在conn没有被关闭的情况下， 才会阻塞,如果一方关闭了conn。则不会阻塞
		//3.1将读取消息直接封装成一个函数,
		//3.1 readPkg 返回一个message,err
		mes, err := readPkg(conn)
		if err != nil {
			fmt.Println("readPkg fails:", err)
			//连接关闭，server也退出
			if err == io.EOF {
				fmt.Println("conn disconnect")
				return
			} else {
				fmt.Println("readPkg err:", err)
				return
			}

		}
		err = serverProcess(conn, &mes)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
	}
}

//3.2.3工具函数
// func writePkg(conn net.Conn, data []byte) (err error) {
// 	//先发送一个长度
// 	var pkgLen uint32
// 	pkgLen = uint32(len(data))
// 	var buf [4]byte
// 	binary.BigEndian.PutUint32(buf[:4], pkgLen)
// 	n, err := conn.Write(buf[:4])
// 	if err != nil || n != 4 {
// 		fmt.Println("conn.Write(pkgLeg) fails")
// 		return
// 	}

// 	//发送内容
// 	n, err = conn.Write(data)
// 	if n != int(pkgLen) || err != nil {
// 		fmt.Println("conn write bytes fails:", err)
// 		return
// 	}
// 	return
// }

//3.2.2编写函数处理登陆请求,
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	//3.2.2.1先从message中取出data，反序列化成loginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.unmarshall loginMes fails")
		return
	}

	//先声明一个response
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//填充数据
	var loginResMes message.LoginResMes

	//如果用户的id = 100, 密码 = 123456， 认为合法，否则不合法
	if loginMes.UserId == 100 && loginMes.UserPw == "123456" {
		//数据库验证
		loginResMes.Code = 200

	} else {
		//不合法
		loginResMes.Code = 500 //表示用户不合法
		loginResMes.Error = "invaliad user"
	}

	//序列化回复完成
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("unmarshall loginResMes fails:", err)
		return
	}

	//将data 赋值给resMes
	resMes.Data = string(data)

	//对整个信息包进行序列化发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json marshal resMes fails:", err)
		return
	}

	//发送data封装函数
	err = writePkg(conn, data)
	if err != nil {
		fmt.Println("write data err:", err)
		return
	}
	return
}

//3.2函数优化，根据客户端发送种类不同，决定调用哪个函数来处理
func serverProcess(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登陆信息
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType:
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}

//3.1封装成函数, readPkg
//3.2优化封装结构,需要readPkg可以读取不同的消息接口
// func readPkg(conn net.Conn) (mes message.Message, err error) {
// 	//获取长度
// 	buf := make([]byte, 8096)
// 	fmt.Println("wait for the message from the sender")
// 	_, err = conn.Read(buf[:4])
// 	if err != nil {
// 		//fmt.Println("len sent error:", err)
// 		//自定义error
// 		err = errors.New("Read len fails.")
// 		return
// 	}
// 	//fmt.Println("read :", buf[:4])

// 	//根据buf[:4]判断到底读多少字节
// 	var pkgLen uint32
// 	pkgLen = binary.BigEndian.Uint32(buf[0:4])

// 	//根据pkgLen读取消息内容
// 	n, err := conn.Read(buf[:pkgLen]) //从conn读取总共pkgLen放进buf中
// 	if n != int(pkgLen) || err != nil {
// 		//fmt.Println("read LoginMes fails")
// 		//err = errors.New("read body fails")
// 		return
// 	}

// 	//反序列化 -> message.Message --> &mes！！
// 	err = json.Unmarshal(buf[:pkgLen], &mes)
// 	if err != nil {
// 		fmt.Println(" LoginMes unmarshall fails")
// 		return
// 	}

// 	return
// }

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
