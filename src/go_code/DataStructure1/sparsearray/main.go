package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type ValNode struct {
	Row int `json:row`
	Col int `json:col`
	Val int `json:val`
}

func main() {
	//1.先创建一个原始的数组
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	//2.输出原始的数组
	// for _, v := range chessMap {
	// 	for _, v2 := range v {
	// 		fmt.Printf("%d\t", v2)
	// 	}
	// 	fmt.Println()
	// }
	//3.转换成稀疏数字，用切片，最好的设计方式就是结构体
	//思路如下：
	//1.遍历chessMap,如果我们发现有一个元素的值，不等于0，就创建一个Node 节点
	//2.将其放入对应的切片里面就好
	var sparseArr []ValNode

	//标准的稀疏数组，应该还有一个记录原始的二维数组的规模和默认值
	valNode := ValNode{
		Row: 11,
		Col: 11,
		Val: 0,
	}

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				//创建一个节点
				valNode = ValNode{
					Row: i,
					Col: j,
					Val: v2,
				}

				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	//输出稀疏数组
	// for i, valNode := range sparseArr {
	// 	fmt.Printf("%d: %d %d %d", i, valNode.row, valNode.col, valNode.val)
	// 	fmt.Println()
	// }

	//将这个数组存在文件里面
	filePath := "/Users/chengpengjiang/Documents/coding/golang/exgolang/src/go_code/DataStructure1/sparsearray/chessmap.data"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()

	if err != nil {
		fmt.Println("open file err:", err)
	}
	fmt.Println("往文件里面写存盘")
	writer := bufio.NewWriter(file)
	//encoder := json.NewEncoder(writer)
	//writer := bufio.NewWriter(file)
	//encoder := json.NewEncoder(writer)
	// for i := 0; i < len(sparseArr); i++ {
	// 	err := encoder.Encode(sparseArr[i])
	// 	if err != nil {
	// 		fmt.Println("json input err:", err)
	// 	}

	// 	// str := fmt.Sprint(data)
	// 	// // fmt.Println(str)
	// 	// writer.WriteString(str)
	// }
	fmt.Println(sparseArr)
	data, err := json.Marshal(sparseArr)
	//fmt.Println(data)
	if err != nil {
		fmt.Println("json err")
	}
	writer.Write(data)
	writer.Flush()
	//如何恢复原始的数组
	//1. 打开文件，恢复数组
	file, err = os.Open(filePath)
	defer file.Close()

	// var temp *ValNode
	reader := bufio.NewReader(file)
	// decoder := json.NewDecoder(reader)
	// err = decoder.Decode(&temp)
	// if err != nil {
	// 	fmt.Println("decoder err")
	// 	return
	// }
	// fmt.Println(temp)
	// var chessMap2 [11][11]int
	data1, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("read err:", err)
		return
	}
	var sparseArr2 []ValNode
	json.Unmarshal(data1, &sparseArr2)
	fmt.Println(sparseArr2)
	// for {

	// 	chessMap2[temp.row][temp.col] = temp.val
	// }

	//2.遍历数据来决定数组的行和列

	// //遍历稀疏数组，
	// //遍历文件的每一行
	var chessMap2 [11][11]int
	for _, valNode := range sparseArr2 {
		if valNode.Val == 0 {
			continue
		}
		chessMap2[valNode.Row][valNode.Col] = valNode.Val
	}
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

}
