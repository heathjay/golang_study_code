package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go_code/TCP/project/version2/common/message"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		//循环读取客户端发消息
		//直接封装成函数，返回一个message和一个err
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("client out,server up")
				return
			} else {
				fmt.Println("reading err=", err)
				return
			}

		}
		fmt.Println("mes=", mes)
	}
}

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
	return
}

func main() {
	fmt.Println("server is listening by port 8889")
	Listerner, err := net.Listen("tcp", "0.0.0.0:8889")
	defer Listerner.Close()
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

}
