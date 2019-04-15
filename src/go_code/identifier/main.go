package main

//从gopath开始 忽略src 默认的
//写到go所在包
import (
	"fmt"
	"go_code/identifier/model"
)

func main() {

	var num int = 10
	var Num int = 20
	fmt.Printf("num=%v Num=%v\n", num, Num)

	// var ab c int = 30

	// var _ int = 400
	// fmt.Println(_)
	fmt.Println(model.HeroName)
}
