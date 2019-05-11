package main

import (
	"fmt"
)

func main() {
	//使用select可以解决从管道内取数据的阻塞问题

	//1.定义一个管道10个数据int
	intChan := make(chan int, 10)

	for i := 0; i < 10; i++ {
		intChan <- i
	}

	//2.定一个管道5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hell" + fmt.Sprintf("%d", i)
	}

	//传统方法在遍历管道时，如果不关闭会阻塞而导致deadlock
	//确定什么时候关闭是一件很难的事情
	//label:
	for {
		//注意：如果管道最终没有关闭，不会一直阻塞而死锁。
		//自动到下一个case匹配
		select {
		case v := <-intChan:
			fmt.Printf("从intChan读取了数据%d\n", v)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取了数据%s\n", v)
		default:
			fmt.Println("都取不到了\n") //程序员可以加入业务逻辑了
			//break label            //return 代表退出该系统,尽量不用label
			return
		}
	}
}
