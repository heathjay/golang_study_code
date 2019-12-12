package process2

import (
	"encoding/json"
	"fmt"
	"go_code/Netprogram/Task4_multi-communication/step6_displayonline/common/message"
	"go_code/Netprogram/Task4_multi-communication/step6_displayonline/server/model"
	"go_code/Netprogram/Task4_multi-communication/step6_displayonline/server/utils"
	"net"
)

//4.2 UserProcessor 结构体
type UserProcess struct {
	//分析字段
	Conn   net.Conn
	UserId int
}

//7.22编写通知所有在线用户的一个方法
//userid要通知其他的在线用户我上线了，遍历UserMgr，然后一个个发送，然后一个一个
func (this *UserProcess) NotifyOthersOnLineUser(userId int) {
	for id, up := range userMgr.onlineUsers {
		//过滤掉自己
		if id == userId {
			continue
		}
		//开始通知【单独写一个方法】
		up.NotifyMeOnline(userId)
	}

}

//7.33
func (this *UserProcess) NotifyMeOnline(userId int) {
	//组装我们的消息NotifyUserStatusMes
	//同一个代码有很多个人接受可以提出去
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		println("notifyonline json err:", err)
		return
	}
	//将序列化后的mes赋给data
	mes.Data = string(data)
	//将mes再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("notifyonline json err:", err)
		return
	}

	//发送创建一个transfer
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("write err", err)
		return
	}
}

//6.3register方法
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//取出message的data反序列化成RegiterMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.unmarshall registerMes fails")
		return
	}
	//先声明一个response
	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	//填充数据
	var registerResMes message.RegisterResMes
	//去数据库完成注册
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	} else {
		registerResMes.Code = 200
	}
	//序列化回复完成
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("unmarshall registerResMes fails:", err)
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
	return

}

//3.2.2编写函数处理登陆请求,
//4.2.2 conn没有必要一直有，被UserProcessor 携带。
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
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
	/*
		5.4需要到redis数据库去完成验证
			1. 使用model.MyUserDao
	*/

	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPw)

	if err != nil {
		//5.4先测试成功，然后根据具体的错误信息
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500 //表示用户不合法
			loginResMes.Error = err.Error()

		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}

	} else {
		loginResMes.Code = 200
		//7.1这里因为用户已经登陆成功于是我们就把该登陆成功的用户放入到userMgr中
		//将登陆成功的用户的id赋给this
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		//7.2通知在看
		this.NotifyOthersOnLineUser(loginMes.UserId)
		//7.3将当前在线用户的id放入到loginResMes.UserIds字段中
		//进行遍历userMgr中的onlineUsers
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}
		fmt.Println(user, " login sucessfully")
	}

	//如果用户的id = 100, 密码 = 123456， 认为合法，否则不合法
	// if loginMes.UserId == 100 && loginMes.UserPw == "123456" {
	// 	//数据库验证
	// 	loginResMes.Code = 200

	// } else {
	// 	//不合法
	// 	loginResMes.Code = 500 //表示用户不合法
	// 	loginResMes.Error = "invaliad user"
	// }

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
