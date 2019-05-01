package main

import (
	"fmt"

	"github.com/ggvylf/learngo/chapter10/fengzhuangdemo01/model"
)

func main() {
	a := model.NewAccount("hellow", 30, "123456")
	if a != nil {
		fmt.Println("账号创建成功")
	} else {
		fmt.Println("账号创建失败")
		return
	}

	fmt.Println("账号信息=", a)

	a.SetName("abcdef")
	a.SetMoney(500)
	a.SetPasswd("666666")
	if (a.GetName() == "") || (a.GetMoney() == 0) || (a.GetPasswd() == "") {
		fmt.Println("修改信息失败")
		return
	} else {
		fmt.Println("账号信息修改后=", a)

	}

}
