package main

import (
	"reflect"
	"testing"
)

type user struct {
	UesrId string
	Name   string
}

func TestReflectStruct(t *testing.T) {
	var (
		model *user
		st    reflect.Type
		elem  reflect.Value
	)

	//获取*user类型的Type
	st = reflect.TypeOf(model)
	t.Log("reflect.TypeOf", st.Kind().String())

	//获取st指向的字段类型
	st = st.Elem()

	//返回一个Value类型，包含了指向Type的指针
	elem = reflect.New(st)

	t.Log("reflect.New", elem.Kind().String())
	t.Log("reflect.New.Elem", elem.Elem().Kind().String())

	//通过该类型断言把interface{}转换成ptr类型
	model = elem.Interface().(*user)

	//拿到elem指向的值
	elem = elem.Elem()

	//给字段赋值
	elem.FieldByName("UserId").Setstring("007")
	elem.FieldByName("Name").Setstring("tommy")

	t.Log("model.Name=", model.Name)
}
