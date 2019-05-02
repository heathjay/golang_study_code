package main

import (
	"errors"
	"fmt"
)

func test() {
	defer func() {
		//recover 是一个内置函数可以捕获到异常
		if err := recover(); err != nil {
			//说明捕获到错误
			fmt.Println("err=,", err)
			//发送邮件给某某管理员
			fmt.Println("管理员：错误出现了")
		}
	}()

	num1 := 100
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

//函数去读取以配置文件init.conf的信息
//如果文件名传入不正确，我们就返回一个自定义的错误。

func readConf(name string) (err error) {
	if name == "config.ini" {
		//读取
		return nil
	} else {
		//返回一个自定义的错误
		return errors.New("读取文件错误")
	}

}

func test02() {
	err := readConf("config2.ini")
	if err != nil {
		//如果读取文件发生错误，就输出这个错误，并且终止程序
		panic(err)
	}

	fmt.Println("test02继续执行....")
}
func main() {

	//  1. 即使错了，仍然继续处理下面
	//	2. 错误预警机制
	// test()
	// fmt.Println("main()下面的代码：")
	test02()
}
