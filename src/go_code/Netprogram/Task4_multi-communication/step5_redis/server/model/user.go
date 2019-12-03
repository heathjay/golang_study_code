package model

/*
	5.1定一个用户的结构体
		为了序列化成功，用户用来传输的字符串的key和结构体的字段对应的tag名字保持一致，否则会失败
		server 层可以用来调用Dao，如果有多个Dao就需要一个server层
*/

type User struct {
	//tag不要忘记
	UserId   int    `json:"userId"`
	UserPw   string `json:"userPw"`
	UserName string `json:"userName"`
}
