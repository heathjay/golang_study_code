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

func unmarshalStruct() {
	//说明str是读取到的，文件或者网络传输获得的
	str := "{\"Name\":\"jiang\",\"Age\":19,\"Skill\":\"speak\"}"

	//定一个结构体实例
	var stu Student
	err := json.Unmarshal([]byte(str), &stu)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println(stu)
	}
}

func unmarMap() {
	//说明str是读取到的，文件或者网络传输获得的
	str := "{\"age\":30,\"name\":\"红孩儿\",\"skill\":\"喷火\"}"

	//定义一个map
	var a map[string]interface{}
	//反序列化就不需要make，被封装到unmarshal里面
	//a = make(map[string]interface{})

	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println(a)
	}
}

func unmarSlice() {
	//说明str是读取到的，文件或者网络传输获得的
	str := "[{\"age\":19,\"name\":\"jake\"},{\"address\":[\"上海\",\"北京\"],\"pu\":191991}]"

	//定义一个slice
	var slice []map[string]interface{}
	//反序列化就不需要make，被封装到unmarshal里面
	//a = make(map[string]interface{})

	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println(slice)
	}
}

func main() {
	//将json字段改成map,struct，slice
	unmarshalStruct()
	unmarMap()
	unmarSlice()
}
