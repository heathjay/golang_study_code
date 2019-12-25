package main

import (
	"fmt"
)

/*
1.先用一个二维数组模拟迷宫
	1.1定下规则：
		如果元素值为1，代表一堵墙
		如果元素为0，代表还没走过的点
		如果元素的值为2，代表一个通路
2.编写一个函数，完成老鼠找路
	myMap	- 地图，保证同一个地图，因此使用引用
	i,j		-起点
3.递归分析，什么情况下找到出路
	如果myMap[6][5]==2找到
	2. 如果该点可以探测且没探测过，可以
	3.怎么探测
		1. 假设能通，但是需要探测

*/

func SetWay(myMap *[8][7]int, i int, j int) bool {
	if myMap[6][5] == 2 {
		return true
	} else {
		if myMap[i][j] == 0 {
			//说明可以探测
			//假设能探测
			myMap[i][j] = 2
			if SetWay(myMap, i+1, j) {
				//向上探测
				//下右上左
				return true
			} else if SetWay(myMap, i, j+1) {
				return true
			} else if SetWay(myMap, i-1, j) {
				return true
			} else if SetWay(myMap, i, j-1) {
				return true
			} else {
				//都没有走通说明是一个死路
				myMap[i][j] = 3
				return false
			}

		} else {
			//说明是一堵墙
			return false
		}
	}
}
func main() {
	var myMap [8][7]int
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}
	myMap[3][1] = 1
	myMap[3][2] = 1
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}
	SetWay(&myMap, 1, 1)
	fmt.Println("finish")
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}
}
