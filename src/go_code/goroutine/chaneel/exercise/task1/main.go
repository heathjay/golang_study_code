package main

import "fmt"

//启动一个协程，将1-2000放入一个channel中，
//启动8个协程，从numChan取出数，并计算1+。。+n，放在存放的resChan
//最后8个协程协同完成工作后，再遍历ersChan，显示结果
//注意考入 resChan chan int 是否合适
func WriteData(numChan chan int) {
	for i := 0; i < 2000; i++ {
		fmt.Println("write:", i)
		numChan <- i
	}
	close(numChan)
}

func CalData(numChan chan int, resChan chan int, exitChan chan bool) {
	for {
		n, ok := <-numChan
		if !ok {
			break
		}
		fmt.Println("cal :", n)
		res := 0
		for i := 1; i <= n; i++ {
			res += i
		}
		resChan <- res
	}
	fmt.Println("有一个完成")
	exitChan <- true
}

func main() {
	numChan := make(chan int, 2000)
	exitChan := make(chan bool, 8)
	resChan := make(chan int, 2000)

	go WriteData(numChan)

	for i := 1; i <= 8; i++ {
		fmt.Println("起线程:", i)
		go CalData(numChan, resChan, exitChan)
	}

	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}

		//当我们从exitChan取出了8个结果，就可以放心地关闭primeNUm
		close(resChan)
		close(exitChan)
	}()
	i := 1

	for {
		res, ok := <-resChan
		if !ok {
			break
		}
		//将结果输出
		fmt.Println("第", i, "结果是:", res)
		i++
	}

	fmt.Println("主线程退出")
}
