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

//在字段有默认值的情况下仍要传入0值
//方法1 使用指针
// type User struct {
// 	ID int64
// 	Name ×string `gorm:"default:'noname'"`
// 	Age  int64
// }
// u:=User{Name:new(string),Age:18)}
// db.Create(&u)
//
//方法2 使用Scanner/Valuer
// type User struct {
// 	ID int64
// 	Name sql.NullString `gorm:"default:'noname'"`
// 	Age  int64
// }
// u:=User{Name:sql.NullString("",true),Age:18)}
// db.Create(&u)

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
	u1 := User{Name: "tom", Age: 18}

	u2 := User{Age: 33}

	//创建数据
	//判断主键是否为空，true表示为空，false表示已经存在
	fmt.Println("创建前主键的状态=", db.NewRecord(u1))
	db.Create(&u1)
	fmt.Println("创建后主键的状态=", db.NewRecord(u1))
	db.Create(&u2)

}
