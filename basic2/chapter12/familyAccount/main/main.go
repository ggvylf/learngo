package main

import (
	"fmt"

	"github.com/ggvylf/learngo/chapter12/familyAccount/utils"
)

func main() {
	fmt.Println("面向对象")
	utils.NewFamilyAccount().MainMenu()
}
