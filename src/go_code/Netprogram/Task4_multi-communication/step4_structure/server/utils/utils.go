package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go_code/Netprogram/Task4_multi-communication/step4_structure/common/message"
	"net"
)

//4.1将方法关联到结构体中
type Transfer struct {
	//分析需要什么字段
	//1. 连接
	//2. 缓冲
	//暂时不考虑封装
	Conn net.Conn
	Buf  [8064]byte //传输时，使用的缓冲

}

//3.1封装成函数, readPkg
//3.2优化封装结构,需要readPkg可以读取不同的消息接口
func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//获取长度
	//buf := make([]byte, 8096)
	fmt.Println("wait for the message from the sender")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//fmt.Println("len sent error:", err)
		//自定义error
		err = errors.New("Read len fails.")
		return
	}
	//fmt.Println("read :", buf[:4])

	//根据buf[:4]判断到底读多少字节
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])

	//根据pkgLen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen]) //从conn读取总共pkgLen放进buf中
	if n != int(pkgLen) || err != nil {
		//fmt.Println("read LoginMes fails")
		//err = errors.New("read body fails")
		return
	}

	//反序列化 -> message.Message --> &mes！！
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println(" LoginMes unmarshall fails")
		return
	}

	return
}

//3.2.3工具函数
func (this *Transfer) WritePkg(data []byte) (err error) {
	//先发送一个长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)
	n, err := this.Conn.Write(this.Buf[:4])
	if err != nil || n != 4 {
		fmt.Println("conn.Write(pkgLeg) fails")
		return
	}

	//发送内容
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn write bytes fails:", err)
		return
	}
	return
}
