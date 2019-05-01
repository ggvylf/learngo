package main

import (
	"fmt"
	"sort"
)

//map嵌套
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
	var a map[int]int
	a = make(map[int]int, 10)
	a[1] = 1
	a[2] = 2
	fmt.Println(a)
	fmt.Println(len(a))

	var b map[string]string = map[string]string{
		"key1": "value1",
		"key2": "vaule2",
	}

	fmt.Println(b)

	val, ok := a[2]
	if ok {
		fmt.Println(val)

	}

	muiltMap()
	sliceOfMap()
	sortOfMap()
	switchMaps()

}
