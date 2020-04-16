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

	user := new(User)

	// // 根据主键查询第一条记录
	// db.First(&user)
	// //// SELECT * FROM users ORDER BY id LIMIT 1;

	// // 随机获取一条记录
	// db.Take(&user)
	// //// SELECT * FROM users LIMIT 1;

	// // 根据主键查询最后一条记录
	// db.Last(&user)
	// //// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	// // 查询所有的记录
	// db.Find(&users)
	// //// SELECT * FROM users;

	// // 查询指定的某条记录(仅当主键为整型时可用)
	db.First(&user, 2)
	// //// SELECT * FROM users WHERE id = 10;

	


	fmt.Println("user=", user)

}
