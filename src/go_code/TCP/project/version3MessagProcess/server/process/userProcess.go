package processor

import (
	"encoding/json"
	"fmt"
	utils "go_code/TCP/project/version3MessagProcess/common"
	"go_code/TCP/project/version3MessagProcess/common/message"
	"net"
)

type UserProcess struct {
	//字段分析
	//conn
	Conn net.Conn
}

func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//核心代码
	var loginmes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginmes)
	if err != nil {
		fmt.Println("serverProcessLogin json.Unmarshal fails:", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//申明一个loginRes类型
	var loginResMes message.LoginResMes

	//如果用户id为100，密码为123456认为合法，否则不合法
	if loginmes.UserId == 100 && loginmes.UserPwd == "123456" {
		//合法
		loginResMes.Code = 200
		loginResMes.Error = "ok"
	} else {
		//不合法
		loginResMes.Code = 500
		loginResMes.Error = "please register first"
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.marshal fails:", err)
	}
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.marshal fails:", err)
		return
	}

	//发送用函数来弄
	//因为使用mvc先创建一个transfer实例，然后再进行读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
