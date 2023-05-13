package main

import "gorm.io/gorm"

// 与另一个模型建立了一对一的连接。 这种模型的每一个实例都“属于”另一个模型的一个实例。
// 注意是从属关系

// `User` 属于 `Company`，`CompanyID` 是外键
type User struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company
}

type Company struct {
	ID   int
	Name string
}
