package process2

import (
	"fmt"
)

/*
7.1
//因为UserMgr，实例在服务器端有且只有一个
//因为在很多地方都会使用到，因此将它定义为全局变量

*/

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//对userMgr初始化
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

//完成对onlineUsers添加
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

//完成删除
func (this *UserMgr) DeleteOnlineUser(userId int) {
	//map删除
	delete(this.onlineUsers, userId)
}

//查询,返回用户列表
func (this *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return this.onlineUsers
}

//传递过来一个id返回map的值，点对点通信
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	//如何从map中取出一个值带检测的方式
	up, ok := this.onlineUsers[userId]
	if !ok {
		//当前不在线的用户
		err = fmt.Errorf("用户%d不存在", userId)
		return
	}
	return
}
