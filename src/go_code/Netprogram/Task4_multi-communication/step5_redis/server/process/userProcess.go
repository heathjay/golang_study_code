package process2

import (
	"encoding/json"
	"fmt"
	"go_code/Netprogram/Task4_multi-communication/step5_redis/common/message"
	"go_code/Netprogram/Task4_multi-communication/step5_redis/server/utils"
	"net"
)

//4.2 UserProcessor 结构体
type UserProcessor struct {
	//分析字段
	Conn net.Conn
}

//3.2.2编写函数处理登陆请求,
//4.2.2 conn没有必要一直有，被UserProcessor 携带。
func (this *UserProcessor) ServerProcessLogin(mes *message.Message) (err error) {
	//3.2.2.1先从message中取出data，反序列化成loginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.unmarshall loginMes fails")
		return
	}

	//先声明一个response
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//填充数据
	var loginResMes message.LoginResMes

	//如果用户的id = 100, 密码 = 123456， 认为合法，否则不合法
	if loginMes.UserId == 100 && loginMes.UserPw == "123456" {
		//数据库验证
		loginResMes.Code = 200

	} else {
		//不合法
		loginResMes.Code = 500 //表示用户不合法
		loginResMes.Error = "invaliad user"
	}

	//序列化回复完成
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("unmarshall loginResMes fails:", err)
		return
	}

	//将data 赋值给resMes
	resMes.Data = string(data)

	//对整个信息包进行序列化发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json marshal resMes fails:", err)
		return
	}

	//发送data封装函数
	//4.2需要调用Transfer对象
	//因为使用了分层模式，先创建transfer实例，然后来读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("write data err:", err)
		return
	}
	return
}
