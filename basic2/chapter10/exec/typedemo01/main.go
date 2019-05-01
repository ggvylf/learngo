package main

import "fmt"

func TypeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("the %d item type bool,value is %v\n", index+1, x)
		case string:
			fmt.Printf("the %d item type string,value is %v\n", index+1, x)
		case int, int32, int64:
			fmt.Printf("the %d item type int,value is %v\n", index+1, x)
		case float32:
			fmt.Printf("the %d item type float32,value is %v\n", index+1, x)
		case float64:
			fmt.Printf("the %d item type float64,value is %v\n", index+1, x)
		case Student:
			fmt.Printf("the %d item type Studnet,value is %v\n", index+1, x)
		case *Student:
			fmt.Printf("the %d item type *Student,value is %v\n", index+1, x)
		default:
			fmt.Println("not fount")
		}
	}
}

type Student struct{}

func main() {
	s1 := Student{}
	s2 := &Student{}
	TypeJudge(1, 1.3, "hehe", false, s1, s2)

}
