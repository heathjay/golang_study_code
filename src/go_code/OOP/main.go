package main

import (
	"encoding/json"
	"fmt"
)

//声明一个结构体
//type 结构体名称 struct{
//字段名称 字段的类型
//}
//结构体首字母是大写，可以被打包引用
//小写私有，只能在本包使用
// 1. 字段名 字段类型
// 2. 如果没有赋值，默认，**指针，slice和map的零值为nil，即没有分配空间**，
// 3. 不同结构体变量的字段是独立的，互不影响，一个结构体变量的改变，
type Cat struct {
	Name  string
	Age   int
	Color string
	Score [3]float64        //类型是数组
	ptr   *int              //指针
	slice []int             //切片
	map1  map[string]string //map

}

type Monster struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, RightDown Point
}

type A struct {
	number int
}
type B struct {
	number int
}

type A1 A

func main() {
	//使用struct来完成养猫案例，
	//年龄属性--字段
	//创建一个Cat的变量
	// var Cat1 Cat //var a int ---int 也可以绑定方法
	// Cat1.Name = "tom"
	// Cat1.Age = 3
	// Cat1.Color = "red"
	// fmt.Println(Cat1)
	// // fmt.Printf("cat1的地址：%p", &Cat1)

	// //定义一个结构体变量
	// var Cat2 Cat
	// fmt.Println(Cat2)
	// //使用slice make
	// //仍然需要make
	// Cat2.slice = make([]int, 10)
	// Cat2.slice[0] = 100
	// Cat2.map1 = make(map[string]string)
	// Cat2.map1["name"] = "tom"
	// fmt.Println(Cat2)

	// //定义两个妖怪
	// //值拷贝
	// var monster1 Monster
	// monster1.Name = "牛魔王"
	// monster1.Age = 500
	// monster2 := &monster1
	// (*monster2).Name = "青牛精"
	// fmt.Println(monster1)
	// fmt.Println(*monster2)
	// //slice引用类型
	// var slice []int = make([]int, 1)
	// slice[0] = 100
	// slice2 := slice
	// slice2[0] = 12
	// fmt.Println(slice2)
	// fmt.Println(slice)

	//方式1.直接声明
	// //方式2.={}  --var person Person = Person{}
	// var monster1 = Monster{"tom", 20}
	// fmt.Println("monster1", monster1)
	// //方式3.返回指针&
	// var monster2 *Monster = new(Monster)
	// //因为是指针，因此标准的赋值方式是
	// (*monster2).Name = "marry"
	// (*monster2).Age = 30
	// fmt.Println(*monster2)
	// //太啰嗦，简化，等价于：
	// monster2.Name = "jone"
	// monster2.Age = 100
	// //在底层会对monster2.Age进行处理，会给monster2加上取值运算，
	// fmt.Println(*monster2)
	// //方式4--{}
	// //案例：
	// var monster3 *Monster = &Monster{"foot", 99}
	// //因为他已经是一个指针，所以标准访问应该取值符.
	// //(*monster3).Name = "jjj"
	// //monster3.Age = 77
	// fmt.Println(*monster3)

	// //说明:1.第三种和第四种方式返回是结构体指针
	// //结构体指针访问字段的标准方式应该是:(*结构体指针).字段名字,比如(*person).Name = "tom"
	// //但go做了一个简化，也支持 结构体指针.字段名，比如person.Name = "tom"

	// //结构体的内存分布
	// //1.支持值拷贝，copy改变不影响后面的，如果想拷贝就是地址拷贝
	// //2.变量总是存在内存中，那么结构体变量在内存究竟是怎么样的存在
	// // var p1 Monster
	// // p1.Age = 10
	// // p1.Name = "good"
	// // var p2 *Monster = &p1
	// // //fmt.Println(*p2.Name) ---报错， . 的运算优先级较高
	// // fmt.Println((*p2).Name)

	// //结构体使用细节和注意事项
	// //1.结构体的所有字段在内存中都是连续分布的，通过地址的加减来快速获取某一个字段的值，速度非常快
	// //r1.leftUp.x 地址0xc00008c020, r1.leftUp.y 地址0xc00008c028 r1.RightDown.x 地址0xc00008c030 r1.RightDown.y 地址0xc00008c038
	// //64位 / 8bit = 8byte --int 的空间大小
	// r1 := Rect{Point{1, 2}, Point{3, 4}}
	// fmt.Printf("r1.leftUp.x 地址%p, r1.leftUp.y 地址%p r1.RightDown.x 地址%p r1.RightDown.y 地址%p \n",
	// 	&r1.leftUp.x, &r1.leftUp.y, &r1.RightDown.x, &r1.RightDown.y)
	// //如果rect存的不是具体的值而是指针就不能这样了
	// //P1指向的地址和P2指向的地址不一样
	// var P1 *Point = &Point{1, 2}
	// var P2 *Point = &Point{3, 4}
	// r2 := Rect{*P1, *P2}
	// //本身连续但是p1,p2指向的地址可能不连续
	// //r2.leftUp 地址0xc00008c040 r2.RightDown地址0xc00008c050
	// fmt.Printf("r2.leftUp 地址%p r2.RightDown地址%p \n", &r2.leftUp, &r2.RightDown)

	// //2.结构体是用户单独定义的类型，和其他类型进行转换时，需要又完全相同的字段（名字/个数/类型）
	// var a A
	// var b B
	// a = A(b)
	// fmt.Println(a, b)

	//3.当结构体使用type进行重新定义的时候（相当于取别名）,golang认为是新的数据类型，但是相互间可以强制转换
	//4.struct 的每个字段上，可以写一个tag，该tag可以通过反射机制获得，常见的使用场景是序列化和反序列化【举例说明】
	//将struct变量进行json处理---做成字符串进行传递--serialization---可以改成name（如果不实用首字母的话）
	monster := Monster{"牛魔王", 500}
	//服务器给客户端发送，需要序列化，结构体如何序列化
	//encoding/json包---marshal==会
	JsonMonster, _ := json.Marshal(monster)
	fmt.Println("jsonStr", string(JsonMonster))
	//jsonStr {"Name":"牛魔王","Age":500}
	//很多程序员不是很喜欢首字母大写，所以用tag:`json:"name"`
	//jsonStr {"name":"牛魔王","age":500}
	//使用到了反射机制，后面仍然会介绍

}
