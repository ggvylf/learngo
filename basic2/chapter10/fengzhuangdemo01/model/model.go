package model

import "fmt"

type account struct {
	name   string
	money  float64
	passwd string
}

//工厂模式
func NewAccount(name string, money float64, passwd string) *account {
	if len(name) < 6 || len(name) > 10 {
		return nil
	}

	if money <= 20 {
		return nil
	}

	if len(passwd) != 6 {
		return nil
	}

	return &account{
		name:   name,
		money:  money,
		passwd: passwd,
	}
}

func (a *account) SetName(name string) {
	if len(name) >= 6 && len(name) <= 10 {
		a.name = name
	} else {
		fmt.Println("账号长度在6~10")
		a.name = ""
	}
}

func (a *account) GetName() string {
	return a.name
}

func (a *account) SetMoney(money float64) {
	if money > 20 {
		a.money = money
	} else {
		fmt.Println("余额不够20")
		a.money = 0
	}

}

func (a *account) GetMoney() float64 {
	return a.money
}

func (a *account) SetPasswd(passwd string) {
	if len(passwd) != 6 {
		fmt.Println("密码长度必须为6")
		a.passwd = ""
	} else {
		a.passwd = passwd
	}

}

func (a *account) GetPasswd() string {
	return a.passwd
}
