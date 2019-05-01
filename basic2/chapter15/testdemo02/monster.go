package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (this *Monster) Store() bool {
	data, err := json.Marshal(this)
	if err != nil {
		return false
	}

	filePath := "./monster.txt"
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		fmt.Println(err)
	}

	return true

}

func (this *Monster) ReStore() bool {
	filePath := "./monster.txt"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return false
	}

	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true

}

func main() {
	m01 := &Monster{"tom", 18, "eat"}
	m01.Store()
	m01.ReStore()
	fmt.Println(m01)

}
