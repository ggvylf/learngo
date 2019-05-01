package main

import "fmt"

type Stu struct {
	Name    string
	Age     int
	Address string
}

func main() {

	students := make(map[string]Stu)
	stu01 := Stu{
		"tom",
		18,
		"aaa",
	}
	stu02 := Stu{
		"jerry",
		19,
		"bbb",
	}

	students["001"] = stu01
	students["002"] = stu02

	fmt.Println(students)

	for k, v := range students {
		fmt.Printf("stu number is %v\n", k)
		fmt.Printf("stu name is %v\n", v.Name)
		fmt.Printf("stu age is %v\n", v.Age)
		fmt.Printf("stu add is %v\n", v.Address)
		fmt.Println()

	}

}
