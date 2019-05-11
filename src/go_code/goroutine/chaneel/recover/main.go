package main

import (
	"fmt"
	"time"
)

//函数
func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}

//函数
func test() {
	//这里可以使用错误处理机制，defer + recover
	defer func() {
		//捕获test跑出的panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()
	//定一个map
	var myMap map[int]string
	//直接是错的
	myMap[0] = "golang" //error

}
func main() {
	//一个协程出现panic
	//导致整个程序崩溃
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("main() say ok=", i)
	}
}
