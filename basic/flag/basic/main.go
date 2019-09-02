package main

import (
	"flag"
	"fmt"
	"os"
)

func getargs() {
	fmt.Printf("参数=%v,长度=%d\n", os.Args, len(os.Args))
	for k, v := range os.Args {
		fmt.Printf("第%d个参数=%v\n", k, v)
	}
}

func getflag() {
	//返回值是一个指针
	name := flag.String("name", "tom", "姓名")
	age := flag.Int("age", 18, "年龄")
	flag.Parse()
	fmt.Printf("参数name=%s,参数age=%d\n", *name, *age)

	var name2 string
	var age2 int
	//赋值给变量
	flag.StringVar(&name2, "name2", "tom", "姓名")
	flag.IntVar(&age2, "age2", 18, "年龄")
	flag.Parse()
	fmt.Printf("参数name2=%s,参数age2=%d\n", name2, age2)
}

func main() {

	getargs()
	getflag()
}
