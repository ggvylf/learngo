package main

import "fmt"

func main() {
	//值类型 普通变量
	a := 3

	//& 对普通变量取内存地址
	fmt.Printf("a的类型=%T,a=%v，内存地址&a=%p\n", a, a, &a)

	//定义一个指针类型，并且赋值为a的内存地址
	var p *int = &a

	//输出指针类型的值，p的值是变量a的内存地址，&p是p在内存中的地址，*p是p的值的地址中存储的数据,即a
	fmt.Printf("指针p的类型=%T,p=%v，内存地址&p=%p,存储指针对应的值*p=%v\n", p, p, &p, *p)

	//指针变量p内存地址对应的值，等于p的值
	fmt.Printf("对指针的地址取值，指针的值是一个内存地址，即p,p=%v，*&p=%v\n",p, *&p)
}
