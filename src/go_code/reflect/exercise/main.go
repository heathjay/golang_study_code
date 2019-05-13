package main

import (
	"fmt"
	"reflect"
)

//1.给一个变量，请使用发射来得到他的reflect.Value,然后获取对应的Type，KInd和值。并将reflect.Value 转成interface{},再将interface{}转换成float64
func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b)

	rType := rVal.Type()
	rKind := rVal.Kind()
	fmt.Printf("rType:%v rKind:%v\n", rType, rKind)

	iVal := rVal.Interface()
	intVal := rVal.Float()
	fmt.Printf("intVal:%v,intVal type:%T\n", intVal, intVal)
	num := iVal.(float64)
	fmt.Printf("num:%v,num:%T\n", num, num)

}

//2.

func main() {
	var v float64 = 1.2
	reflect01(v)

	var str string = "tom"
	// fs := reflect.ValueOf(str)
	//panic: reflect: reflect.Value.SetString using unaddressable value
	//fs.SetString("jack")
	fs := reflect.ValueOf(&str)
	fs.Elem().SetString("jac")
	fmt.Printf("fs:%v\n", str)
}
