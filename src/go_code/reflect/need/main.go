package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

type Monster struct {
	Name string
	Age  int
}

//1. 编写一个案例，演示对基本类型、interface{}，Value进行操作
func reflectTest01(b interface{}) {
	//通过num进行获取传入变量的type，kind value
	//1.:先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	//他真正的类型是reflect.Type,
	//原先的int不能调用东西的，
	fmt.Println("type:", rTyp)

	//2.获取到reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("reflect:", rVal)

	//3.获取kind
	//rVal
	fmt.Println("kind1:", rVal.Kind())
	//rType
	fmt.Println("kind:", rTyp.Kind())

	//但他不是真正原先的100
	//如果是100可以进行下面的操作：
	// n2 := 2 + rVal//mismatch int ---reflect.Value 类型
	//如何真正拿到int的值？
	n2 := 2 + rVal.Int()
	fmt.Println(n2)
	fmt.Printf("rVal=%v rVal type= %T\n", rVal, rVal)
	//转成interface
	iV := rVal.Interface()
	//将interface通过断言重新回来
	num2 := iV.(int)
	fmt.Printf("num2=%v num2 type= %T\n", num2, num2)
}

func reflectTest02(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rtype:", rType)

	rVal := reflect.ValueOf(b)

	iV := rVal.Interface()
	fmt.Printf("iV = %v iV type :=%T\n", iV, iV)
	//但是这样的iv无法直接取出name，age之类的
	//反射的本质是运行的，编译的阶段无法确认stu是Student
	//实现了运行时的反射。

	//实现断言，怎么确定他是student？
	stu, ok := iV.(Student)
	if ok {
		fmt.Println("student's name : ", stu.Name)
	}
	//或者用switch iv.(type){
	//case bool:
	//	}
	switch iV.(type) {
	case Student:
		fmt.Println("i am student")
		fmt.Println("student's name : ", iV.(Student).Name)
	case Monster:
		fmt.Println("i am monster")
		fmt.Println("monster's name : ", iV.(Monster).Name)
	}
}
func main() {
	//要求如下
	//先定义一个int
	var num int = 100
	//目标是通过num拿到一系列信息
	reflectTest01(num)

	//对结构体进行反射，
	stu := Student{
		Name: "tom",
		Age:  19,
	}
	reflectTest02(stu)

}
