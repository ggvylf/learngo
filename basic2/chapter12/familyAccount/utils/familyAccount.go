package utils

import "fmt"

type FamilyAccount struct {
	key          string
	loop         bool
	balance      float64
	money        float64
	note         string
	details      string
	detailsEmpty bool
}

func (f *FamilyAccount) ShowDetails() {
	fmt.Println("-----收支明细-----")
	if !f.detailsEmpty {
		fmt.Println("没有收支")
		return
	}
	fmt.Println(f.details)
	fmt.Println()
}
func (f *FamilyAccount) AddDetails() {
	fmt.Println("----登记收入-----")
	fmt.Println("本次收入金额=")
	fmt.Scanln(&f.money)
	fmt.Println("本次收入说明")
	fmt.Scanln(&f.note)

	f.balance += f.money
	f.details += fmt.Sprintf("\n收入\t%v\t%v\t\t%v", f.balance, f.money, f.note)
	f.detailsEmpty = true

}

func (f *FamilyAccount) DelDetails() {
	fmt.Println("----登记支出-----")
	fmt.Println("本次支出金额=")
	fmt.Scanln(&f.money)
	fmt.Println("本次支出说明")
	fmt.Scanln(&f.note)

	if f.money > f.balance {
		fmt.Printf("当前余额不足，当前余额为%v\n", f.balance)
		return
	}
	f.balance -= f.money
	f.details += fmt.Sprintf("\n支出\t%v\t%v\t\t%v", f.balance, f.money, f.note)

}

func (f *FamilyAccount) Exit() {
	isexit := ""
	fmt.Println("确定退出么(y/n)：")
	fmt.Scanln(&isexit)
	if isexit == "y" {
		f.loop = false
	} else {
		return
	}

}

//工厂模式，返回×FamilyAccount实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:  "",
		loop: true,

		balance:      0.0,
		money:        0.0,
		note:         "",
		details:      "类型\t余额\t收支金额\t说明",
		detailsEmpty: false,
	}
}

func (f *FamilyAccount) MainMenu() {
	for {
		fmt.Println("-----家庭收支------")
		fmt.Println("1. 收支明细")
		fmt.Println("2. 登记收入")
		fmt.Println("3. 登记支出")
		fmt.Println("4. 退出")
		fmt.Print("请选择 1-4 :")

		fmt.Scanln(&f.key)

		switch f.key {
		case "1":
			f.ShowDetails()

		case "2":
			f.AddDetails()

		case "3":
			f.DelDetails()

		case "4":
			f.Exit()

		default:
			fmt.Println("请重新输入")

		}
		if !f.loop {
			break
		}

	}
	fmt.Println("已经退出")
}
