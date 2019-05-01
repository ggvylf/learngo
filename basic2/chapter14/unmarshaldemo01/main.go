package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func UnMarshalStruct() {
	str := "{\"Name\":\"tom\",\"Age\":18}"

	var stu Student

	err := json.Unmarshal([]byte(str), &stu)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stu)

}

func UnMarshalMap() {
	//反序列化map不需要make，json.UnMarshal()通过类型断言已经在内部封装了make()
	str := "{\"Name\":\"tom\",\"Name\":\"world\"}"

	var map1 map[string]interface{}

	err := json.Unmarshal([]byte(str), &map1)

	if err != nil {
		fmt.Println(err)
	}
	for k, v := range map1 {
		fmt.Println(k, v)
	}

}

func UnMarshalSlice() {
	str := "[{\"hehe\":\"haha\"},{\"hello\":\"world\"}]"

	var s1 []map[string]interface{}

	err := json.Unmarshal([]byte(str), &s1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s1)
}

func main() {
	UnMarshalStruct()
	UnMarshalMap()
	UnMarshalSlice()

}
