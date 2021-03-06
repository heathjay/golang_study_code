package process

import (
	"fmt"
	"go_code/Netprogram/Task4_multi-communication/step5_redis/client/utils"
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
	fmt.Scanln(&key)

	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("发送消息")
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
		fmt.Println("mes = ", mes)

	}
}
