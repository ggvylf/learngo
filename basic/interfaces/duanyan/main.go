//类型断言
package main

import "fmt"

func Test(a interface{}) {
	//断言，a是个int类型，如果不是int类型就走if流程
	b, ok := a.(int)
	if !ok {
		fmt.Println("not a int")
		return
	}
	b += 3
	fmt.Printf("%d\n", b)
}

func testType(items ...interface{}) {
	for _, v := range items {
		switch v.(type) {
		case int, int32, int64:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("bool")
		default:
			fmt.Println("unknow")
		}
	}
}

func main() {
	var a int
	Test(a)

	testType("123", 123, 46)
}
