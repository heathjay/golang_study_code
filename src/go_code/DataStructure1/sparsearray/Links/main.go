/*
	单链表的基本方式，
	案例说明：
		第一种方法在添加英雄的时候，
*/

package main

import "fmt"

type HeroNode struct {
	name     string
	next     *HeroNode
	no       int
	nickName string
}

var head *HeroNode

//增加
func init() {
	//初始化
	//1.头节点
	head = &HeroNode{}
}

//3.给链表插入节点
func InsertHeroNode(front *HeroNode, newNode *HeroNode) {
	//先找到链表的最后一位
	//创建一个辅助节点
	temp := front
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newNode
}

//4..编写2种插入方法，根据no 的编号从小到大插入
func InsertHeroNode2(front *HeroNode, newNode *HeroNode) {
	//先找到适当的节点
	//创建一个辅助节点
	temp := front
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newNode.no {
			break
		} else if temp.next.no == newNode.no {
			//说明我们的这链表已经有这个no,就不然插入
			flag = false
			break
		}

		temp = temp.next
	}

	if !flag {
		fmt.Println("对不其，已经存在no=", newNode.no)
		return
	} else {
		newNode.next = temp.next
		temp.next = newNode
	}

}

//显示链表的所有节点信息
func ListHeroNode(front *HeroNode) {
	temp := head
	//1。判断该链表是不是空链表
	if temp.next == nil {
		fmt.Println("empty")
		return
	}
	//遍历这个链表
	for {
		fmt.Printf("节点信息如下：[%d, %s, %s ]==>", temp.next.no, temp.next.name, temp.next.nickName)
		//判断是否是链表最后
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
	fmt.Println()
}

//查找
func Find(front *HeroNode, id int) {
	temp := front
	for {
		if temp.next == nil {
			fmt.Println("there is no such node")
			return
		} else if temp.no == id {
			fmt.Printf("node %d : name:%s\n", temp.no, temp.name)

			return
		}
		temp = temp.next
	}
}

//改动
func Change(front *HeroNode, newNode *HeroNode) {
	temp := front
	for {
		if temp.next == nil {
			fmt.Println("there is no such node")
			return
		} else if temp.next.no == newNode.no {
			newNode.next = temp.next.next
			temp.next = newNode
			fmt.Printf("node %d : name:%s\n", newNode.no, newNode.name)
			return
		}
		temp = temp.next
	}
}

//删除
func DelHeroNode(front *HeroNode, id int) {
	temp := front
	for {
		if temp.next == nil {
			fmt.Println("there is no such node")
			break
		} else if temp.next.no == id {
			temp.next = temp.next.next
			return
		}
		temp = temp.next
	}
}
func main() {
	//2.构建一个新的HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "松江",
		nickName: "rain",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "起灵",
		nickName: "rainx",
	}
	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickName: "rainx",
	}
	//产生随机
	//但是要求必须从小到大
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero2)
	InsertHeroNode2(head, hero3)
	ListHeroNode(head)
	Find(head, 1)
	Find(head, 5)
	hero4 := &HeroNode{
		no:       2,
		name:     "bad",
		nickName: "jj",
	}
	Change(head, hero4)
	ListHeroNode(head)
	DelHeroNode(head, 1)
	ListHeroNode(head)
}
