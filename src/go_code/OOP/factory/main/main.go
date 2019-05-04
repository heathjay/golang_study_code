package main

import (
	"fmt"
	"go_code/OOP/factory/model"
)

func main() {
	// var stu = model.Student{
	// 	Name: "tom",
	// 	Age:  19,
	// }
	// fmt.Println(stu)

	var stu = model.NewStudent("tom", 88)
	fmt.Println(*stu)
	fmt.Println("name=", stu.Name, "age=", stu.GetAge())

	//如果model包内，student 中name变成小写。
	// func (s *student) GetAge() int {
	// 	return s.age
	// }

}
