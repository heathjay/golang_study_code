package main

import (
	"fmt"
)

type Student struct {
}

//编写一个函数用来判断输入的参数是什么类型
func TypeJudge(items ...interface{}) {

	for i, x := range items {
		//type 只能在switch里面进行实现
		//
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型，值是%v\n", i, x)
		case float32:
			fmt.Printf("第%v个参数是float32类型，值是%v\n", i, x)
		case float64:
			fmt.Printf("第%v个参数是float64类型，值是%v\n", i, x)
		case int, int32, int64:
			fmt.Printf("第%v个参数是int类型，值是%v\n", i, x)
		case string:
			fmt.Printf("第%v个参数是string类型，值是%v\n", i, x)
		case Student:
			fmt.Printf("第%v个参数是student类型，值是%v\n", i, x)
		case *Student:
			fmt.Printf("第%v个参数是*student类型，值是%v\n", i, x)
		default:
			fmt.Println("不确定\n")
		}
	}
}

func main() {
	var n1 float32 = 1.1
	var n2 float64 = 30.0
	var n3 int32 = 88
	var name string = "good"
	address := "上海"
	n4 := 300

	stu1 := Student{}
	stu2 := &Student{}

	TypeJudge(n1, n2, n3, name, address, n4, stu1, stu2)
}
