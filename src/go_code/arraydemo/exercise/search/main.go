package main

import (
	"fmt"
)

func main() {
	//定一个数组

	var str [4]string = [4]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "亲衣服网"}
	var heroName = ""
	fmt.Println("input names:")
	fmt.Scanln(&heroName)

	//顺序查找的第一种方式
	// for i := 0; i < len(str); i++ {
	// 	if heroName == str[i] {
	// 		fmt.Printf("found%v,下标%v...\n", heroName, i)
	// 		break
	// 	} else if i == len(heroName)-1 {
	// 		fmt.Println("bad job")
	// 	}
	// }

	//顺序查找第二种方式
	index := -1
	//找到了下标改变
	for i := 0; i < len(str); i++ {
		if heroName == str[i] {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到%v，下标%v\n", str[index], index)
	} else {
		fmt.Println("bad job")
	}

	//
}
