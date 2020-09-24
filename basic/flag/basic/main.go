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

	//赋值给变量
	var name2 string
	var age2 int
	flag.StringVar(&name2, "name2", "tom", "姓名")
	flag.IntVar(&age2, "age2", 18, "年龄")

	//从os.args[1:]中读取参数，这个要放在所有参数定义之后
	flag.Parse()

	fmt.Printf("参数name=%s,参数age=%d\n", *name, *age)
	fmt.Printf("参数name2=%s,参数age2=%d\n", name2, age2)

	fmt.Printf("没有指定flag的命令行参数=%v\n", flag.Args())

	fmt.Printf("没有指定flag命令行参数的个数=%v\n", flag.NArg())

	fmt.Printf("指定flag命令行参数的个数=%v\n", flag.NFlag())
}

func main() {

	//getargs()
	getflag()
}
