package process2

import (
	"encoding/json"
	"fmt"
	"go_code/Netprogram/Task4_multi-communication/step7_sms/common/message"
	"go_code/Netprogram/Task4_multi-communication/step7_sms/server/utils"
	"net"
)

/*
	8.5

*/

type SmsProcess struct {
	//...
}

//写方法
func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	//遍历服务器端的onlineMgr map【int】*UserProcess
	//将消息转发出去
	//取出message内容，SmsMes
	//可以拓展离线留言
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.unmarshal err=", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("mes json err:", err)
		return
	}
	for id, up := range userMgr.onlineUsers {
		//还需要过滤掉自己
		if id == smsMes.UserId {
			continue
		}
		this.SendToEachOnlineUser(data, up.Conn)
	}
}

//8.5写一个send方法
func (this *SmsProcess) SendToEachOnlineUser(data []byte, conn net.Conn) {
	//创建一个tf
	tf := &utils.Transfer{
		Conn: conn, //
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息:", err)
		return
	}
}
