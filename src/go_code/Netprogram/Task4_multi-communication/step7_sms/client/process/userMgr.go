package process

import (
	"fmt"
	"go_code/Netprogram/Task4_multi-communication/step7_sms/client/model"
	"go_code/Netprogram/Task4_multi-communication/step7_sms/common/message"
)

//7.44客户端需要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser

//初始化工作在登陆成功后

//7.5专门写一个方法来处理返回的信息
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	//判断是否原先已经存在user
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		//原来没有
		user = &message.User{
			UserId:     notifyUserStatusMes.UserId,
			UserStatus: notifyUserStatusMes.Status,
		}
		onlineUsers[notifyUserStatusMes.UserId] = user
	} else {
		user.UserStatus = notifyUserStatusMes.Status
		onlineUsers[notifyUserStatusMes.UserId] = user
	}
	outputOnelineUser()
}

//7.66显示当前在线的用户，在客户端显示当前在线的用户
func outputOnelineUser() {
	//遍历一把就可以了
	//当前用户所在列表是：
	fmt.Println("当前在线列表是：")
	for id, _ := range onlineUsers {
		//如果不显示自己可以进行过滤
		fmt.Println("用户id:\t，", id)
	}
}
