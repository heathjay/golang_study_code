package main

import "fmt"

// //定义全局变量
// var n1 = 100
// var n2 = 200
// var name = "jack"

// //也可以改成一次性声明
// var (
// 	n3    = 300
// 	n4    = 900
// 	name2 = "mary"
// )

func main() {

	// //golang的变量使用方式
	// //第一种，指定变量类型，声明后若不赋值，使用默认值
	// //int默认值0
	// var i int
	// fmt.Println("i=", i) //加在一起

	// //第二种，根据值自行判断（类型推导）
	// var num = 10.11
	// fmt.Println("num=", num)

	// //第三种：省略var，注意：=左侧的变不应该是已经声明郭娥，否则导滞编译错误
	// //下面的方式等价于var name string   name = "tom"
	// //:不能省略
	// name := "tom"
	// fmt.Println("name =", name)

	// //一次性定义多个变量
	// var n1, n2, n3 int
	// fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	// //一次性声明多个变量2
	// var n4, nam, n5 = 100, "tom", 888
	// fmt.Println("n4=", n4, "nam=", nam, "n5=", n5)

	//一次性声明多个变量3，同样可以使用类型推导
	//n1, name, n3 := 100, "tom", 888
	// fmt.Println("n1=", n1, "name=", name, "n2=", n2)
	// fmt.Println("n3=", n3, "name2=", name2, "n4=", n4)

	//变量使用的注意事项
	var i int = 10
	i = 30
	i = 50
	j := 10
	var r = i + j
	fmt.Println("r=", r)
	fmt.Println("i=", i)

	//i = 1.2 //int 类型不能表示小数
	//var i int = 99  i := 59//重复定义

	//加号使用
	var str1 = "hello"
	var str2 = "world"
	var res = str1 + str2
	fmt.Println("res=", res)

	//常量

	var num int
	//必须给值
	const tax int = 0
	//不能进行修改
	// tax = 10
	//只能修饰基本数据类型
	//
	fmt.Println(num, tax)

	// const name = "tome" //ok
	// const tax float64 = 0.08 //ok
	// const a int //ok
	// const b1 = 9/33//bad
	// const c = getVal()//bad

	//简洁写法
	const (
		a = 1
		b = 1
	)
	//专业写法
	const (
		a1 = iota //0
		b1
		c
	)

}
