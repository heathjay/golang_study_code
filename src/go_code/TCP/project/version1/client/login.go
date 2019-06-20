package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/TCP/project/version1/common/message"
	"net"
)

//error 可以有表述信息
func Login(userId int, userPwd string) (err error) {
	//定协议'
	// fmt.Println("userId:", userId, "userPwd:", userPwd)
	// return nil

	//连接到服务器端口
	//理论上直接读配置文件
	//结合读取的知识点
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dail err:", err)
		return
	}
	defer conn.Close()

	//2.准备发送
	var mes message.Message
	mes.Type = message.LoginMesType
	//序列化实质内容到data
	var loginmes message.LoginMes
	loginmes.UserId = userId
	loginmes.UserPwd = userPwd
	loginmes.UserName = "jok"
	//将loginMessage序列化
	data, err := json.Marshal(loginmes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	mes.Data = string(data)

	//message marshal
	data, err = json.Marshal(loginmes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//到这个时候，data就是我们要发送的消息
	//先发送data长度给服务器，
	//先获取长度，在转换成byte切片
	//l := len(data)
	//package binary ---数字转换成字节序列
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//binary.PutUint32(bytes[0:4], pkgLen)
	n, err := conn.Write(buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fails:", err)
		return
	}

	//fmt.Printf("client sends the message len:%d content:=%s", len(data), string(data))

	return nil
}
