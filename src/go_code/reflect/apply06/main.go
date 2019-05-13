package main

import (
	"fmt"
	"reflect"
)

//编写一个cal结构体，有两个字段num1,num2
//方法GetSub(name string)
//使用反射遍历Cal结构体所有的字段信息
//使用反射机制完成对GetSub的调用，输出形式：tom完成了减法运算，8-3=5

type Cal struct {
	Num1 float64
	Num2 float64
}

func (this Cal) GetSub(name string) {
	fmt.Printf("%v 完成了 %v -%v =%v", name, this.Num1, this.Num2, this.Num1-this.Num2)
}

func (this Cal) GetSum(name string) {
	fmt.Printf("%v 完成了 %v -%v =%v", name, this.Num1, this.Num2, this.Num1+this.Num2)
}

func GetReflect() {

	//创建一个结构体
	var (
		cal *Cal
		st  reflect.Value
		ele reflect.Type
		str []reflect.Value
	)
	ele = reflect.TypeOf(cal)
	ele = ele.Elem()
	st = reflect.New(ele)

	cal = st.Interface().(*Cal)
	fmt.Println("st:", st)
	st = st.Elem()
	st.FieldByName("Num1").SetFloat(8.0)
	st.FieldByName("Num2").SetFloat(5.0)
	fmt.Println("st kind:", st.Kind().String())
	fmt.Println(cal)
	val := reflect.ValueOf(cal)
	//获取字段数目
	n := val.Elem().NumField()
	//n := val.NumField()
	fmt.Println("good")
	num := make([]reflect.Value, 2)
	for i := 0; i < n; i++ {
		num[i] = val.Elem().Field(i)
	}
	str = make([]reflect.Value, 1)
	str[0] = reflect.ValueOf("tom")
	val.Method(0).Call(str)

}

func main() {

	GetReflect()
}
