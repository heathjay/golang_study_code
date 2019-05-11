package main

import (
	"fmt"
	"strconv"
	"time"
)

// - 在主线程中开启一个goroutine，改协程每隔1s输出“hello world”
// - 在主线程中也每个一秒输出
// - 要求主线程和协程同时进行

//编写一个函数，每隔1s输出hello, world
func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello world" + strconv.Itoa(i+1))
		time.Sleep(time.Second)
	}
}

func main() {

	//如果直接调用？
	//test()顺序执行
	go test() //非常快起了线程

	for i := 0; i < 10; i++ {
		fmt.Println("hello golang" + strconv.Itoa(i+1))
		time.Sleep(time.Second)
	}
}
