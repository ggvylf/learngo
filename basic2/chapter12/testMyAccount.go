package main

import (
	"fmt"
)

func main() {
	key := ""
	loop := true

	balance := 0.0
	money := 0.0
	note := ""
	details := "类型\t余额\t收支金额\t说明"
	detailsEmpty := false

	for {
		fmt.Println("-----家庭收支------")
		fmt.Println("1. 收支明细")
		fmt.Println("2. 登记收入")
		fmt.Println("3. 登记支出")
		fmt.Println("4. 退出")
		fmt.Print("请选择 1-4 :")

		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("-----收支明细-----")
			if !detailsEmpty {
				fmt.Println("没有收支")
				break
			}
			fmt.Println(details)
			fmt.Println()

		case "2":
			fmt.Println("----登记收入-----")
			fmt.Println("本次收入金额=")
			fmt.Scanln(&money)
			fmt.Println("本次收入说明")
			fmt.Scanln(&note)

			balance += money
			details += fmt.Sprintf("\n收入\t%v\t%v\t\t%v", balance, money, note)
			detailsEmpty = true
		case "3":
			fmt.Println("----登记支出-----")
			fmt.Println("本次支出金额=")
			fmt.Scanln(&money)
			fmt.Println("本次支出说明")
			fmt.Scanln(&note)

			if money > balance {
				fmt.Printf("当前余额不足，当前余额为%v\n", balance)
				break
			} else {

			}

			balance -= money
			details += fmt.Sprintf("\n支出\t%v\t%v\t\t%v", balance, money, note)
		case "4":
			isexit := ""
			fmt.Println("确定退出么(y/n)：")
			fmt.Scanln(&isexit)
			if isexit == "y" {
				loop = false
			} else {
				break
			}

		default:
			fmt.Println("请重新输入")

		}
		if !loop {
			break
		}

	}
	fmt.Println("已经退出")

}
