package main

import (
	"fmt"
	"time"
)

func panDuan() {
	start := time.Date(1990, time.January, 1, 0, 0, 0, 0, time.Local)

	var year int
	var month int
	var day int
	fmt.Println("year=")
	fmt.Scan(&year)
	fmt.Println("month=")
	fmt.Scan(&month)
	fmt.Println("day=")
	fmt.Scan(&day)

	end := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)

	sum := (end.Unix() - start.Unix()) / 3600 / 24

	switch sum % 5 {
	case 1, 2, 3:
		fmt.Println("打鱼")
	default:
		fmt.Println("晒网")
	}

}

func main() {
	panDuan()

}
