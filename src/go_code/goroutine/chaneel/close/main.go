package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 3)
	intChan <- 10
	intChan <- 33
	close(intChan)
	//这时候不能往里面写入数据
	//intChan <- 300
	//读取ok
	num1 := <-intChan
	fmt.Println(num1)
	fmt.Println("oooll")

	//遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}
	//遍历
	//遍历时不能使用普通的for循环，
	close(intChan2)
	//如果没有close会出现死锁，但数据全部被取出
	for v := range intChan2 {
		fmt.Print("--,", v)
	}

}
