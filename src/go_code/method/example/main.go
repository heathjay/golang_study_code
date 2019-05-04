package main

import (
	"fmt"
)

//面向对象编程的应用实例
//1.方法：
//a.声明结构体，确定结构体名称
//b.编写结构体的字段
//c.编写结构体的方法

//一个student结构体，包含name、gender、age、id、score字段分别为string、string、int、int、float64
//在main方法中，创建student 类型，访问方法say返回信息中包含的所有字段

type Student struct {
	Name   string
	Gender string
	Age    int
	Id     int
	score  float64
}

func (stu *Student) say() {
	fmt.Println("name=", stu.Name, "age=", stu.Age, "gender:=", stu.Gender, "id=", stu.Name, "score=", stu.score)
}

//盒子案例
//创建一个Box结构体，在其中声明三个字段表示一个立方体的长宽高，长宽高都要在终端获取
//声明一个方法获取立方体的体积
//创建一个Box结构体变量，打印给定尺寸的立方体体积

type Box struct {
	len   float64
	width float64
	hight float64
}

//一个景区根据游人的年龄收取不同价格的门票，比如年龄大于18，收费20元，其他情况门票免费
type Ticket struct {
	Name string
	Age  int
}

func (ti *Ticket) judgeTicket() int {
	var res int
	if ti.Age > 18 {
		res = 20
	} else {
		res = 0
	}
	return res
}
func (box *Box) getVol() float64 {
	return box.width * box.len * box.hight
}
func main() {
	var stu = Student{
		Name:   "jiang",
		Age:    19,
		Gender: "男",
		Id:     18,
		score:  99,
	}
	stu.say()

	var box Box
	box.len = 1.1
	box.width = 2.0
	box.hight = 3.0
	vol := box.getVol()
	fmt.Println("体积=", vol)

	for {
		var ti Ticket
		fmt.Println("input name")
		fmt.Scanln(&ti.Name)
		if ti.Name == "exit" {
			break
		}
		fmt.Println("input age")
		fmt.Scanln(&ti.Age)
		res := ti.judgeTicket()
		fmt.Println("游客名字为:", ti.Name, "游客年龄:", ti.Age, "价钱：", res)
	}

}
