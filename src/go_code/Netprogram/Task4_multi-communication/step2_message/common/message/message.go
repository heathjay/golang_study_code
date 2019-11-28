package message

//定义消息常量
const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
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
}
