package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

func (s Student) Hehe() {
	fmt.Println("hehe")
}

//方法要大写，否则val.NumMethod()获取不到
func (s Student) Set(name string, age int, score float32) {
	s.Name = name
	s.Age = age
	s.Score = score
	fmt.Println(s)
}

func test(b interface{}) {
	//变量的类型，返回reflect.Type类型
	t := reflect.TypeOf(b)
	fmt.Println(t)

	//要分析啥，先转成reflect.ValueOf类型，然后在进行其他操作
	//变量的值，返回reflect.ValueOf类型
	val := reflect.ValueOf(b)
	fmt.Println(val)

	//变量的类别，返回一个常量
	k := val.Kind()
	fmt.Println(k)

	//转换成接口类型
	i := val.Interface()
	fmt.Println(i)

	stu, ok := i.(Student)
	if ok {
		fmt.Printf("%v  %T\n", stu, stu)
	}

}

func testInt(b interface{}) {
	val := reflect.ValueOf(b)
	c := val.Int()
	fmt.Printf("%d\n", c)
}

//运行时修改变量
func setInt(b interface{}) {
	val := reflect.ValueOf(b)
	//用Elem()对指针进行操作，等同于使用*，这里设置一个int类型
	val.Elem().SetInt(100)
	fmt.Println(val.Elem().Int())
}

//操作struct
func testStruct(b interface{}) {
	val := reflect.ValueOf(b)
	//判断类型是否是struct
	k := val.Kind()
	if k != reflect.Struct {
		fmt.Println("not struct!!")
	}
	//输出struct中字段的数量和字段对应的值
	f := val.NumField()
	fmt.Println(f)
	for i := 0; i < f; i++ {
		fmt.Printf("%d,%v，%v\n", i, val.Field(i), val.Field(i).Kind())
	}

	//输出struct对应的方法的数量
	f2 := val.NumMethod()
	fmt.Println(f2)

	//调用struct的方法
	f3 := val.MethodByName("Hehe")
	params := []reflect.Value{}

	f3.Call(params)

}

func main() {
	a := 1233
	test(a)

	b := Student{"tom", 18, 22.4}
	test(b)

	testInt(a)

	fmt.Println(a)
	//注意传地址
	setInt(&a)
	fmt.Println(a)

	testStruct(b)

}
