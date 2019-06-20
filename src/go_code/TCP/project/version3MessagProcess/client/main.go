package main

import (
	"fmt"
)

//定义两个全局变量
//一个表示用户id
//一个表示用户密码
var userId int
var userPwd string

func display() {
	var input int
	var loop = true
	for loop {
		fmt.Println("---------------welcome-----------------------")
		fmt.Println("---------------1.login in-----------------------")
		fmt.Println("---------------2.register-----------------------")
		fmt.Println("---------------3.exit-----------------------")
		fmt.Println("---------------select(1-3)-----------------------")

		fmt.Scanln(&input)
		switch input {
		case 1:
			fmt.Println("1")
			loop = false
		case 2:
			fmt.Println("2")
			loop = false
		case 3:
			fmt.Println("3")
			loop = false
		default:
			fmt.Println("again")
		}
	}

	if input == 1 {
		fmt.Println("input id")
		//格式化
		fmt.Scanln(&userId)
		fmt.Println("input pwd")
		//格式化
		fmt.Scanln(&userPwd)
		//先把登陆函数写在另一个文件里面
		Login(userId, userPwd)
		// if err != nil {
		// 	fmt.Println("failure")
		// } else {
		// 	fmt.Println("success")
		// }
	} else {

	}
}

func main() {
	display()
}
