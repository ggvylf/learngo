package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini: "address"`
	Port     int    `ini: "port"`
	Username string `ini: "username"`
	Paasword string `ini: "password"`
}

type RedisConfig struct {
	Host     string `ini: "host"`
	Port     int    `ini: "port"`
	Paasword string `ini: "password"`
	Database int    `ini: "database"`
}

type Config struct {
	MysqlConfig `ini: "mysql"`
	RedisConfig `ini: "redis"`
}

func loadIni(fileName string, data interface{}) (err error) {

	t := reflect.TypeOf(data)

	//判断data的类型是否是否为指针类型
	if t.Kind() != reflect.Ptr {
		err = errors.New("data should be a pointer")
		return
	}

	// // 判断data的值是否为结构体
	// 	t.Elem().Kind() != reflect.Struct {
	// 		err = errors.New("data should be a pointer struct")
	// 		return
	// 	}

	//读取配置文件
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}

	//把读取到的内容转换成字符串，根据换行符切割
	lineSlice := strings.Split(string(b), "\n")
	fmt.Printf("lineSlice=%#v\n", lineSlice)

	var structName string

	for idx, line := range lineSlice {

		//去掉line首位的空格
		line = strings.TrimSpace(line)

		//判断是否是注释，是的话跳过
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "；") {
			continue
		}

		//已[开头的表示节 section
		if strings.HasPrefix(line, "[") {
			//开头和结尾不是以[]为准，语法错误
			if line[0] != '[' && line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}

			//如果[]中间的内容为空，语法错误
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
			}

			//根据sectionName去data中找到对应的结构体

			for i := 0; i < t.Elem().NumField(); i++ {

				field := t.Elem().Field(i)
				fmt.Println(field.Tag.Get("ini"))

				//找到对应的结构体名称，把字段名字记录下来
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}

		} else {

		}
		return

	}
}

func main() {

	var cfg Config
	file := "./a.ini"
	err := loadIni(file, &cfg)
	if err != nil {
		fmt.Println(err)
	}

}
