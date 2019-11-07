package main

import (
	"fmt"
	"os"
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

func init() {
	db, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mydb")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	Db = db

}

func main() {
	defer Db.Close()
	// insert()
	// update()
	//delete()
	selects()
}
