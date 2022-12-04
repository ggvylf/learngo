package main

import (
	"errors"
	"fmt"
	"reflect"
)

// 填充不同struct的相同字段

type Person struct {
	Name string `json:name`
	Age  int    `json:age`
}

type Student struct {
	Name  string
	Age   int
	Class string
}

// 填充函数，入参是接口类型和对应struct值的map
func fill(x interface{}, setting map[string]interface{}) error {

	//先判断传进来的x是不是struct类型
	if reflect.TypeOf(x).Kind() != reflect.Ptr {
		if reflect.TypeOf(x).Elem().Kind() != reflect.Struct {
			return errors.New("type not struct")

		}

	}

	// 判断setting是否为空
	if setting == nil {
		return errors.New("setting is nil")
	}

	var (
		field reflect.StructField
		ok    bool
	)

	//从setting里取kv 并且赋值
	for k, v := range setting {
		if field, ok = (reflect.ValueOf(x).Elem().Type().FieldByName(k)); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(x)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}

	return nil

}

func main() {
	set := map[string]interface{}{
		"Name": "tom",
		"Age":  18,
	}

	p := &Person{}
	fill(p, set)
	fmt.Println(p)
	s := &Student{}
	fill(s, set)
	fmt.Println(s)

}
