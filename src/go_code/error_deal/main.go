package main

import (
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
func main() {

	//  1. 即使错了，仍然继续处理下面
	//	2. 错误预警机制
	test()
	fmt.Println("main()下面的代码：")

}
