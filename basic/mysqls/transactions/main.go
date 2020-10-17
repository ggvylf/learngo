package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type Person struct {
	Name     string    `db:"name"`
	Age      int       `db:"age"`
	Money    float64   `db:"rmb"`
	Gender   bool      `db："gender"`
	Birthday time.Time `db："birthday"`
}

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/mydb"
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Println("检查dsn信息，err=", err)
		return
	}
	defer db.Close()

	err = db.Ping()

	if err != nil {
		fmt.Println("mysql ping失败，err=", err)
		return
	}

	return

}

func transactions() {
	//开启事务
	tx, _ := db.Begin()

	//执行sql语句
	ret1, e1 := tx.Exec("insert into person(name,age,rmb) values(?,?,?)", "alex", 13, 111)
	ret2, e2 := tx.Exec("update person set rmb =333 where name like ?", "alex")
	ret3, e3 := tx.Exec("delete from person where name like ?", "alex")

	//如果任意步骤执行失败，则回滚
	if e1 != nil || e2 != nil || e3 != nil {

		fmt.Println("事务执行失败！", e1, e2, e3)

		//回滚事务
		tx.Rollback()

	} else {

		//提交事务
		tx.Commit()
		ra1, _ := ret1.RowsAffected()
		ra2, _ := ret2.RowsAffected()
		ra3, _ := ret3.RowsAffected()

		fmt.Println("事务执行成功，受影响的行=", ra1+ra2+ra3)
	}

}

func selects() {
	var p []Person
	err := db.Select(&p, "select * from person")
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("mysql链接失败，err=", err)
		return
	}

	transactions()
	selects()
}
