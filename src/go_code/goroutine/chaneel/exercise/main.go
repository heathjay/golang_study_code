package main

import (
	"fmt"
)

func ReadData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("readdata :=", v)
	}
	exitChan <- true
	close(exitChan)
}

func WriteData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("write data:", i)
	}
	close(intChan) //关闭
}

func main() {
	//创建两个管道
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go ReadData(intChan, exitChan)
	go WriteData(intChan)
	//如果没做同步，就会秒退
	//协程销毁

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

}
