package main

import "fmt"

//全局匿名函数

var (
	func1 = func(n1, n2 int) int {
		return n1 + n2
	}
)

func main() {
	//定义匿名函数时就直接调用----只能调用一次
	//求两个数的和

	res := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res:=", res)

	//将匿名函数直接赋值给一个变量---可以多次

	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	//可以通过a完成调用
	res2 := a(10, 8)
	fmt.Println("res2:=", res2)

	//全局匿名函数的接收
	res4 := func1(4, 9)
	fmt.Println("res4:", res4)

	//闭包的基本介绍
	//闭包就是一个函数和与其相关的引用环境组合的一个整体

}
