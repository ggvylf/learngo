// https://go.dev/doc/tutorial/generics
// https://gobyexample.com/generics

package main

import "fmt"

// 泛型函数的固定格式
// K 是comparable接口类型，实现了所有可进行比较操作==和!=的数据类型
// V 如果是使用any表示任意类型，或者是可以通过|来限制具体的类型
// func xxxx[K comparable, V int64|float64] {} {
//      dosometing
// }

// 接收K类型的key 和任意类型的value V，并返回K的slice
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// List链表，用来存放element的头尾直针
type List[T any] struct {
	head, tail *element[T]
}

// elemetn的结构体，支持任意类型的key和value
type element[T any] struct {
	next *element[T]
	val  T
}

// 把元素填充到list
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// 从list中遍历所有元素
func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	// 声明一个map key是int，value是string

	var m = map[int]string{
		1: "2",
		2: "4",
		4: "8",
	}

	// 直接调用，不指定K和V的类型
	fmt.Println("keys:", MapKeys(m))
	fmt.Printf("keys type: %#v\n", MapKeys(m))

	// 指定K和V类型的调用
	_ = MapKeys[int, string](m)

	// int类型的List
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())

	// string类型的List
	lst2 := List[string]{}
	lst2.Push("hello")
	lst2.Push("my")
	lst2.Push("world")
	lst2.Push("golang")
	fmt.Println("list2:", lst2.GetAll())
}
