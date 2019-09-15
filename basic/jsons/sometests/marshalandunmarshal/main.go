package marshalandunmarshal

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Rmb   float64
	Sex   bool
	Hobby []string
}

func structtojson() {
	person := Person{"tom", 11, 123.45, true, []string{"hehe", "haha"}}
	fmt.Println(person)
	bytes, err := json.Marshal(person)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))

	dataMap1 := make(map[string]interface{})
	dataMap1["name"] = "jerry"
	dataMap1["hobby"] = []string{"aa", "bb", "cc"}

	dataMap2 := make(map[string]interface{})
	dataMap2["name"] = "jim"
	dataMap2["hobby"] = []string{"aaa", "bbb", "ccc"}
	dataMap3 := make(map[string]interface{})
	dataMap3["name"] = "ann"
	dataMap3["hobby"] = []string{"a", "b", "c"}

	dataslice := make([]map[string]interface{}, 0)
	dataslice = append(dataslice, dataMap1, dataMap2, dataMap3)
	bytes, err = json.Marshal(dataslice)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
}

func jsontomap() {
	jsonStr := `{"Name":"tom","Age":11,"Rmb":123.45,"Sex":true,"Hobby":["hehe","haha"]}`
	jsonBytes := []byte(jsonStr)
	dataMap := make(map[string]interface{})
	err := json.Unmarshal(jsonBytes, &dataMap)
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range dataMap {
		fmt.Println(k, v)
	}

}

func jsontostruct() {
	jsonStr := `{"Name":"tom","Age":11,"Rmb":123.45,"Sex":true,"Hobby":["hehe","haha"]}`
	jsonBytes := []byte(jsonStr)
	p := new(Person)
	err := json.Unmarshal(jsonBytes, &p)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(*p)
	fmt.Println(p.Name)
	fmt.Println(p.Age)
}

func jsontomapslice() {
	jsonStr := `[{"hobby":["aa","bb","cc"],"name":"jerry"},{"hobby":["aaa","bbb","ccc"],"name":"jim"},{"hobby":["a","b","c"],"name":"ann"}]
`
	jsonBytes := []byte(jsonStr)
	dataSlice := make([]map[string]interface{}, 0)
	err := json.Unmarshal(jsonBytes, &dataSlice)
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range dataSlice {
		fmt.Println(k, v)
	}
}

func jsontostructslice() {
	jsonStr := `[{"hobby":["aa","bb","cc"],"name":"jerry"},{"hobby":["aaa","bbb","ccc"],"name":"jim"},{"hobby":["a","b","c"],"name":"ann"}]
`
	jsonBytes := []byte(jsonStr)
	dataSlice := make([]Person, 0)
	err := json.Unmarshal(jsonBytes, &dataSlice)
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range dataSlice {
		fmt.Println(k, v)
	}

}

func main() {

	structtojson()
	jsontomap()
	jsontostruct()
	jsontomapslice()
	jsontostructslice()

}
