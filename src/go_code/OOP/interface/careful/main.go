package main

import "fmt"

type AInterface interface {
	Say()
}

type Student struct {
	Name string
}

func (stu Student) Say() {
	fmt.Println("stu say")
}

//2.不一定是struct来实现interface

type integer int

func (in integer) Say() {
	fmt.Println("integer saying,i:", in)
}

//3.一个自定义类型可以实现多个接口
type BInterface interface {
	Hello()
}

type Mix struct {
	Name string
}

func (m Mix) Hello() {
	fmt.Println("M say hello")
}

func (m Mix) Say() {
	fmt.Println("m say good")
}

//4.继承接口
type CInterface interface {
	AInterface
	BInterface
	ctest()
}

type Boy struct {
}

func (b Boy) Say() {

}
func (b Boy) Hello() {

}
func (b Boy) ctest() {

}

//5.空接口的实现
type T interface {
}

func main() {

	//1.不能创建一个接口实例，但能指向一个实现该接口的自定义类型的实例
	var stu Student
	var a AInterface = stu
	var i integer = 10
	var b AInterface = i

	a.Say()
	b.Say()

	//3.mix实现了ainterface 和 binterface
	var m Mix = Mix{"jiang"}
	var a2 AInterface = m
	var b2 BInterface = m
	a2.Say()
	b2.Hello()

	//4.接口可以实现继承
	var boy Boy
	var cboy CInterface = boy
	cboy.Hello()

	//5.空接口
	var t T = stu
	fmt.Println(t)
	var t2 interface{} = stu
	var num1 float64 = 8.8
	t2 = num1
	fmt.Println(t2)
}
