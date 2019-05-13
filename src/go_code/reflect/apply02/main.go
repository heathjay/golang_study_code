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

func TestStruct(b interface{}) {
	typ := reflect.TypeOf(b)
	val := reflect.ValueOf(b)
	numField := val.Elem().NumField()
	val.Elem().Field(0).SetString("jiajijajia")
	for i := 0; i < numField; i++ {
		fmt.Printf("filed : %d: 值为=%v\n", i, val.Elem().Field(i))
		tagVal := typ.Elem().Field(i).Tag.Get("json")
		fmt.Printf("tag= %s\n", tagVal)
	}

}

func main() {
	var monster = Monster{
		Name:  "jiba",
		Age:   199,
		Score: 99,
	}
	TestStruct(&monster)
	fmt.Println(monster)

}
