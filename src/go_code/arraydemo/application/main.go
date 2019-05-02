package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//1. 创建一个byte类型的26个元素数组，分别放置A-Z,使用for循环访问所有元素并且打印出来
	var arraystr [26]byte
	for i := 0; i < 26; i++ {
		arraystr[i] = 'A' + byte(i)
	}
	for _, v := range arraystr {
		fmt.Printf("%c ", v)
	}

	//2. 求一个数组的最大值

	var intArr = [...]int{1, -1, 9, 90, 11}
	maxTemp := intArr[0]
	Index := 0
	for i := 1; i < len(intArr); i++ {
		if maxTemp < (intArr)[i] {
			maxTemp = (intArr)[i]
			Index = i
		}
	}
	fmt.Println("index:", Index, "max:", maxTemp)

	//求一个数组的和和平均值

	var intArr2 = [5]int{1, -1, 9, 90, 11}
	sum := 0
	for _, v := range intArr2 {
		sum += v
	}
	fmt.Printf("sum:%v 平均值= %v\n", sum, float64(sum)/float64(len(intArr2)))

	//3.随机生成5个数，并加以反转
	//为了每次生成随机数不一样，我们需要给一个seed值
	rand.Seed((time.Now().UnixNano()))
	var intArr3 [5]int
	len := len(intArr3)
	fmt.Println("交换前")
	for i := 0; i < len; i++ {
		intArr3[i] = rand.Intn(100)
		fmt.Printf("%v ", intArr3[i])
	}
	for i := 0; i < len/2; i++ {
		temp := intArr3[i]
		intArr3[i] = intArr3[len-i-1]
		intArr3[len-i-1] = temp
	}
	fmt.Println("交换后")
	for _, v := range intArr3 {
		fmt.Printf("%v ", v)
	}

}
