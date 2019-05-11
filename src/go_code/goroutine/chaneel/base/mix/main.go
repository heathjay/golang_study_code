package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	//struct 管道
	cat1 := Cat{Name: "goo", Age: 19}
	var map1 map[string]string
	map1 = make(map[string]string, 10)
	map1["good"] = "bad"
	map1["first"] = "second"

	var mixChan chan interface{}
	mixChan = make(chan interface{}, 10)
	mixChan <- cat1
	mixChan <- map1
	mixChan <- 10

	a := <-mixChan

	//报错的原因在于interface 需要类型断言
	//fmt.Println(a.Name)不能通过编译器
	b := a.(Cat)
	fmt.Println(b.Name)
}
