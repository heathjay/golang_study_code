package main

import (
	"fmt"
	"strings"
)

func AddUpper() func(int) int {
	var n int = 10
	var str = "hello"
	return func(x int) int {
		n = n + x
		str += "a"
		fmt.Println("str= ", str) //str1="hello"
		return n
	}
}

func makeSuffix(suffix string) func(string) string {

	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {

	//使用前面的代码
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))

	// 1.makeSuffix(suffix string)可以接受一个文件后缀名.jpg，并返回一个闭包
	// 2.调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀，
	// 则返回文件名.jpg 如果已经有.jpg后缀，则返回原文件名
	// 3.要求使用闭包完成
	// 4.string.HasSuffix
	//返回一个闭包
	f2 := makeSuffix(".jpg")
	fmt.Println("文件名处理后：", f2("winter"))
	fmt.Println("有后缀：", f2("bird.jpg"))

}
