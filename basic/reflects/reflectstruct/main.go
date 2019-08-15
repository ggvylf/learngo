package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"jsonname" xml:"xmlname"`
	Score int    `json:"score"`
}

func main() {
	s1 := Student{"tom", 99}
	stureftype := reflect.TypeOf(s1)
	//打印struct的名称和类型
	fmt.Printf("struct的名称=%v,struct的类型=%v\n", stureftype.Name(), stureftype.Kind())

	//获取struct中字段的个数
	fmt.Println("struct中的字段个数=", stureftype.NumField())

	//遍历struct中的字段和tag
	for i := 0; i < stureftype.NumField(); i++ {
		obj := stureftype.Field(i)
		fmt.Printf("字段的名称=%v，字段的类型=%v，字段的tag=%v\n", obj.Name, obj.Type, obj.Tag)
		//取对应tag的值
		fmt.Printf("tag为json的值=%v,tag为xml的值=%v\n", obj.Tag.Get("json"), obj.Tag.Get("xml"))
	}
	//根据字段名字取值
	v, _ := stureftype.FieldByName("Score")
	fmt.Println("Score字段的值=")

}
