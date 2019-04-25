package main

import (
	"fmt"
)

func test() {
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
