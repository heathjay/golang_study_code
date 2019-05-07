package main

import (
	"fmt"
)

//编写一个小学生考试系统

type Student struct {
	Name  string
	Age   int
	Score int
}

//将pupile和graduate共有方法也绑定到student

func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩= %v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	if score <= 100 && score >= 0 {
		stu.Score = score
	}
}

//给*Student增加一个方法，那么pupil 和 Graduate 都可以使用该方法
func (stu *Student) getSum(n1 int, n2 int) {
	fmt.Println("res:=", n1+n2)
}

type Pupile struct {
	Student //嵌入了匿名结构体，没有名字
}

//特有方法
func (p *Pupile) test() {
	fmt.Println("stu is testing")
}

type Gruaduate struct {
	Student
}

//特有方法
func (gru Gruaduate) testing() {
	fmt.Println("大学生正在考试中....")
}

//演示一下匿名字段是基本数据类型的使用
type E struct {
	int
}

func main() {

	//使用方法会发生改变
	pupile := &Pupile{}
	pupile.Student.Name = "tom"
	pupile.Student.Age = 19
	pupile.test()
	pupile.Student.SetScore(99)
	pupile.Student.ShowInfo()
	pupile.Student.getSum(33, 33)

	//访问简化，先看b中绑定否，再没有找匿名结构体中是否绑定
	graduate := &Gruaduate{}
	graduate.Name = "merry"
	graduate.Age = 29
	graduate.testing()
	graduate.SetScore(66)
	graduate.ShowInfo()

	graduatea := &Gruaduate{
		Student{"nihao", 19, 99},
	}

	graduatea.ShowInfo()
	var e E
	e.int = 20
	fmt.Println("e=", e)
}
