package main

import (
	"fmt"
)

type BirdAble interface {
	flying()
}

type FishAble interface {
	swim()
}

//Monkey结构体
type Monkey struct {
	Name string
}

func (this *Monkey) climing() {
	fmt.Println(this.Name, "生来就会爬树...")
}

type LittleMonkey struct {
	Monkey
}

func (this LittleMonkey) flying() {
	fmt.Println(this.Name, "通过学习会飞翔")
}
func (this LittleMonkey) swim() {
	fmt.Println(this.Name, "通过学习会游泳")
}

func main() {

	//特别的猴子希望有功能扩展，但他不希望破坏合理性，即实例的时候进行功能性拓展
	var lm = LittleMonkey{
		Monkey{
			Name: "孙悟空",
		},
	}
	lm.climing()
	lm.flying()
	lm.swim()
}
