package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	utils "go_code/TCP/project/version3MessagProcess/common"
	"go_code/TCP/project/version3MessagProcess/common/message"
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
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//mes.Data = string(data)
	//到这个时候，data就是我们要发送的消息
	//先发送data长度给服务器，
	//先获取长度，在转换成byte切片
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

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fails:", err)
		return
	}
	//time.Sleep(5 * time.Second)
	mes, err = utils.ReadPkg(conn)
	if err != nil {
		fmt.Println("mes receive err:", err)
		return
	}
	//把mes的data部分反序列化成LoginResMes
	var ResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &ResMes)
	//fmt.Println(ResMes)
	if ResMes.Code == 200 {
		fmt.Println("login success")
	} else if ResMes.Code == 500 {
		fmt.Println(ResMes.Error)
	}

	//这里还需要处理服务器返回的消息
	return
}
