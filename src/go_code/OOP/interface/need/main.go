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
func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	//关键点
	//识别是哪一种，然后调用相关函数
	computer.Working(phone)
	computer.Working(camera)
}
