package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"jsonname" xml:"xmlname"`
	Score int    `json:"score"`
}

func (s Student) Study(a, b int) {
	str := "is study"
	fmt.Println(str)
	fmt.Println(a + b)

}

func (s Student) Sleep(a, b string) {
	str := "is sleeping"
	fmt.Println(str)
	fmt.Println(a + b)
}

func PrintMethod(x interface{}) {
	xtype := reflect.TypeOf(x)
	xvalue := reflect.ValueOf(x)
	//获取方法个数
	fmt.Println("结构体包含的方法个数=", xtype.NumMethod())

	//遍历所有的方法名称和类型，这里要使用值的反射，才能获取到具体的方法来调用
	for i := 0; i < xvalue.NumMethod(); i++ {
		methodtype := xvalue.Method(i).Type()
		fmt.Printf("结构体中方法的名称=%v，方法的类型=%v\n", xtype.Method(i).Name, methodtype)

	}

	args := []reflect.Value{}
	strargs := []reflect.Value{}
	args = append(args, reflect.ValueOf(1))
	args = append(args, reflect.ValueOf(3))
	strargs = append(strargs, reflect.ValueOf("2"))
	strargs = append(strargs, reflect.ValueOf("4"))
	//调用方法
	xvalue.Elem().Method(0).Call(args)
	xvalue.Elem().Method(1).Call(strargs)

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
	//根据字段名字查找对应字段
	s1score, _ := stureftype.FieldByName("Score")
	fmt.Println("s1的Score字段=", s1score)

	PrintMethod(s1)

}
