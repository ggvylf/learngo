package main

type Book struct {
	ID    int64  `db:"db"`
	Title string `db:"title"`
	Price int64  `db:"price"`
}
