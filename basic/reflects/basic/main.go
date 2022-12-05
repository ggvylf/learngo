package main

import (
	"fmt"
	"reflect"
)

type Cat struct{}
type Dog struct{}

// 使用反射判断接口类型
func ReflectType(x interface{}) {
	//如何判断接口类型 1，类型断言 2. 反射
	//接口=动态接口类型+对应动态接口类型的值
	//接口类型可以通过Reflect.TypeOf()获取
	//对应接口类型的值可以通过reflect.ValueOf()获取

	inftype := reflect.TypeOf(x)

	//Name()表示的是接口名称，kind()表示接口的类型
	//slice map pointer array等类型的Name()都是返回空
	fmt.Printf("变量%v的 接口类型=%v,接口名称=%v,接口种类=%v\n", x, inftype, inftype.Name(), inftype.Kind())

}

// 获取变量的类型
func GetType() {
	ReflectType(3)
	ReflectType([]int{1, 2, 3})
	var c Cat
	ReflectType(c)
	var d Dog
	ReflectType(d)
}

// 通过反射获取接口对应类型的值
func ReflectValue(x interface{}) {
	infval := reflect.ValueOf(x)
	fmt.Printf("接口反射值=%v,接口反射值的类型=%v\n", infval, infval.Kind())

	//判断infval的值的类型并转换成对应的类型的值
	switch infval.Kind() {
	//和reflect包中定义的类型比较
	case reflect.Float64:
		{
			//通过反射转换原始值
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

// 获取实例对象的值和类型
func GetValue() {
	// float64类型要转换成float32类型
	ReflectValue(3.33)
	ReflectValue(123)
}

// 通过反射修改变量的值，注意函数传递的是值的拷贝，必须传地址
func ReflectSetValue(x interface{}) {

	//通过反射获取值
	infval := reflect.ValueOf(x)

	//Elem()方法返回v的接口或者是指针，如果v的Kind()不是interface{}或者是ptr会panic,如果v是nil则返回0
	//通过Elem()方法才能获取指针对应的值
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

	// int类型修改成200
	// float64类型修改成4.4444
	var a int32 = 123
	var b float64 = 3.33

	//函数是值的拷贝，必须传地址。
	fmt.Println("a修改前值=", a)
	ReflectSetValue(&a)
	fmt.Println("a修改反射后的值=", a)

	fmt.Println("b修改前值=", b)
	ReflectSetValue(&b)
	fmt.Println("b修改反射后的值=", b)
}

// IsNil()会判断变量持有的值是否为nil，注意变量的的类型必须是chan func interface map pointer slice之一，否则会painc
func ReflectIsNil() {
	var a *int

	fmt.Println("指针a的值否为空=", reflect.ValueOf(a).IsNil())
	//基本变量会panic
	// var b int
	// fmt.Println("int类型b的值是否为空=", reflect.ValueOf(b).IsNil())
}

// isVaild()会返回变量是否持有一个值，如果值是0则会返回false，此时该值除了IsVaild() String Kind以外的方法都会panic
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

	//判断sturct否存在某个方法
	p2 := struct{}{}
	if reflect.ValueOf(p2).MethodByName("abc").IsValid() {
		fmt.Println("存在abc方法")
	} else {
		fmt.Println("不存在abc方法")
	}

	//判断map中是否存在key
	k1 := map[int]int{1: 11, 2: 22, 3: 33}
	fmt.Println(k1)

	//这里要把int类型的转换为Value类型
	if reflect.ValueOf(k1).MapIndex(reflect.ValueOf(4)).IsValid() {
		fmt.Println("index为4的元素存在")
	} else {
		fmt.Println("index为4的元素不存在")
	}
}
func main() {
	//反射指的是程序在运行时，可以访问、检测和修改它本身的状态或行为的一种能力。

	GetType()
	GetValue()
	SetValue()
	ReflectIsNil()
	ReflectIsValid()
}
