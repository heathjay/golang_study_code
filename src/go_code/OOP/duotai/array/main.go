package main

import (
	"fmt"
)

//定一个接口
type Usb interface {
	Start()
	Stop()
}

//结构体
type Phone struct {
	Name string
}

//让phone实现usb接口
func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

//让相机Camera来实现Usb方法
type Camera struct {
	Name string
}

func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

//Computer
type Computer struct {
}

//能接收一个interface接口类型的变量
//只要是实现了Usb接口
func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

type Point struct {
	x int
	y int
}

func main() {
	//定义一个usb接口数组，可以存放phone和camera结构体变量
	//这里就体现出多态数组
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Camera{"ke"}
	usbArr[2] = Camera{"sony"}
	fmt.Println(usbArr)

	var a interface{}
	var point = Point{1, 2}
	a = point
	//把a重新给一个另一个point变量
	var b Point
	//需要判断a到底是谁
	//b = a.(Point)
	b = a.(Point)
	fmt.Println(b)

	var x interface{}
	var b1 float32 = 1.1
	x = b1 // 把b赋值给x
	//空接口可以赋给任何类型
	//x ==>float32【使用类型断言】
	//带检测的类型断言,成功true，不成功false
	// y, err := x.(float32)
	if y, err := x.(float32); !err {

		fmt.Println("convert fail")
	} else {
		fmt.Printf("y的类型是 %T 他的值是%v\n", y, y)
	}

	fmt.Println("继续执行")

}
