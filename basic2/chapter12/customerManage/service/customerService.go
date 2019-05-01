package service

import (
	"github.com/ggvylf/learngo/chapter12/customerManage/model"
)

type CustomerService struct {
	customers   []model.Customer
	customerNum int
}

func NewCustomerService() *CustomerService {
	cv := &CustomerService{}
	cv.customerNum = 1
	customer := model.NewCoustomer(1, "tom", "male", 20, "123", "tom@163.com")
	cv.customers = append(cv.customers, customer)
	return cv
}

func (cv *CustomerService) List() []model.Customer {
	return cv.customers
}

func (cv *CustomerService) Add(customer model.Customer) bool {
	cv.customerNum++
	customer.Id = cv.customerNum
	cv.customers = append(cv.customers, customer)
	return true

}

func (cv *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(cv.customers); i++ {
		if cv.customers[i].Id == id {
			index = i
		}
	}
	return index
}

func (cv *CustomerService) Delete(id int) bool {
	index := cv.FindById(id)

	if index == -1 {
		return false
	}
	cv.customers = append(cv.customers[:index], cv.customers[index+1:]...)
	return true
}

func (cv *CustomerService) Edit(id int, name string, gender string, age int, phone string, email string) bool {
	index := cv.FindById(id)

	cv.customers[index].Name = name
	cv.customers[index].Gender = gender
	cv.customers[index].Age = age
	cv.customers[index].Phone = phone
	cv.customers[index].Email = email
	return true
}
