package main

import "fmt"

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {

	//添加第一个节点，并且构成环形
	if head.next == nil {

		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head
		fmt.Println(newCatNode.no, "加入到链表")
		return
	}

	//其他节点
	temp := head

	//找到环装链表的末尾
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}

	//加入链表
	temp.next = newCatNode
	newCatNode.next = head

}

func ListCircelLink(head *CatNode) {
	temp := head

	//判断是否为空联链表
	if temp.next == nil {
		fmt.Println("empty")
		return
	}

	for {
		fmt.Printf("[no=%d name=%s]-->", temp.no, temp.name)
		if temp.next == head {
			fmt.Printf("tail")
			break
		}
		temp = temp.next
	}
	fmt.Println()
}

//如果删除的是head，head会变化
func DelCatNode(head *CatNode, no int) *CatNode {
	temp := head
	helper := head

	//判断是否为空
	if temp.next == nil {
		fmt.Println("empty")
		return head
	}

	//只有一个节点
	if temp.next == head {
		//no就是head的no
		if temp.no == no {
			temp.next = nil
		}
		return head
	}

	//helper定位到链表最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	//2个以上节点
	flag := true
	for {
		//temp到链表尾部,但是没有比较temp跟最后一个元素，放在后边比较
		if temp.next == head {
			break
		}

		if head.no == no {
			head = temp.next
			break
		}

		//找到对应的id
		if temp.no == no {
			//用helper删除temp
			helper.next = temp.next
			fmt.Printf("no=%d cat is del\n", no)
			//已经删除了
			flag = false
			break
		}
		//temp负责移动，helper负责删除temp
		temp = temp.next
		helper = helper.next

	}

	//比较最后一个元素
	if flag {
		if temp.no == no {
			helper.next = temp.next
			fmt.Printf("no=%d cat is del in flag\n", no)
		} else {
			fmt.Println("id不匹配")
		}

	}

	return head

}

func main() {
	head := &CatNode{}

	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}

	cat2 := &CatNode{
		no:   2,
		name: "jerry",
	}

	cat3 := &CatNode{
		no:   3,
		name: "ann",
	}

	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)

	ListCircelLink(head)

	head = DelCatNode(head, 3)

	ListCircelLink(head)
}
