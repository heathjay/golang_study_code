package main

import (
	"fmt"
)

func feibo(n int) []uint64 {
	//接受一个n int
	// 能够将斐波那契数列放到切片
	var slice []uint64 = make([]uint64, n)
	slice[0] = 1
	slice[1] = 1
	for i := 2; i < n; i++ {
		slice[i] = slice[i-1] + slice[i-2]
	}
	return slice

}
func main() {
	slice := feibo(10)
	fmt.Println(slice)
}
