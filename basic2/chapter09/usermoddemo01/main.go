package main

import "fmt"

func modifyUser(users map[string]map[string]string, name string) {
	_, ok := users[name]

	if !ok {
		users[name] = make(map[string]string)
		users[name]["nickname"] = "nick" + name
		users[name]["pwd"] = "123456"
	} else {
		users[name]["pwd"] = "888888"
	}

}

func main() {
	users := make(map[string]map[string]string)
	users["tom"] = make(map[string]string)
	users["tom"] = map[string]string{
		"nickname": "tommy",
		"pwd":      "123456",
	}

	users["jim"] = make(map[string]string)
	users["jim"] = map[string]string{
		"nickname": "jimmy",
		"pwd":      "123456",
	}

	fmt.Println(users)

	modifyUser(users, "jim")

	fmt.Println(users)

}
