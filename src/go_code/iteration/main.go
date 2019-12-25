package main

import (
	"fmt"
)

func test(num int) {
	if num > 2 {
		num--
		test(num)
	}
	fmt.Println("n=", num)
}
func main() {
	test(4)
}
