/*
	step 4:
		main.go: 1. 根据用户的选择，调用对应的处理器
		uerProcessor.go ： 1. 处理和用户相关的业务包括登陆和注册
		smsProcessor.go: 2.处理和短消息相关的逻辑， 私聊和群发
		utils.go ： 可以抽取出来一起使用，或者再写一类，公用的需要保证同步，或者上线前进行整合
		server.go : 客户端登陆后不会马上退出，1. 显示登陆成功的命令，2. 上线后，保持和服务器端的通信， 启动一个协程不停地与服务器进行交互 3. 当读取到服务器端发送的消息后，就会显示在界面上。
		客户端一旦退出，需要保留的东西，

*/

package main

import (
	"fmt"
	"go_code/Netprogram/Task4_multi-communication/step4_structure/client/process"
)

//用户id 和密码
var userId int
var userPw string

func main() {
	//接受输入
	var key int
	//判断是否还继续显示菜单
	//var loop = true
	//
	for true {
		fmt.Println("welcome......")
		fmt.Println("\t\t\t 1. login in")
		fmt.Println("\t\t\t 2. register")
		fmt.Println("\t\t\t 3. exit")
		fmt.Println("\t\t\t select")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			//fmt.Println("login in successfully")
			fmt.Println("请输入用户的id号码")
			fmt.Scanln(&userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanln(&userPw)
			//登陆写在另一个文件,写到另一个文件
			//设计一个message协议
			//loop = false

			//4.1完成登陆
			//4.1创建一个UserProcess 的实例
			up := &process.UserProcess{}
			up.Login(userId, userPw)
		case 2:
			fmt.Println("registeration ")
			//loop = false
		case 3:
			fmt.Println("bye")
			//loop = true
		}
	}
	//根据用户输入，显示新的信息

	// if key == 1 {
	// 	//说明用户要登陆
	// 	fmt.Println("请输入用户的id号码")
	// 	fmt.Scanln(&userId)
	// 	fmt.Println("请输入用户的密码")
	// 	fmt.Scanln(&userPw)
	// 	//登陆写在另一个文件,写到另一个文件
	// 	//设计一个message协议

	// 	//4.1 重新调用
	// 	login(userId, userPw)
	// 	// if err != nil {
	// 	// 	fmt.Println("login error")
	// 	// } else {
	// 	// 	fmt.Println("login succ")
	// 	// }
	// } else {
	// 	fmt.Println("register")
	// }
}
