package main

import (
	"fmt"
	"sync"
)

//要计算1-200各个数的阶乘，并且把各个数的阶乘放在map中，最后显示出来，要求使用goroutine完成
//1.编写一个函数，来计算各个数的阶乘并放入到一个map中，
//2.启动协程多少，统计的结果放入map中国呢
//3.map应该是一个全局变量

var (
	myMap = make(map[int]int)
	//声明一个全局的互斥锁
	//synchronize
	lock sync.Mutex
)

//test函数就是l计算n！，放入到mymap
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//加锁
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}
func main() {
	//这里开启多个协程完成任务，
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	//time.Sleep(time.Second * 10)
	//这里我们输出结果，放入到map中
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%v]=%v", i, v)
	}
	lock.Unlock()
}
