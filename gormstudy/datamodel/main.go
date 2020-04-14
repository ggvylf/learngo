package main

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	//内嵌gorm.Modle
	gorm.Model

	Name         string
	Age          sql.NullInt64
	Birthdya     *time.Time
	Email        string  `gorm: "type:varchar(100);unique_index"` //制定大小病创建唯一索引
	Role         string  `gorm:"size:255`                         //指定长度
	MemberNmuber *string `gorm:"unique;not null"`                 //`gorm:"column:mem_nums;unique;not null"` 指定字段名称
	Num          int     `gorm:"AUTO_INCREMENT"`
	Address      string  `gorm:"index:addr"` //创建名称为addr的索引
	IgnoreMe     int     `gorm:"-"`          //忽略该字段
}

type Animal struct {
	AnimalID int64 `gorm:"primary_key"` //明确指定主键
	Type     string
}

//修改表名
// func (a Animal) TableName() string {
// 	//根据字段来判断
// 	if a.Type=="admin" {
// 		return "admin_myAnimal"
// 	}else {
// 		return "myAnimal"
// 	}

// }

func main() {
	//修改表名的默认命名规则
	// gorm.DefaultTableNameHandler=func( db *gorm.DB,defualtTableName string) string {
	// 	return "prefix_"+defualtTableName
	// }

	//连接数据库
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/mytest?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//禁用复数表名
	// db.SingularTable(true)

	//指定表名
	//使用User结构体创建名称为new_user的表
	// db.Table("new_users").CreateTable(&User{})
	db.AutoMigrate(&User{})

}
