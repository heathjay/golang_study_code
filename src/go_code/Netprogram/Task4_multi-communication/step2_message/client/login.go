package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/Netprogram/Task4_multi-communication/step2_message/common/message"
	"net"
)

//尽量用error的形式返回，
//error可以返回各种错误，自定义，描述信息
//变量尽量使用一样的
func login(userId int, userPw string) (err error) {
	//开始定协议
	//fmt.Printf("userId= %d,userPw = %s", userId, userPw)
	//return nil
	//1.连接到服务器
	//理论上说需要读取配置文件
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("conn err")
		return
	}
	defer conn.Close()
	//2.准备发送信息
	var mes message.Message
	mes.Type = message.LoginMesType

	//3.一个loginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPw = userPw

	//4.marsh
	data, err := json.Marshal(loginMes)
	//mes.Data = loginMes	//
	if err != nil {
		fmt.Println("json error")
		return
	}

	//data 切片形式
	mes.Data = string(data)
	//6.mes 序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes json error")
		return
	}

	//7.发送消息，
	//7.1 先发送消息长度
	//conn.Write（）byte 切片
	//先得到长度，然后得到切片encoding/binary---简单数字与字节序列 PutUnit32([]byte,uint32)

	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if err != nil || n != 4 {
		fmt.Println("conn.Write(pkgLeg) fails")
		return
	}
	fmt.Println("send len succ:", len(data))
	return
}
