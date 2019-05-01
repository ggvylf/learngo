package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

//定义一个结构体的slice
type HeroSlice []Hero

//实现sort.Sort的方法

func (hs HeroSlice) Len() int {
	return len(hs)
}

//排序最核心的方法，制定要排序的元素和升降序
// func (hs HeroSlice) Less(i, j int) bool {
// 	return hs[i].Age < hs[j].Age
// }

func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Name < hs[j].Name
}

func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	var hs HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("hero%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		hs = append(hs, hero)
	}

	fmt.Println("before sort:")
	for _, v := range hs {
		fmt.Println(v)
	}
	sort.Sort(hs)
	fmt.Println("after sort:")

	for _, v := range hs {
		fmt.Printf("%v\n", v)
	}
}
