package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go_code/TCP/project/version3MessagProcess/common/message"
	"net"
)

//把这些方法关联到结构体中
type Transfer struct {
	//分析有哪些字段
	//conn
	Conn net.Conn
	//buf,传输的时候当作切片
	Buf [1024 * 8]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 1024*4)
	//fmt.Println("读取客户端发送的数据")
	//conn.Read 在conn没有被关闭的情况才会组赛
	//如果客户端关闭了
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//fmt.Println("connect.read is error:", err)
		//err = errors.New("reading pkg header error")
		return
	}
	//	fmt.Println("good job,buf:= ", buf[0:4])

	//内容长度
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	//fmt.Println(pkgLen)
	//根据pkgLen读取内容
	//从conn里面读取放在buf内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		//err = errors.New("read n error")
		//fmt.Println("conn.read fail err=", err)
		return
	}
	//把pkgLen反序列成message类型
	//记得&mes
	v := this.Buf[0:pkgLen]
	err = json.Unmarshal(v, &mes)
	if err != nil {
		err = errors.New("json,Unmarshal err")
		//fmt.Println("json,Unmarshal err:", err)
		return
	}
	//fmt.Printf("%v", m)
	//err = severProcessMes(conn, &mes)
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	//长度
	pkgLen := uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	n, err := this.Conn.Write(buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fails:", err)
		return
	}

	n, err = this.Conn.Write(data)
	if uint32(n) != pkgLen || err != nil {
		fmt.Println("conn.Write(data) fails:", err)
		return
	}
	return
}
