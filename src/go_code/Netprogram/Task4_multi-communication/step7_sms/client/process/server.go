package process

/*
8.1  - client 发送消息，
  - 写一个新的包，构建一个新的结构体，
  smsMes{
      User
      Content...//内容
  }
  - 客户端也需要维护一个连接，之前只是在用根本没有维护，用Conn去读，但没有保存
  - 在model层加一个内容，model下面新建一个结构体，
  CurUser{
    Conn
    User
  }
  - 添加一个smsProcessor.go

*/
import (
	"encoding/json"
	"fmt"
	"go_code/Netprogram/Task4_multi-communication/step7_sms/client/utils"
	"go_code/Netprogram/Task4_multi-communication/step7_sms/common/message"
	"net"
	"os"
)

//4.1显示登陆成功后的界面

func ShowMenu() {
	fmt.Println("---------恭喜xxx登陆成功----------")
	fmt.Println("---------1. 显示在线列表----------")
	fmt.Println("---------2. 发送消息----------")
	fmt.Println("---------3. 信息列表----------")
	fmt.Println("---------4. 退出系统----------")
	fmt.Println("请选择1-4")
	var key int

	var content string

	//8.4因为我们总会使用到smsProcess的实例，因此我们将其定义在switch外部
	smsProcess := &SmsProcess{}
	fmt.Scanln(&key)

	switch key {
	case 1:
		outputOnelineUser()
		fmt.Println("显示在线用户列表")
	case 2:

		fmt.Println("请输入你想说的话")
		fmt.Scanln(&content)
		//8.3实例化，只要连接不断，实例不会被销毁
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统") // 走之前可以先给server发一个消息，入库或者缓存
		os.Exit(0)
	default:
		fmt.Println("你输入的不对")
	}

}

//4.1和服务器端保持通讯
func serverProcessMes(Conn net.Conn) {
	/*创建一个transfer 实例， 不停地读取服务器的消息*/
	tf := &utils.Transfer{
		Conn: Conn,
	}
	for {
		fmt.Println("客户端,正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg() //会阻塞在这里
		if err != nil {
			fmt.Println("readPkg err=", err)
			return
		}

		//如果读取到了消息又是下一步处理逻辑
		//fmt.Println("mes = ", mes)
		//7.44
		switch mes.Type {
		case message.NotifyUserStatusMesType: //通知某人上线了
			//1.提取notifyMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updateUserStatus(&notifyUserStatusMes)
			//2. 把这个用户的信息，状态保存在客户的map中map[int] User
			//
		case message.SmsMesType:
			//有人群发消息
			outputGroupMes(&mes)
		default:
			//无法处理该类型
			fmt.Println("暂时无法识别该类型")
		}

	}
}
