package main

import "fmt"

/*环形单向链表*/
/*一只猫也能实现环形*/

type CatNode struct {
	name string
	no   int
	next *CatNode
}

func InserCatNode(head *CatNode, newCatNode *CatNode) {
	//判断是不是添加第一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head
		fmt.Println(newCatNode, "已经加入到环形的列表中")
		return
	}
	//先定义一个临时变量
	//找到环形的最后
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	//加入到
	temp.next = newCatNode
	newCatNode.next = head

}

//输出这个环形的链表
func ListCirleLink(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("empty")
		return
	}

	for {
		fmt.Println("猫的信息:", temp.name, "->")
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

//删除一只猫
func DelCatNode(head *CatNode, id int) *CatNode {
	//头节点也有值
	//1.先让temp指向head
	//2.定义一个helper指向环形链表的最后
	//让head不动,总栈的head【main】 head[delNode],里面的head变空，外面的head也没有
	//3.让temp和要删除的id进行比较，如果相同，则同helper完成删除，这里必须考虑如果删除就是头节点怎么办
	temp := head
	helper := head
	//判断是否为空
	if temp.next == nil {
		fmt.Println("这是一个空的环形链表不能删除")
		return head
	}

	//假设就只有一个节点
	if temp.next == head {
		//只有一个节点
		temp.next = nil
		return head
	}
	flag := true
	//将helper 定位在链表最后
	for {
		if helper.next == head {
			break
		}

		helper = helper.next
	}
	//假设有两个以上的节点
	for {

		if temp.next == head {
			//如果到这来，说明我比较到最后一个【最后一个还没比较】
			break
		}
		if temp.no == id {
			if temp == head {
				//说明删除的是头节点
				head = head.next
			}
			//恭喜找到
			//我们也可以在这里直接删除
			helper.next = temp.next
			fmt.Printf("猫猫被删掉=%d\n", id)
			flag = false
			break
		}
		temp = temp.next     //temp用来比较
		helper = helper.next //移动的价值，一旦找到删除的节点helper,一旦有这个就删除
	}
	//这里还有比较一次
	if flag {
		//如果flag为真，则我们上面没有删除
		if temp.no == id {
			helper.next = temp.next
			fmt.Println("毛毛=%d\n", id)
		} else {
			fmt.Printf("没有这个猫:%d\n", id)
		}
	}
	return head

}

func main() {
	//这里我们初始化一个环形链表的头节点
	head := &CatNode{}

	//创建一只新猫
	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	cat2 := &CatNode{
		no:   2,
		name: "tom2",
	}
	// cat3 := &CatNode{
	// 	no:   1,
	// 	name: "tom3",
	// }
	InserCatNode(head, cat1)
	InserCatNode(head, cat2)
	// InserCatNode(head, cat3)
	head = DelCatNode(head, 1)
	ListCirleLink(head)
}
