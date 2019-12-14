package process

//8.3群发
//发送群聊的消息

import (
	"encoding/json"
	"fmt"
	"go_code/Netprogram/Task4_multi-communication/step7_sms/client/utils"
	"go_code/Netprogram/Task4_multi-communication/step7_sms/common/message"
)

type SmsProcess struct {
}

//8.3群发
//发送群聊的消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {

	//1.创建一个mes实例
	var mes message.Message
	mes.Type = message.SmsMesType

	//2.创建一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//3.序列化
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("sendGroup序列化失败:", err)
		return
	}

	mes.Data = string(data)
	//4.对mes再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("sendGroup序列化失败:", err)
		return
	}
	//5.发送给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("sendGroup err:", err)
		return
	}
	return
}
