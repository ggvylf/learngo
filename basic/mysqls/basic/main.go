package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	Name     string    `db:"name"`
	Age      int       `db:"age"`
	Money    float64   `db:"rmb"`
	Gender   bool      `db："gender"`
	Birthday time.Time `db："birthday"`
}

func insert() {
	// 执行语句
	result, err := Db.Exec("insert into person(name,age,rmb,gender,birthday) values(?,?,?,?,?)", "peter", 11, 444, true, 19930101)
	if err != nil {
		fmt.Println(err)
	}

	rosAffected, _ := result.RowsAffected()
	lastInsertId, _ := result.LastInsertId()
	fmt.Println("受影响的行数=", rosAffected)
	fmt.Println("最后一行的id=", lastInsertId)

}

func update() {
	result, err := Db.Exec("update person set age= ? where name like ?;", 44, "%peter%")
	if err != nil {
		fmt.Println(err)
	}

	rosAffected, _ := result.RowsAffected()
	lastInsertId, _ := result.LastInsertId()
	fmt.Println("受影响的行数=", rosAffected)
	fmt.Println("最后一行的id=", lastInsertId)

}

func delete() {
	result, err := Db.Exec("delete from person where name  like ?;", "%peter%")
	if err != nil {
		fmt.Println(err)
	}

	rosAffected, _ := result.RowsAffected()
	lastInsertId, _ := result.LastInsertId()
	fmt.Println("受影响的行数=", rosAffected)
	fmt.Println("最后一行的id=", lastInsertId)
}

func selects() {
	var ps []Person
	err := Db.Select(&ps, "select name,age,rmb from person")
	// err := Db.Select(&ps, "select name,age,rmb from person where name like ?", "%peter%")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("查询成功！", ps)

}

var Db *sqlx.DB

func initDB() (err error) {
	//数据库信息
	//username:password@tcp(ip:port)/dbname
	dsn := "root:123456@tcp(127.0.0.1:3306)/mydb"
	//连接数据库
	//这里不要使用:= 不然创建的就是临时变量，而不是全局变量，会报空指针的问题
	Db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Println("dsn格式错误，err=", err)
		return

	}

	// 尝试连接数据库
	err = Db.Ping()
	if err != nil {
		fmt.Println("数据库ping失败，err=", err)
		return

	}

	//最大连接数
	Db.SetMaxOpenConns(200)
	//最大空闲链接
	Db.SetMaxIdleConns(1000)

	return
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("数据库初始化失败")
	}

	Db.QueryRow()

	defer Db.Close()
	// insert()
	// update()
	//delete()
	selects()
}
