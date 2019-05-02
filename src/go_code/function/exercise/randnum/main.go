package main

import (
	"fmt"
	"math/rand"
	"time"
)

//编写一个函数，随机猜数字，随机生成一个1-100的整数，有10次机会，如果第一次就猜中，返回“good job”
//如果2--3次猜中，提醒“smart”
//如果4--9次猜中，提示“soso”
//最后一次猜中，提示“finally”
//如果一次没中，提示说"badjob“
func main() {
	rand.Seed((time.Now().Unix()))
	var num = rand.Intn(100)
	var input int
	var count int
	for i := 1; i <= 11; i++ {
		count = i
		fmt.Println("input your ", i, " try:")
		fmt.Scanln(&input)
		if input == num || count > 10 {
			break
		}
	}

	switch count {
	case 1:
		fmt.Println("goodjob")
	case 2, 3:
		fmt.Println("smart")
	case 4, 5, 6, 7, 8, 9:
		fmt.Println("soso")
	case 10:
		fmt.Println("soso")
	default:
		fmt.Println("badjob")
		fmt.Println("ans=", num)
	}
}
