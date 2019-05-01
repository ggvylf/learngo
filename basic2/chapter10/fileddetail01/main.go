package main

import "fmt"

type Persion struct {
	Name  string
	Age   int
	Score [5]float64
	ptr   *int
	slice []int
	map1  map[string]string
}

func main() {
	var p1 Persion
	fmt.Println(p1)
	p1.Name = "tom"
	p1.Score = [5]float64{1.0, 2.0, 3.0, 4.0, 5.0}
	p1.slice = make([]int, 5)
	p1.map1 = make(map[string]string)
	p1.map1["hello"] = "world"
	fmt.Println(p1)

}
