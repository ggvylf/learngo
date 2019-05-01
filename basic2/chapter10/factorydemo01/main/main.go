package main

import (
	"fmt"

	"github.com/ggvylf/learngo/chapter10/factorydemo01/model"
)

func main() {
	stu := model.NewStudent("tom", 12)
	fmt.Println(*stu)
	fmt.Println("name=", stu.Name, "age=", stu.GetAge())

}
