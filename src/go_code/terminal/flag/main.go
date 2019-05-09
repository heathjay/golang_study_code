package main

import (
	"flag"
	"fmt"
)

func main() {
	//定义变量接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int
	//user用来接收用户命令行中输入的-u后面的参数
	//u代表 就是-u指定的参数
	//“默认值
	//对这个参数的说明
	flag.StringVar(&user, "u", "", "用户名默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码默认为空")
	flag.StringVar(&host, "h", "localhost", "主机默认为空")
	flag.IntVar(&port, "port", 3306, "端口默认为空")

	//非常重要的操作，就是转换，必须调用该方法
	flag.Parse()
	fmt.Printf("user=%v pwd=%v host=%v port=%v\n", user, pwd, host, port)

}
