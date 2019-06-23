package process

//显示登陆成功的界面
//保持通讯
import (
	"fmt"
	utils "go_code/TCP/project/version3MessagProcess/common"
	"net"
	"os"
)

type ServerUint struct {
	UserName string
	UserId   int
	Conn     net.Conn
}

//显示登陆成功后的界面....

func (this *ServerUint) ShowMenu() {
	//1.显示在线用户列表
	//2.发送信息
	//3.信息列表
	//4.

	fmt.Println("-----------hello", this.UserName, "-----------")
	fmt.Println("-----------1. display online server---------")
	fmt.Println("-----------2. send message			---------")
	fmt.Println("-----------3. information list     ---------")
	fmt.Println("-----------4. exit					---------")
	fmt.Println("-----------select (1-4)			---------")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("show")
	case 2:
		fmt.Println("send")
	case 3:
		fmt.Println("list history")
	case 4:
		fmt.Println("bye")
		//可以告诉服务器我要走
		//可以
		os.Exit(0)
	default:
		fmt.Println("select again")
	}

}

func (this *ServerUint) serverProcessMes() {
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	for {
		fmt.Println(this.UserName, " is waiting for information")
		mes, err := tf.ReadPkg() //会进行阻塞
		if err != nil {
			fmt.Println("server down:", err)
			return
		}
		fmt.Println(mes)
	}
}
