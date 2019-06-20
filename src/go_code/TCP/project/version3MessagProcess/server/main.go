package main

import (
	"encoding/json"
	"fmt"
	utils "go_code/TCP/project/version3MessagProcess/common"
	"go_code/TCP/project/version3MessagProcess/common/message"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		//循环读取客户端发消息
		//直接封装成函数，返回一个message和一个err
		mes, err := utils.ReadPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("client out,server up")
				return
			} else {
				fmt.Println("reading err=", err)
				return
			}

		}
		//fmt.Println("mes=", mes)
		err = severProcessMes(conn, &mes)
		if err != nil {
			return
		}
	}
}

//登陆请求
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	//核心代码
	var loginmes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginmes)
	if err != nil {
		fmt.Println("serverProcessLogin json.Unmarshal fails:", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//申明一个loginRes类型
	var loginResMes message.LoginResMes

	//如果用户id为100，密码为123456认为合法，否则不合法
	if loginmes.UserId == 100 && loginmes.UserPwd == "123456" {
		//合法
		loginResMes.Code = 200
		loginResMes.Error = "ok"
	} else {
		//不合法
		loginResMes.Code = 500
		loginResMes.Error = "please register first"
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.marshal fails:", err)
	}
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.marshal fails:", err)
		return
	}

	//发送用函数来弄
	err = utils.WritePkg(conn, data)
	return
}

/*
func writePkg(conn net.Conn, data []byte) (err error) {
	//长度
	pkgLen := uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	n, err := conn.Write(buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fails:", err)
		return
	}

	n, err = conn.Write(data)
	if uint32(n) != pkgLen || err != nil {
		fmt.Println("conn.Write(data) fails:", err)
		return
	}
	return
}
*/
//根据客户端发送的消息种类不同，决定调用哪一个函数
func severProcessMes(conn net.Conn, mes *message.Message) (err error) {
	//用switch比较好
	switch mes.Type {
	case message.LoginMesType:
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType:
		fmt.Println("register")
	default:
		fmt.Println("input again")
		return
	}
	return nil
}

/*
func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 1024*4)
	fmt.Println("读取客户端发送的数据")
	//conn.Read 在conn没有被关闭的情况才会组赛
	//如果客户端关闭了
	_, err = conn.Read(buf[:4])
	if err != nil {
		//fmt.Println("connect.read is error:", err)
		//err = errors.New("reading pkg header error")
		return
	}
	fmt.Println("good job,buf:= ", buf[0:4])

	//内容长度
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	fmt.Println(pkgLen)
	//根据pkgLen读取内容
	//从conn里面读取放在buf内容
	n, err := conn.Read(buf[:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		//err = errors.New("read n error")
		//fmt.Println("conn.read fail err=", err)
		return
	}
	//把pkgLen反序列成message类型
	//记得&mes
	v := buf[0:pkgLen]
	err = json.Unmarshal(v, &mes)
	if err != nil {
		err = errors.New("json,Unmarshal err")
		//fmt.Println("json,Unmarshal err:", err)
		return
	}
	//fmt.Printf("%v", m)
	err = severProcessMes(conn, &mes)
	return
}
*/
func main() {
	fmt.Println("server is listening by port 8889")
	Listerner, err := net.Listen("tcp", "0.0.0.0:8889")

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
	defer Listerner.Close()
}
