// step3: 工具包
package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go_code/Netprogram/Task4_multi-communication/step4_structure/common/message"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	//获取长度
	buf := make([]byte, 8096)
	fmt.Println("wait for the message from the sender")
	_, err = conn.Read(buf[:4])
	if err != nil {
		//fmt.Println("len sent error:", err)
		//自定义error
		err = errors.New("Read len fails.")
		return
	}
	//fmt.Println("read :", buf[:4])

	//根据buf[:4]判断到底读多少字节
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])

	//根据pkgLen读取消息内容
	n, err := conn.Read(buf[:pkgLen]) //从conn读取总共pkgLen放进buf中
	if n != int(pkgLen) || err != nil {
		//fmt.Println("read LoginMes fails")
		//err = errors.New("read body fails")
		return
	}

	//反序列化 -> message.Message --> &mes！！
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println(" LoginMes unmarshall fails")
		return
	}

	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	//先发送一个长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if err != nil || n != 4 {
		fmt.Println("conn.Write(pkgLeg) fails")
		return
	}

	//发送内容
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn write bytes fails:", err)
		return
	}
	return
}
