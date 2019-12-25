package main

import (
	"fmt"
)

type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

//前序遍历,root,left,right

func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("no=%d,name=%s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

//中序遍历 left,root,right
func InfixOrder(node *Hero) {
	if node != nil {

		InfixOrder(node.Left)
		fmt.Printf("no=%d,name=%s\n", node.No, node.Name)
		InfixOrder(node.Right)
	}
}

//后序 left,right,root
func PostOrder(node *Hero) {
	if node != nil {

		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("no=%d,name=%s\n", node.No, node.Name)

	}
}

func main() {
	root := &Hero{
		No:   1,
		Name: "songjiang",
	}

	left1 := &Hero{
		No:   2,
		Name: "wu yong",
	}
	node10 := &Hero{
		No:   10,
		Name: "10",
	}

	node11 := &Hero{
		No:   12,
		Name: "12",
	}
	left1.Left = node10
	left1.Right = node11
	right1 := &Hero{
		No:   3,
		Name: "lujuyi",
	}

	root.Left = left1
	root.Right = right1
	right2 := &Hero{
		No:   4,
		Name: "nihao",
	}
	right1.Right = right2
	//PreOrder(root)
	//InfixOrder(root)
	PostOrder(root)

}
