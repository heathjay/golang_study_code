package main

import "fmt"

func sum(n1 int, n2 int) int {
	defer fmt.Println("ok1 n1=", n1) //ok n1 = 10
	defer fmt.Println("ok2 n2=", n2) //ok2 n2 = 20
	//增加一句--相关值拷贝同时入栈
	n1++           //n1 = 11
	n2++           //n2 = 21
	res := n1 + n2 //res = 32
	fmt.Println("sum=", res)
	return res
}

func main() {
	sum(10, 20)
}
