package main

import (
	"fmt"
	"go_code/Netprogram/Task4_multi-communication/step5_redis/common/message"
	process2 "go_code/Netprogram/Task4_multi-communication/step5_redis/server/process"
	"go_code/Netprogram/Task4_multi-communication/step5_redis/server/utils"
	"io"
	"net"
)

//4.3 剥离processor
type Processor struct {
	Conn net.Conn
}

//3.2函数优化，根据客户端发送种类不同，决定调用哪个函数来处理
func (this *Processor) serverProcess(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登陆信息
		//4.3.1调用下一层逻辑
		up := &process2.UserProcessor{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}

func (this *Processor) processor2() (err error) {
	for {

		//conn.Read 只有在conn没有被关闭的情况下， 才会阻塞,如果一方关闭了conn。则不会阻塞
		//3.1将读取消息直接封装成一个函数,
		//3.1 readPkg 返回一个message,err
		//4.3创建一个tf实例，用来读包
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("readPkg fails:", err)
			//连接关闭，server也退出
			if err == io.EOF {
				fmt.Println("conn disconnect")
				return err
			} else {
				fmt.Println("readPkg err:", err)
				return err
			}

		}
		err = this.serverProcess(&mes)
		if err != nil {
			fmt.Println("err:", err)
			return err
		}
	}
}
