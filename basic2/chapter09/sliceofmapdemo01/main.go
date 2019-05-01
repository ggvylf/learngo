package main

import (
	"fmt"
)

func main() {
	var users []map[string]string
	users = make([]map[string]string, 2)

	users[0] = make(map[string]string, 2)
	users[0]["name"] = "tom"
	users[0]["sex"] = "male"

	users[1] = make(map[string]string, 2)
	users[1]["name"] = "jerry"
	users[1]["sex"] = "female"

	newuser := map[string]string{
		"name": "ann",
		"sex":  "male",
	}

	users = append(users, newuser)

	fmt.Println(users)

}
