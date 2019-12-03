/*
	step5: redis保存
		用hash
		100  用户信息 str
		只需要进行hash检索就能找到对应的用户信息
		2.加一层module层，夹在server和redis中间，model【数据】
			2.1 添加一个user结构体 user.go
			2.2 userDao.go 数据库访问对象 dao：data access object
				2.2.1编写user对象实例操作的各种方法。主要就是对其操作的方法包括增删改查。

			2.3 error.go 自定义错误
		3.客户端和服务器端的交互，通过utils
		4.服务器端用dao操作后台
		5.初始化连接池，从连接池里面初始化一批连接池，通过服务器来维护。redis.go


		5.1输入用户信息提示相应结果信息
			1. 用户不存在，你可以重新注册再登陆
			2. 你的密码不正确
		5.2 需要对redis进行初始化，redis.go放在main包里面
*/

/*
	6.注册步骤
		1. common/message/message.go 新增加两个message类型
*/
package main

import (
	"fmt"
	"go_code/Netprogram/Task4_multi-communication/step5_redis/server/model"
	"net"
	"time"
)

//gorountine
func process(conn net.Conn) {
	//读取用户发送信息
	defer conn.Close()

	//1 接受len
	//4.3调用总控
	processor := &Processor{
		Conn: conn,
	}
	err := processor.processor2()
	if err != nil {
		fmt.Println("goroutine err:", err)
		return
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

// //3.2.2编写函数处理登陆请求,
// func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
// 	//3.2.2.1先从message中取出data，反序列化成loginMes
// 	var loginMes message.LoginMes
// 	err = json.Unmarshal([]byte(mes.Data), &loginMes)
// 	if err != nil {
// 		fmt.Println("json.unmarshall loginMes fails")
// 		return
// 	}

// 	//先声明一个response
// 	var resMes message.Message
// 	resMes.Type = message.LoginResMesType

// 	//填充数据
// 	var loginResMes message.LoginResMes

// 	//如果用户的id = 100, 密码 = 123456， 认为合法，否则不合法
// 	if loginMes.UserId == 100 && loginMes.UserPw == "123456" {
// 		//数据库验证
// 		loginResMes.Code = 200

// 	} else {
// 		//不合法
// 		loginResMes.Code = 500 //表示用户不合法
// 		loginResMes.Error = "invaliad user"
// 	}

// 	//序列化回复完成
// 	data, err := json.Marshal(loginResMes)
// 	if err != nil {
// 		fmt.Println("unmarshall loginResMes fails:", err)
// 		return
// 	}

// 	//将data 赋值给resMes
// 	resMes.Data = string(data)

// 	//对整个信息包进行序列化发送
// 	data, err = json.Marshal(resMes)
// 	if err != nil {
// 		fmt.Println("json marshal resMes fails:", err)
// 		return
// 	}

// 	//发送data封装函数
// 	err = writePkg(conn, data)
// 	if err != nil {
// 		fmt.Println("write data err:", err)
// 		return
// 	}
// 	return
// }

// //3.2函数优化，根据客户端发送种类不同，决定调用哪个函数来处理
// func serverProcess(conn net.Conn, mes *message.Message) (err error) {
// 	switch mes.Type {
// 	case message.LoginMesType:
// 		//处理登陆信息
// 		err = serverProcessLogin(conn, mes)
// 	case message.RegisterMesType:
// 	default:
// 		fmt.Println("消息类型不存在，无法处理")
// 	}
// 	return
// }

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

/*
	5.3编写一个函数完成对UserDao的初始化任务
*/
func initUserDao() {
	//这里的pool本身就是一个全局变量
	//注意初始化的顺序问题，先调用initPool再调用iniUserDao
	model.MyUserDao = model.NewUserDao(pool)
}

func init() {
	//5.2当服务器启动时，初始化redis连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	//5.3初始化MyUserDao
	initUserDao()
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
