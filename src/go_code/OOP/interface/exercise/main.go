package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//接口的最佳实践
//1.声明hero结构体
type Hero struct {
	Name string
	Age  int
}

//2.声明一个hero结构体切片类型
type HeroSlice []Hero

//3.实现interface 接口 :3个方法来满足
func (hs HeroSlice) Len() int {
	return len(hs)
}

//less这个方法决定你使用什么标准进行排序
//1.按年龄从小到达排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
	//改成对名字排序
	//return hs[i].Name < hs[j].Name
}

func (hs HeroSlice) Swap(i, j int) {
	//交换有点繁琐
	// tmp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = tmp

	//等价于
	hs[i], hs[j] = hs[j], hs[i]
}

//学生结构体
type Student struct {
	Name  string
	Age   int
	Score float64
}

//实现了interface 这个接口 就可以满足了。
func main() {
	//先定一个数组/切片
	//实现对hero 的排序
	var intSlice = []int{-1, 0, 10, 7, 99}
	//要求对intSlice 进行排序
	//1.冒泡排序
	//2.使用系统的方法

	sort.Ints(intSlice) //sort内部排序会直接影响刀切片
	fmt.Println(intSlice)

	//对一个结构体切片进行排序
	//1.冒泡排序....
	//2.系统提供的方法

	//定一个切片
	var heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄～%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}
	fmt.Println("排序前的顺序:")

	for _, v := range heros {
		fmt.Println(v)
	}

	//调用sort.sort函数
	sort.Sort(heros)
	fmt.Println("排序后的顺序:")

	for _, v := range heros {
		fmt.Println(v)
	}
}
