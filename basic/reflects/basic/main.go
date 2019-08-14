package main

import (
	"fmt"
	"reflect"
)

type Cat struct{}
type Dog struct{}

func ReflectType(x interface{}) {
	//如何判断接口类型 1，类型断言 2. 反射
	inftype := reflect.TypeOf(x)
	//Name()表示的是接口名称，kind()表示接口的类型
	//slice map pointer array等类型的Name()都是返回空
	fmt.Printf("接口类型=%v,接口名称=%v,接口类型=%v\n", inftype, inftype.Name(), inftype.Kind())
}

func GetType() {
	ReflectType(3)
	ReflectType([]int{1, 2, 3})
	var c Cat
	ReflectType(c)
	var d Dog
	ReflectType(d)
}

func ReflectValue(x interface{}) {
	infval := reflect.ValueOf(x)
	fmt.Printf("接口反射值=%v,接口反射值的类型=%v\n", infval, infval.Kind())

	//判断infval的值的类型
	switch infval.Kind() {
	//和reflect包中定义的类型比较
	case reflect.Float64:
		{
			//转换值
			ret := float32(infval.Float())
			fmt.Printf("转换后的值=%v,转换后的类型=%T\n", ret, ret)
		}
	case reflect.Int32:
		{
			ret := int32(infval.Int())
			fmt.Printf("转换后的值=%v,转换后的类型=%T\n", ret, ret)
		}
	}

}

func GetValue() {
	ReflectValue(3.33)
	ReflectValue(123)
}

func ReflectSetValue(x interface{}) {
	//通过反射获取值
	infval := reflect.ValueOf(x)

	//通过Elem()方法获取对应指针的值
	switch infval.Elem().Kind() {
	case reflect.Float64:
		{
			infval.Elem().SetFloat(4.444)
		}
	case reflect.Int32:
		{
			infval.Elem().SetInt(200)
		}
	}

}

func SetValue() {
	var a int32 = 123
	var b float64 = 3.33
	ReflectSetValue(&a)
	fmt.Println("a修改反射后的值=", a)
	ReflectSetValue(&b)
	fmt.Println("b修改反射后的值=", b)
}

func ReflectIsNil() {
	var a *int
	//IsNil()会判断变量持有的值是否为nil，注意变量的的类型必须是chan func interface map pointer slice之一，否则会painc
	fmt.Println("指针a的值否为空=", reflect.ValueOf(a).IsNil())
	//基本变量会panic
	// var b int
	// fmt.Println("int类型b的值是否为空=", reflect.ValueOf(b).IsNil())
}

func ReflectIsValid() {
	type Person struct {
		Name string
		Age  int
	}

	p1 := Person{"tom", 11}
	//IsValid()判断返回值是否有效
	//FieldByName()判断struct中是否存在名称为Name的字段。
	if reflect.ValueOf(p1).FieldByName("Name").IsValid() {
		fmt.Println("存在字段Name")
	}

	//判断是否存在某个方法
	p2 := struct{}{}
	if reflect.ValueOf(p2).MethodByName("abc").IsValid() {
		fmt.Println("存在acb方法")
	} else {
		fmt.Println("不存在acb方法")
	}

	//判断map中是否存在key
	k1 := map[int]int{1: 11, 2: 22, 3: 33}
	//这里要把int类型的转换为Value类型
	if reflect.ValueOf(k1).MapIndex(reflect.ValueOf(4)).IsValid() {
		fmt.Println("index为4的元素存在")
	} else {
		fmt.Println("index为4的元素不存在")
	}
}
func main() {
	GetType()
	GetValue()
	SetValue()
	ReflectIsNil()
	ReflectIsValid()
}
