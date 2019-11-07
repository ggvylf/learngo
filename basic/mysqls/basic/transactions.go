package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	db, _ := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mydb")
	Db = db
}

func transactions() {
	tx, _ := Db.Begin()

	ret1, e1 := tx.Exec("insert into person(name,age,rmb) values(?,?,?)", "alex", 13, 111)
	ret2, e2 := tx.Exec("update person set rmb =333 where name like ?", "alex")
	ret3, e3 := tx.Exec("delete from person where name like ?", "alex")

	if e1 != nil || e2 != nil || e3 != nil {

		fmt.Println("事物执行失败！", e1, e2, e3)

		tx.Rollback()
	} else {
		tx.Commit()
		ra1, _ := ret1.RowsAffected()
		ra2, _ := ret2.RowsAffected()
		ra3, _ := ret3.RowsAffected()

		fmt.Println("事物执行成功，受影响的行=", ra1+ra2+ra3)
	}

}

func main() {
	defer Db.Close()
	transactions()
}
