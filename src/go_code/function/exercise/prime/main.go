package main

import (
	"fmt"
)

//编写一个函数，输出100以内所有素数，每行显示5个

func primeCount(limitation int) {
	fmt.Println("primeCount below than ", limitation)
	var count = 0
	var sum = 0
	for i := 1; i <= 100; i++ {
		var flag = 1
		for j := 1; j < i; j++ {
			if i%j == 0 && j != 1 {
				flag = 0
				break
			}
		}
		if flag == 1 {
			count++
			sum += i
			fmt.Print(i, " ")
			if count%5 == 0 {
				fmt.Print("\n")
			}
		}
	}
	fmt.Println("sum:= ", sum)
}
func main() {
	primeCount(100)
}
