package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("-----start-----")
	fmt.Println(s)
	fmt.Println("------end-------")
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{}) {
	val := reflect.ValueOf(a)
	typ := reflect.TypeOf(a)
	fmt.Println(typ)
	kd := val.Kind()

	//判断是否是结构体
	if kd != reflect.Struct {
		fmt.Println("plz input struct")
		return
	}

	//获取字段的数量
	num := val.NumField()
	fmt.Println("num:", num)

	//变量结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("filed : %d: 值为=%v\n", i, val.Field(i))
		//获取到struct标签，注意需要通过reflect.Type来获取tag标签的值
		//获取标签就不能用val
		//type.Field()--返回structField---一个结构体里面有tag get(key) `json: "name"`
		//有些字段you 标签，有些没有标签
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("filed : %d: tag为=%v\n", i, tagVal)
		}
	}

	numMethod := val.NumMethod()
	fmt.Println("多少个方法:", numMethod)

	//调用方法
	//获取到第二个方法method(1)---进行方法排序的时候是按照函数的名字开始排序，ascii码
	//call调用该方法 call(in []value) reflect.value切片 返回一个[]Value
	val.Method(1).Call(nil)

	//声明一个[]reflect.Value切	片
	//调用的时候需要进行放入到切片的操作
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params)
	fmt.Println("len:", len(res), "res=", res[0].Int())
}

func main() {
	var a Monster = Monster{
		Name:  "xixi",
		Age:   400,
		Score: 20.0,
	}

	TestStruct(a)
}
