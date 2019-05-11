package main

import (
	"fmt"
)

func send(ch chan<- int, exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
	var a struct{}
	exitChan <- a
}

func recv(ch <-chan int, exitChan chan struct{}) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	var a struct{}
	exitChan <- a
}

func main() {
	// //	默认情况下，管道是双向
	// //声明为只写
	// var chan2 chan<- int
	// chan2 = make(chan int, 3)
	// chan2 <- 20
	// fmt.Println("chan2=", chan2)

	// //声明只读
	// var chan3 <-chan int
	// num2 := <-chan3
	// fmt.Println("num2", num2)

	//应用实例
	var ch chan int
	ch = make(chan int, 10)
	exitChan := make(chan struct{}, 2)
	go send(ch, exitChan)
	go recv(ch, exitChan)

	var total = 0
	for _ = range exitChan {
		total++
		if total == 2 {
			break
		}
	}

	fmt.Println("结束")
}
