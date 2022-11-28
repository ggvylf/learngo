package main

import (
	"fmt"
	"sort"
)

// map嵌套
func muiltMap() {
	var c map[int]map[int]string
	c = make(map[int]map[int]string)
	c[1] = make(map[int]string)
	c[1][1] = "heehe"

	//修改之前先判断，不存在再初始化
	_, ok := c[2]
	if !ok {
		c[2] = make(map[int]string)
	}

	//追加或修改
	c[2][1] = "jajaa"
	c[2][2] = "xxxx"

	fmt.Println(c)

	//删除
	delete(c[2], 2)
	fmt.Println(c)
	delete(c, 2)
	fmt.Println(c)

}

func sliceOfMap() {

	d := make([]map[int]int, 5)
	fmt.Printf("%T\n", d)

	d[0] = make(map[int]int)
	d[0][1] = 1
	fmt.Println(d)

}

func sortOfMap() {
	e := make(map[int]int, 3)
	e[3] = 6
	e[1] = 8
	e[2] = 5

	for k, v := range e {
		fmt.Printf("%d,%d\n", k, v)
	}

	var es []int
	for k, _ := range e {
		es = append(es, k)
	}
	fmt.Println(es)

	sort.Ints(es)
	fmt.Println(es)

	for _, v := range es {
		fmt.Println(v, e[v])
	}

}

func switchMaps() {
	f := make(map[int]string)
	g := make(map[string]int)

	f[1] = "hehe"
	f[2] = "haha"
	fmt.Println(f)

	for k, v := range f {
		g[v] = k
	}
	fmt.Println(g)

}

func main() {

	//map初始化和map的长度
	initandlen()

	//判断map中的元素是否存在
	iskeyexist()

	//map遍历
	rangemap()

	//map的value是一个方法
	mapvalueisfunc()

	//map实现set
	mapasset()

	// muiltMap()
	// sliceOfMap()
	// sortOfMap()
	// switchMaps()

}

func initandlen() {

	m1 := map[int]int{1: 1, 2: 2}
	fmt.Printf("len is %d\n", len(m1))

	m2 := map[int]int{}
	m2[1] = 1
	fmt.Printf("len is %d\n", len(m2))

	//注意 map不能用cap统计容量
	m3 := make(map[int]int, 10)
	fmt.Printf("len is %d\n", len(m3))
}

func iskeyexist() {
	//key不存在的时候 默认返回key类型的nil值
	m1 := map[int]int{}
	m2 := map[int]string{}
	fmt.Printf("m1[1]=%+v,type is %T\n", m1[1], m1[1])
	fmt.Printf("m2[1]=%+v,type is %T\n", m2[1], m2[1])

	//判断key是否存在
	m1[1] = 3
	if v, ok := m1[1]; ok {
		fmt.Printf("key 1 is exists,value=%v\n", v)
	} else {
		fmt.Println("key 2 is not exist")
	}

	if v, ok := m1[2]; ok {
		fmt.Printf("key 2 is exists,value=%v\n", v)
	} else {
		fmt.Println("key 2 is not exist")
	}

}

func rangemap() {
	m1 := map[string]int{
		"one": 1,
		"two": 2,
	}

	for k, v := range m1 {
		fmt.Printf("key is %s,value is %d\n", k, v)
	}

}

func mapvalueisfunc() {

	m := map[int]func(op int) int{}

	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}

	m[3] = func(op int) int {
		return op * op * op
	}

	fmt.Println(m[1](2), m[2](2), m[3](2))

}

func mapasset() {
	//声明set map[type]bool
	mySet := map[int]bool{}

	//判断元素是否存在，只要判断value即可
	mySet[1] = true
	if mySet[1] {
		fmt.Println("mySet[1] is exist")
	} else {
		fmt.Println("mySet[1] is not exist")
	}

	//set的长度
	mySet[2] = true
	fmt.Printf("mySet len is %d\n", len(mySet))

	//删除元素
	n := 1
	delete(mySet, n)
	if mySet[n] {
		fmt.Printf("mySet[%d] is exist\n", n)
	} else {
		fmt.Printf("mySet[%d] is not exist\n", n)
	}
}
