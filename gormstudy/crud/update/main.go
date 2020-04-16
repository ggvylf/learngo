package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//定义数据模型
type User struct {
	ID   int64
	Name string `gorm:"default:'noname'"`
	Age  int64
}

func main() {

	//链接数据库
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/mytest?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//把模型和数据库中的表对应
	db.AutoMigrate(&User{})

	//创建数据
	user := User{Name: "tom", Age: 18}

	db.Model(&user).Update("name", "jerry")

	db.First(&user)
	fmt.Println(user)

}
