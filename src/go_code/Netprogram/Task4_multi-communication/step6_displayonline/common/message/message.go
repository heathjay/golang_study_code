package message

//定义消息常量
const (
	LoginMesType       = "LoginMes"
	LoginResMesType    = "LoginResMes"
	RegisterMesType    = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
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
