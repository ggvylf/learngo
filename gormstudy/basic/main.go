package main


import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)


type UserInfo struct {
	ID uint
	Name string
	Gender string
	Hobby string
}


func main() {

	//连接数据库
	db,err:=gorm.Open("mysql","root:123456@(127.0.0.1:3306)/mytest?charset=utf8mb4&parseTime=True&loc=Local")
	if err!=nil {
		panic(err)
	}
	defer db.Close()


	//创建表
	db.AutoMigrate(&UserInfo{})
	u1:=UserInfo{1,"tom","male","drink"}

	//创建数据
	db.Create(&u1)


	//查询
	var u UserInfo

	//查询表中的第一条数据
	db.First(&u)
	fmt.Printf("u=%#v\n",u)

	//更新
	db.Model(&u).Update("hobby","book")
	db.First(&u)
	fmt.Println("after update")
	fmt.Printf("u=%#v\n",u)

	//删除
	db.Delete(&u)




}