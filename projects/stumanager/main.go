package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("欢迎来到学员管理系统")
	fmt.Println("1.添加学员")
	fmt.Println("2.编辑学员")
	fmt.Println("3.显示所有学员")
	fmt.Println("4.退出")

}

func getstu() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("输入学员信息")
	fmt.Scanf("%d\n", &id)
	fmt.Scanf("%s\n", &name)
	fmt.Scanf("%s\n", &class)
	stu := newStudent(id, name, class)
	return stu
}

func main() {
	sm := newStudentMgr(20)
	for {
		showMenu()
		var input int
		fmt.Print("请输入序号:")
		fmt.Scanf("%d\n", &input)
		switch input {
		case 1:
			stu := getstu()
			sm.addStudent(stu)
		case 2:
			stu := getstu()
			sm.editStudent(stu)
		case 3:
			sm.showallStudent()
		case 4:
			os.Exit(0)

		}
	}
}
