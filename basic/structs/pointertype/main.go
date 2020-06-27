package main 


import (
	"fmt"
)

type Person struct {
	Name,Gender string
}

//函数内部操作的修改的是传入变量的值拷贝
func f1(p Person) {
	p.Gender="female"

}


//传递指针可以直接修改传入变量的值
func f2(p *Person) {
	//语法糖
	//标准写法 (*p).Gender="female"
	p.Gender="female"

}

func main() {
	p:=Person{"tom","male"}

	f1(p)
	fmt.Println("p=",p)



	f2(&p)
	fmt.Println("p=",p)

}