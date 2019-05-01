	package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(192.168.56.11:3306)/gotest")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

//增加
func Add() {
	r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)
}

//查询
func Select() {
	var person []Person
	err := Db.Select(&person, "select user_id, username, sex, email from person")
	//err := Db.Select(&person, "select user_id, username, sex, email from person where user_id=?", 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("select succ:", person)

}

//修改
func Edit() {
	_, err := Db.Exec("update person set username=? where user_id=?", "stu0001", 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

}

//删除
func Drop() {

	// _, err := Db.Exec("delete from person where user_id=?", 1)
	_, err := Db.Exec("delete from person")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("delete succ")

}

func main() {
	Add()
	Select()
	Edit()
	Drop()

}
