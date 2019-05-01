package main

import "fmt"

func days() {
	for {
		var year, month int
		fmt.Print("year=")
		fmt.Scanln(&year)
		fmt.Print("month=")
		fmt.Scan(&month)
		if month >= 12 || month <= 1 {
			fmt.Println("month not match")
			continue
		}
		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			switch month {
			case 2:
				fmt.Println(29)
			case 1, 3, 5, 7, 8, 10, 12:
				fmt.Println(31)
			default:
				fmt.Println(30)
			}
		} else {
			switch month {
			case 2:
				fmt.Println(28)
			case 1, 3, 5, 7, 8, 10, 12:
				fmt.Println(31)
			default:
				fmt.Println(30)
			}
		}
	}

}

func main() {
	days()
}
