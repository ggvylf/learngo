package main

import (
	"fmt"

	"github.com/ggvylf/learngo/chapter12/customerManage/model"
	"github.com/ggvylf/learngo/chapter12/customerManage/service"
)

type customerView struct {
	key             string
	loop            bool
	customerService *service.CustomerService
}

func (cv *customerView) list() {
	customer := cv.customerService.List()
	fmt.Println("-------客户列表-------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customer); i++ {
		fmt.Println(customer[i].Getinfo())
	}
	fmt.Println("\n-------客户列表完成-------\n")

}

func (cv *customerView) add() {
	fmt.Println("-------添加客户--------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱:")
	email := ""
	fmt.Scanln(&email)

	customer := model.NewCoustomer2(name, gender, age, phone, email)

	if cv.customerService.Add(customer) {
		fmt.Println("---------添加用户成功---------")
	} else {
		fmt.Println("---------添加用户失败---------")
	}

}

func (cv *customerView) delete() {
	fmt.Println("-------删除客户--------")
	fmt.Println("请输入用户的id（-1表示退出）")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Println("退出删除")
		return
	}

	fmt.Println("确认是否删除（y/n）")
	isok := ""
	fmt.Scanln(&isok)
	if isok == "y" {
		if cv.customerService.Delete(id) {
			fmt.Println("已删除")
		} else {
			fmt.Println("输入id不存在")
		}
	}
}

func (cv *customerView) edit() {
	fmt.Println("-------修改客户--------")
	fmt.Println("请输入用户的id（-1表示退出）")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Println("退出修改")
		return
	}

	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱:")
	email := ""
	fmt.Scanln(&email)

	if cv.customerService.Edit(id, name, gender, age, phone, email) {
		fmt.Println("---------修改用户成功---------\n")
	} else {
		fmt.Println("---------修改用户失败---------\n")
	}

}

func (cv *customerView) MainMenu() {
	for {
		fmt.Println("--------客户信息管理----------")
		fmt.Println("1. 添加客户")
		fmt.Println("2. 修改客户")
		fmt.Println("3. 删除客户")
		fmt.Println("4. 客户列表")
		fmt.Println("5. 退   出")
		fmt.Print("请选择（1-5）：")
		fmt.Scanln(&cv.key)

		switch cv.key {
		case "1":
			cv.add()
		case "2":
			cv.edit()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.loop = false
		default:
			fmt.Println("重新选择")
			break
		}
		if !cv.loop {
			break
		}

	}

	fmt.Println("已退出系统")
}

func main() {
	cv := customerView{
		key:  "",
		loop: true,
	}
	cv.customerService = service.NewCustomerService()

	cv.MainMenu()
}
