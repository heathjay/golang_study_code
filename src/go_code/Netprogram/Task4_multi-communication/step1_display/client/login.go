package main

import (
	"fmt"
)

//尽量用error的形式返回，
//error可以返回各种错误，自定义，描述信息
//变量尽量使用一样的
func login(userId int, userPw string) (err error) {
	//开始定协议
	fmt.Printf("userId= %d,userPw = %s", userId, userPw)
	return nil
}
