package main

import "gorm.io/gorm"

// Many to Many 会在两个 model 中添加一张连接表。
// 会生成一个临时表存放

// User 拥有并属于多种 language，`user_languages` 是连接表
type User struct {
	gorm.Model
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}
