package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string
}

func testString() {
	//将结构体，map,切片, 进行序列化
	stu := Student{
		Name:  "jiang",
		Age:   19,
		Skill: "speak",
	}
	//func Marshal(v interface{})([]byte,error) 任何都可以序列化
	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("err:=", err)
	}

	//输出序列化后的结
	//字段的值可能被当作key--其值为value
	//但是可能会有小写，利用tag标签进行来标志，序列化后的key值按照设计的显示，
	fmt.Printf("序列化后：%v\n", string(data))
}
func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["skill"] = "喷火"
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("err:=", err)
	}
	fmt.Printf("序列化后：%v\n", string(data))
}
func testSlice() {
	//定义切片
	var slice []map[string]interface{}
	//多个map
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "jake"
	m1["age"] = 19

	m2 := make(map[string]interface{})
	a1 := [2]string{"上海", "北京"}
	m2["address"] = a1
	m2["pu"] = 191991

	slice = append(slice, m1, m2)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("err:=", err)
	}

	//输出序列化后的结
	fmt.Printf("序列化后：%v\n", string(data))
}

func main() {
	testString()
	testMap()
	testSlice()
}
