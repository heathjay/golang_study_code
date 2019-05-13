package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

//通过反射，修改int的一个值
//修改一个student的结构体的值

func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b)
	//reflect.SetString
	//rVal --类型是一个指针
	//type:reflect.Value kind:ptr2
	fmt.Printf("type:%T kind:%v\n", rVal, rVal.Kind())
	//rVal.SetInt(20)
	rVal.Elem().SetInt(20)
}

func main() {
	var num int = 2
	//panic: reflect: reflect.Value.SetInt using unaddressable value
	reflect01(&num) //地址拷贝
	fmt.Println("num:", num)
	//可以这样理解：
	// num := 9
	// ptr *int = &num
	// num2 := *ptr //rVal.Elem()
}
