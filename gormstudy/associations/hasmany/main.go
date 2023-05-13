package main

import "gorm.io/gorm"

// 与另一个模型建立了一对多的连接。 不同于 has one，拥有者可以有零或多个关联模型。
// has one要求必须有一个关联关系，has many的关联关系可以为0

// User 有多张 CreditCard，UserID 是外键
type User struct {
	gorm.Model
	CreditCards []CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}
