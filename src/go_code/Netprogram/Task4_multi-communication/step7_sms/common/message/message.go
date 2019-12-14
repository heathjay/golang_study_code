package message

//定义消息常量
const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMesType"
)

//7.33定义几个用户在线的常量
const (
	UserOnline = iota
	UsreOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

//login message

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPw   string `json:"userPw"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code  int    `json:"code"`  //500 未注册 200登陆成功
	Error string `json:"error"` // 错误信息
	//7.3需要包含在线用户的切片，保存用户id的切片
	UserIds []int `json:"userIds"`
}

type RegisterMes struct {
	//6.1注册类型
	User User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"code"`  //400该用户已经占用， 200注册成功
	Error string `json:"error"` // 错误信息
}

//6.1为了配合服务器端推送上线的用户状态变化
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"` //用户的状态

}

//8.1增加一个smsMes//发送的
//SmsReMes可扩展
type SmsMes struct {
	Content string `json:"content"` //内容
	User           //匿名的结构体。继承的应用
}
