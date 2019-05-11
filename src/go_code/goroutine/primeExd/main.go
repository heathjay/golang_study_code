package main

import (
	"fmt"
	"runtime"
	"time"
)

//向intChan放入1-8000个数
func putNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	close(intChan) //		没有close也开始读了
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	//使用for循环

	var flag bool //标示是不是素数
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true //假设是一个素数
		//判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				//说明这个num不是素数
				flag = false
				break
			}
		}

		if flag {
			//将这个数让人到primeChan
			primeChan <- num
		}
	}
	//有一个
	fmt.Println("有一个prime协程取不到数据退出了")
	//还不能关闭primeChan，有可能别人还在处理，向exitChan写入一个true
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	cpus := runtime.NumCPU()
	fmt.Println("cpus:", cpus)
	exitChan := make(chan bool, cpus) //开到cpus大小

	start := time.Now().UnixNano()
	fmt.Println(start)
	//开启一个协程向intChan放入1-8000个数
	go putNum(intChan)
	//开启cpus个协程，从intChan取出数据，并判断是否为素数，如果是就放入到primeChan
	for i := 1; i <= cpus; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}
	go func() {
		for i := 0; i < cpus; i++ {
			<-exitChan
		}
		end := time.Now().UnixNano()
		fmt.Println(end)
		fmt.Println("耗时：", end-start)
		//当我们从exitChan取出了4个结果，就可以放心地关闭primeNUm
		close(primeChan)
	}()

	//遍历我们的primeNum,把结果取出来

	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
		//将结果输出
		// fmt.Println("素数有:", res)
	}

	fmt.Println("主线程退出")
}
