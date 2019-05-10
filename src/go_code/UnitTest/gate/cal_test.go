package main

import (
	"fmt"
	"testing" //引入go的testing框架包
)

//编写要给测试用例，去测试addUpper是否正确
//go test命令。能够自动执行如下形式的任何函数
//用于识别测试程序
func TestAddUpper(t *testing.T) {
	//调用
	res := AddUpper(10)
	if res != 55 {
		//fmt.Println("addUpper(10)执行错误，期望值=%v，实际值=%v", 55, res)
		//输出日志并且停止程序退出
		t.Fatalf("addUpper(10)执行错误，期望值=%v，实际值=%v", 55, res)
	}

	t.Logf("addUpper(10)执行正确，期望值=%v，实际值=%v", 55, res)
}

func TestHello(t *testing.T) {
	//调用
	fmt.Println("hello")
}

func TestGetSub(t *testing.T) {
	//调用
	res := GetSub(10, 3)
	if res != 13 {
		//fmt.Println("addUpper(10)执行错误，期望值=%v，实际值=%v", 55, res)
		//输出日志并且停止程序退出
		t.Fatalf("addUpper(10)执行错误，期望值=%v，实际值=%v", 13, res)
	}

	t.Logf("addUpper(10)执行正确，期望值=%v，实际值=%v", 13, res)
}
