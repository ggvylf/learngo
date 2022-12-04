package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Paasword string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Paasword string `ini:"password"`
	Database int    `ini:"database"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {

	//获取结构体的类型
	t := reflect.TypeOf(data)

	//定义子结构体的字段
	var subStructName string

	//判断data的类型是否是否为指针类型
	if t.Kind() != reflect.Ptr {
		err = errors.New("data should be a pointer")
		return
	}

	// // 判断data的值是否为结构体
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data should be a pointer struct")
		return
	}

	//读取配置文件
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}

	//把读取到的内容转换成字符串切片，根据换行符切割
	lineSlice := strings.Split(string(bytes), "\n")
	fmt.Printf("lineSlice= %#v\n", lineSlice)

	for idx, line := range lineSlice {

		//去掉line首位的空格
		line = strings.TrimSpace(line)

		//空行就跳过
		if len(line) == 0 {
			continue
		}

		//判断是否是注释，是的话跳过
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "；") {
			continue
		}

		//已[开头的表示节 section
		if strings.HasPrefix(line, "[") {
			//开头和结尾不是以[]为准，语法错误
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error，not a [xxx]", idx+1)
				return
			}

			//获取section的名称，并去除空格

			sectionName := strings.TrimSpace(line[1 : len(line)-1])

			//如果[]中间的内容为空，语法错误
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error，is []", idx+1)
				return
			}
			// v := reflect.ValueOf(data)
			//根据sectionName去data中找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {

				//获取结构体中的字段
				field := t.Elem().Field(i)
				//获取Config结构体中 对应字段中tag是ini中的内容 例如`ini: "mysql"`

				//找到对应的子结构体的名称，把字段名赋值给子结构体的名称
				if sectionName == field.Tag.Get("ini") {
					subStructName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, subStructName)
				}
			}

			// 不是[xxx] 那就是正常的子结构体的内容
		} else {

			//判断该行是否有等号，没有等号或者是等号再行首 语法错误。
			if strings.Index(line, "=") == -1 || strings.Index(line, "=") == 0 {
				err = fmt.Errorf("line:%d  syntax error，not = or = is first", idx+1)
				return
			}

			//获取=的索引
			index := strings.Index(line, "=")
			//key在等号之前,value在等号之后
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])

			//通过嵌套结构体字段的名称获取对应的值

			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(subStructName)

			//嵌套结构体的类型信息
			sType := sValue.Type()

			//如果sValue的kind不是结构体
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段应该是一个结构体", subStructName)
			}

			//定义具体子结构体中的字段名称
			var fieldName string
			//定义具体子结构体中的字段类型
			var fieldType reflect.StructField

			//遍历嵌套结构体的每个字段
			for i := 0; i < sValue.NumField(); i++ {
				//获取子结构体字段的名称
				field := sType.Field(i)
				fieldType = field

				//判断tag中的内容是否等于key，如果相等就把从数据获取的key的内容赋值给fieldName
				if fieldType.Tag.Get("ini") == key {

					fieldName = field.Name
					break

				}
			}

			//在结构体中找不到对应的字段
			if len(fieldName) == 0 {
				continue
			}

			//从sValue中找到对应fieldName的内容
			fieldObj := sValue.FieldByName(fieldName)

			//根据值类型进行对应的字段填充
			switch fieldType.Type.Kind() {
			case reflect.String:
				fieldObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				//把value转换成int类型
				valueInt, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return err
				}
				fieldObj.SetInt(valueInt)

			case reflect.Bool:
				valueBool, err := strconv.ParseBool(value)
				if err != nil {
					return err
				}
				fieldObj.SetBool(valueBool)
			}

			fmt.Printf("fileldName=%v,fieldObj=%v\n", fieldName, fieldObj)

		}
	}
	return
}

func main() {

	var cfg Config
	file := "./a.ini"
	err := loadIni(file, &cfg)
	if err != nil {
		fmt.Println(err)
	}

}
