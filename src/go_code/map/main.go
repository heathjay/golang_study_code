package main

import (
	"fmt"
)

func modify(map1 map[int]int) {
	map1[10] = 900
}

type Stu struct {
	Name    string
	Age     int
	Address string
}

func modifyUser(users map[string]map[string]string, name string, nickname string) {

	//判断user中是否有name：key
	_, findRes := users[name]
	if findRes == true {
		fmt.Printf("存在")
		users[name]["pwd"] = "8888"
	} else if findRes == false {
		users[name] = make(map[string]string, 2)
		users[name]["pwd"] = "8888"
		users[name]["nickname"] = nickname
	}

	// if users[name] != nil{
	// 	//有这个用户
	// 	fmt.Printf("存在")
	// 	users[name]["pwd"] = "888888"
	// }else{
	// 	users[name] = make(map[string]string,2)
	// 	users[name]["pwd"] = "8888"
	// 	user[name]["nickname"] = nickname
	// }
}
func main() {
	// var a map[string]string
	// //1.使用map前，先make，先给map分配数据空间
	// a = make(map[string]string, 10)
	// a["non1"] = "松江" //ok
	// a["non2"] = "黄河" //ok
	// a["non1"] = "武松"
	// //2.覆盖前面的non1所对应的内容，
	// fmt.Println(a)

	// //3.map的使用方式
	// //3.1先声明。在make
	// //3.2声明的时候直接make
	// cities := make(map[string]string)
	// cities["no1"] = "天津"
	// cities["no2"] = "北京"
	// cities["no3"] = "上海"
	// fmt.Println(cities)
	// //3.3声明的时候直接赋值

	//map使用的方式
	//存放三个学生信息，每个学生又name和sex信息，map[string]map[string]string
	// stu := make(map[string]map[string]string)
	// stu["stu01"] = make(map[string]string, 3)
	// stu["stu01"]["name"] = "tom"
	// stu["stu01"]["sex"] = "男"
	// stu["stu01"]["address"] = "北京长安街"

	// stu["stu02"] = make(map[string]string, 2) //这句话不能少
	// stu["stu02"]["name"] = "marry"
	// stu["stu02"]["sex"] = "女"
	// stu["stu02"]["address"] = "长安街"
	// fmt.Println(stu)
	// fmt.Println(stu["stu02"])

	// //map的增删改查
	// //map的增加
	// //1. map["key"] = value //如果key没有，就是增加；存在就覆盖
	// //2.删除 delete(map,"key")//如果不存在key也不会报错
	// delete(heros, "heor1")
	// delete(heros, "heros3")
	// fmt.Println(heros)
	// //3.删除map一次性清空，遍历key，逐个删除。或者重新make(...)让原来的成为垃圾，被GC回收。
	// //heros = make(map[string]string)
	// //fmt.Println(heros)
	// //4.查找
	// //val,findRes := heros["no1"]--存在findRes = true
	// val, findRes := heros["hero2"]
	// if findRes == true {
	// 	fmt.Println("有，", val)
	// } else {
	// 	fmt.Println("non")
	// }

	// heros := map[string]string{
	// 	"heor1": "松江",
	// 	"hero2": "将",
	// }
	// heros["hero4"] = "林冲"
	// fmt.Println(heros)
	// fmt.Println("hero有", len(heros), "对key-value")
	//map 的遍历
	//map遍历用for range来遍历
	// for k, val := range heros {
	// 	fmt.Printf("k=%v val:=%v\n", k, val)
	// }
	// //map遍历结构较为复杂的map【string】map[string]string
	// stu := make(map[string]map[string]string)
	// stu["stu01"] = make(map[string]string, 3)
	// stu["stu01"]["name"] = "tom"
	// stu["stu01"]["sex"] = "男"
	// stu["stu01"]["address"] = "北京长安街"

	// stu["stu02"] = make(map[string]string, 2) //这句话不能少
	// stu["stu02"]["name"] = "marry"
	// stu["stu02"]["sex"] = "女"
	// stu["stu02"]["address"] = "长安街"
	// //fmt.Println(stu)
	// fmt.Println(stu["stu02"])
	// for k1, v1 := range stu {
	// 	fmt.Println("k1:=", k1)
	// 	for k2, v2 := range v1 {
	// 		fmt.Printf("\t k2:= %v, v2:%v\n", k2, v2)
	// 	}
	// 	fmt.Println()
	// }

	// map切片
	// 切片的数据类型是map
	// 用map来记录monster的信息,一个monster对应一个map，并且妖怪的个数是可以动态增加的
	// 切片本身也需要make
	// var slice = make([]map[string]string, 2)
	// if slice[0] == nil {
	// 	slice[0] = make(map[string]string, 2)
	// 	slice[0]["name"] = "牛魔王"
	// 	slice[0]["age"] = "500"
	// }
	// if slice[1] == nil {
	// 	slice[1] = make(map[string]string, 2)
	// 	slice[1]["name"] = "玉兔经"
	// 	slice[1]["age"] = "400"
	// }

	// //动态增加map
	// newMonster := map[string]string{
	// 	"name": "网哈哈",
	// 	"age":  "66",
	// }
	// slice = append(slice, newMonster)
	// fmt.Println(slice)

	// //map的key排序
	// //默认无序，每次遍历可能都不一样
	// //1.把key放进切片
	// //2.切片排序，
	// //3.遍历切片，然后按照key来输出map的值
	// map1 := make(map[int]int, 10)
	// map1[10] = 100
	// map1[1] = 13
	// map1[4] = 56
	// map1[8] = 90
	// var keys []int
	// for k, _ := range map1 {
	// 	keys = append(keys, k)
	// }
	// sort.Ints(keys)
	// fmt.Println(keys)
	// for _, v := range keys {
	// 	fmt.Printf("map[%v]=%v\n", v, map1[v])
	// }

	//使用细节
	//1.map是引用类型的，遵守引用类型传递的机制，在一个函数接收map，修改后，会直接修改原来的map--切片也是一样的
	//2.map能够动态增长

	// map1 := make(map[int]int, 2)
	// map1[1] = 90
	// map1[2] = 88
	// map1[10] = 1
	// map1[20] = 2
	// modify(map1)
	// fmt.Println(map1)

	// //3.map的value为student的结构体更加合适 struct
	// //3.1map的key是id
	// //3.1value包含学生名字，年龄，地址
	// stu := make(map[int]Stu, 10)
	// Stu1 := Stu{"tom", 18, "北京"}
	// Stu2 := Stu{"mary", 20, "上海"}
	// stu[1] = Stu1
	// stu[2] = Stu2
	// fmt.Println(stu)
	// //4.遍历
	// for k, v := range stu {
	// 	fmt.Printf("学生的编号：%v，学生的名字：%v,学生的年龄：%v，住址：%v\n", k, v.Name, v.Age, v.Address)
	// }

	//5.练习：
	//使用map[string]map[string]string的map类型 key表示用户名 如果某个用户名存在就讲起密码修改成88888，如果不存在就增加该信息， 包括昵称nickname 和密码
	//编写modifyUser(users map[string]map[string]string, name string)完成上述
	var users = make(map[string]map[string]string, 10)
	users["tom"] = make(map[string]string, 2)
	users["tom"]["pwd"] = "9999"
	users["tom"]["nickname"] = "ooo"
	modifyUser(users, "tom", "tom-jj")
	modifyUser(users, "marr", "tom-wife")
	fmt.Println(users)
}
