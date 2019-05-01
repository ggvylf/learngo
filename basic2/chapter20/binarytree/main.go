package main

import "fmt"

type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

//前序遍历

func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

func MidOrder(node *Hero) {
	if node != nil {
		PreOrder(node.Left)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		PreOrder(node.Right)
	}
}

func AftOrder(node *Hero) {
	if node != nil {
		PreOrder(node.Left)
		PreOrder(node.Right)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
	}
}

func main() {
	root := &Hero{
		No:   1,
		Name: "tom",
	}

	left1 := &Hero{
		No:   2,
		Name: "jerry",
	}

	right1 := &Hero{
		No:   3,
		Name: "ann",
	}

	root.Left = left1
	root.Right = right1

	right2 := &Hero{
		No:   4,
		Name: "jim",
	}

	right1.Right = right2

	PreOrder(root)
	fmt.Println()
	MidOrder(root)
	fmt.Println()
	AftOrder(root)

}
