package main

import (
	"fmt"
	utils "go_code/TCP/project/version3MessagProcess/common"
	"go_code/TCP/project/version3MessagProcess/common/message"
	processor "go_code/TCP/project/version3MessagProcess/server/process"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) severProcessMes(mes *message.Message) (err error) {
	//用switch比较好
	switch mes.Type {
	case message.LoginMesType:
		up := &processor.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		fmt.Println("register")
	default:
		fmt.Println("input again")
		return
	}
	return nil
}

func (this *Processor) processUnit() (err error) {
	for {
		//循环读取客户端发消息
		//直接封装成函数，返回一个message和一个err
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("client out,server up")
				return err
			} else {
				fmt.Println("reading err=", err)
				return err
			}

		}
		//fmt.Println("mes=", mes)

		err = this.severProcessMes(&mes)
		if err != nil {
			return err
		}
	}

}
