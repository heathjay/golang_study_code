package main

import "fmt"

type Hero struct {
	Name  string
	No    int
	Left  *Hero
	Right *Hero
}

//pre-Order // 先输出根节-左边-右边
func PreOrder(node *Hero) {

	if node != nil {
		fmt.Printf("no = %d\n", node.No)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

//中序遍历，左-中-右
func InfixOrder(node *Hero) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("no = %d\n", node.No)

		InfixOrder(node.Right)
	}
}

func main() {
	root := &Hero{
		No:   1,
		Name: "songjiang",
	}

	left := &Hero{
		No:   2,
		Name: "liqiu",
	}

	right := &Hero{
		No:   3,
		Name: "lu",
	}
	root.Left = left
	root.Right = right

	right2 := &Hero{
		No:   4,
		Name: "fei",
	}

	right.Right = right2
	PreOrder(root)
	InfixOrder(root)
}
