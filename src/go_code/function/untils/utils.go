package utils

import "fmt"

//调用前引入
var Age int
var Name string

func Cal(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '/':
		res = n1 / n2
	case '*':
		res = n1 * n2
	default:
		fmt.Println("操作符有错")
	}

	return res
}

func init() {
	Age = 100
	Name = "tom"
	fmt.Println("utils 包。。。。")
}
