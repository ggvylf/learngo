package model

import "fmt"

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

//工厂模式
func NewCoustomer(id int, name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func NewCoustomer2(name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func (cv *Customer) Getinfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", cv.Id, cv.Name, cv.Age, cv.Gender, cv.Phone, cv.Email)
	return info

}
