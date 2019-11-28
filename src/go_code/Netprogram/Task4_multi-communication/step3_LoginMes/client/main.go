package main

import (
	"fmt"
)

//用户id 和密码
var userId int
var userPw string

func main() {
	//接受输入
	var key int
	//判断是否还继续显示菜单
	var loop = true
	//
	for loop {
		fmt.Println("welcome......")
		fmt.Println("\t\t\t 1. login in")
		fmt.Println("\t\t\t 2. register")
		fmt.Println("\t\t\t 3. exit")
		fmt.Println("\t\t\t select")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			//fmt.Println("login in successfully")
			loop = false
		case 2:
			fmt.Println("registeration ")
			loop = false
		case 3:
			fmt.Println("bye")
			loop = true
		}
	}
	//根据用户输入，显示新的信息

	if key == 1 {
		//说明用户要登陆
		fmt.Println("请输入用户的id号码")
		fmt.Scanln(&userId)
		fmt.Println("请输入用户的密码")
		fmt.Scanln(&userPw)
		//登陆写在另一个文件,写到另一个文件
		//设计一个message协议
		login(userId, userPw)
		// if err != nil {
		// 	fmt.Println("login error")
		// } else {
		// 	fmt.Println("login succ")
		// }
	} else {
		fmt.Println("register")
	}
}
