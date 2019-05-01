package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `json:"username"`
	Age  int    `json:"userage"`
}

func (s Student) Print() {
	fmt.Println("----start----")
	fmt.Println(s)
	fmt.Println("-----end-----")
}

func (s Student) GetSum(m, n int) int {
	return m + n
}

func (s Student) Set(name string, age int) {
	s.Name = name
	s.Age = age
}

func TestStruct(i interface{}) {
	rType := reflect.TypeOf(i)
	rVal := reflect.ValueOf(i)
	rKind := rVal.Kind()

	if rKind != reflect.Ptr && rVal.Elem().Kind() == reflect.Struct {
		fmt.Println("接收的变量不是struct，退出")
		return
	}

	//获取结构体的字段个数
	num := rVal.Elem().NumField()
	fmt.Printf("struct有%d个字段\n", num)

	//运行时修改字段的值
	rVal.Elem().Field(0).SetString("jerry")

	//获取字段的值和对应的tag
	for i := 0; i < num; i++ {
		fmt.Printf("第%d个字段的值=%v,对应的tag=%v\n", i+1, rVal.Elem().Field(i), rType.Elem().Field(i).Tag.Get("json"))
	}

	//获取struct的方法个数
	numOfMethod := rVal.Elem().NumMethod()
	fmt.Printf("struct有%d个方法\n", numOfMethod)

	//调用第2个方法，index顺序按照文件名排序
	rVal.Elem().Method(1).Call(nil)

	//调用第一个方法
	//创建切片，构造传入参数
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))

	//调用对应方法，注意res是个切片
	res := rVal.Method(0).Call(params)
	fmt.Println("调用struct方法的结果=", res[0].Int())

}

func main() {
	s := Student{"tom", 18}
	TestStruct(&s)

}
