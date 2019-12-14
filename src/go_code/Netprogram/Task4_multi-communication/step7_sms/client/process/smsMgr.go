package process

import (
	"encoding/json"
	"fmt"
	"go_code/Netprogram/Task4_multi-communication/step7_sms/common/message"
)

/*
	8.6接受群发消息
*/

func outputGroupMes(mes *message.Message) {
	//显示出来就好
	//这个地方mes一定smsmes
	//1. 反序列化
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	//3.显示信息，
	info := fmt.Sprintf("用户id:\t %d 对大家说：\t %s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)

}
