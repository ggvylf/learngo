package main

import "fmt"

type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode
	next     *HeroNode
}

//无序插入
func InsertHeroNode(head *HeroNode, newnode *HeroNode) {
	temp := head

	for {
		//如果next为nil，表示已经到链表的尾部
		if temp.next == nil {

			break
		}
		//不为nil就继续向后移动
		temp = temp.next
	}
	//添加到链表
	temp.next = newnode
	newnode.pre = temp

}

//有序插入
func InsertHeroNode2(head *HeroNode, newnode *HeroNode) {
	temp := head
	flag := true
	for {
		//如果next为nil，表示已经到链表的尾部
		if temp.next == nil {
			break
			//temp的下一个节点的id比newnode的id大，newnode就该插入到temp后边，可以退出循环执行插入操作了
		} else if temp.next.no > newnode.no {
			break
			//表示id冲突
		} else if temp.next.no == newnode.no {
			flag = false
			break
		}
		//移动temp
		temp = temp.next
	}

	if !flag {
		fmt.Println("id冲突，无法插入")
		return
	} else {
		//temp的下一个节点就是newnode的下一个节点
		newnode.next = temp.next
		//newnode的前一个节点是temp
		newnode.pre = temp

		//temp的下一个节点的前一个节点就是newnode

		if temp.next != nil {
			temp.next.pre = newnode
		}

		//temp的下一个节点就是newnode
		temp.next = newnode

	}

}

//删除
func DelHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false

	for {
		//到达队列尾部队，跳出循环
		if temp.next == nil {
			fmt.Println("empty link")
			break
			//temp的下一个节点是要删除的节点，标记并跳出循环
		} else if temp.next.no == id {
			flag = true
			break
		}
		//移动temp到下一个
		temp = temp.next
	}

	if flag {
		//temp下个一节点就是temp的下一个的下一个
		temp.next = temp.next.next
		//temp是最后一个节点
		if temp.next != nil {
			temp.next.pre = temp
		}

	} else {
		fmt.Println("sorry,要删除的id不存在")
	}

}

func ListLink(head *HeroNode) {
	temp := head

	//判断是否是空链表
	if temp.next == nil {
		fmt.Println("empty link")
		return
	}

	//遍历链表
	for {
		fmt.Printf("[%d,%s,%s]-->", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			fmt.Println("tail")
			return
		}

	}
	fmt.Println()

}

func ListLink2(head *HeroNode) {
	temp := head

	//判断是否是空链表
	if temp.next == nil {
		fmt.Println("empty link")
		return
	}

	//temp移动到链表结尾
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next

	}

	//遍历链表
	for {
		fmt.Printf("[%d,%s,%s]-->", temp.no, temp.name, temp.nickname)
		temp = temp.pre

		if temp.pre == nil {
			fmt.Println("head")
			break
		}

	}
	fmt.Println()

}

func main() {

	//创建一个head节点
	head := &HeroNode{}

	//创建HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "tom",
		nickname: "tommy",
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "jim",
		nickname: "jimmy",
	}

	hero3 := &HeroNode{
		no:       3,
		name:     "ann",
		nickname: "ann",
	}

	//加入链表
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero2)

	//删除
	// DelHeroNode(head, 2)

	//显示链表
	ListLink(head)
	ListLink2(head)

}
