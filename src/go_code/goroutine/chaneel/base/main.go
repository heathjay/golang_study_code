package main

import (
	"fmt"
)

func main() {
	//演示一下管道的使用
	//1.创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Println(intChan)

	//本身地址
	fmt.Printf("本身地址:%p\n", &intChan)

	//向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	//管道的长度和cap
	//make多大就多大
	// intChan <- 99
	// intChan <- 100
	//注意：当我们给这个管道写入数据时，不能超过其容量，不然会出现死锁
	fmt.Printf("channel 长度：%v cap=%v\n", len(intChan), cap(intChan))
	//管道的价值在于一边取一边放

	//取数据

	var num2 int
	num2 = <-intChan
	fmt.Println(num2)

	//在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取会报告死锁
	fmt.Printf("channel 长度：%v cap=%v\n", len(intChan), cap(intChan))

}
