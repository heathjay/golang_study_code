package main

import (
	"fmt"
	"time"
)

func main() {

	flag := false
	start := time.Now().UnixNano()
	num := 8000
	for i := 1; i <= num; i++ {
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
			//primeChan <- num
		}
	}
	end := time.Now().UnixNano()
	fmt.Println(end)
	fmt.Println("耗时：", end-start)
}
