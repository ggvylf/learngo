// 包的名称
// 1. 一个文件夹下只能有一个包，同样一个包的文件不能在多个目录下
// 2. 包名可以喝文件夹的名字不同。报名不能包含短横线-
// 3. 包名为main的包是应用程序的入口。编译不包含main包的源代码不会得到可执行文件。
package add

import "fmt"

func init() {
	fmt.Println("add.go init func")
}

// 可见性问题：引用该包或者包中的变量、常量、函数等，只要首字母大写表示可以被外部引用。
func Add(a, b int) int {
	return a + b
}
