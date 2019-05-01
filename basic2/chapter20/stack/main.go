package main

import (
	"errors"
	"fmt"
)

type Stack struct {

	//最大存放的个数
	MaxTop int
	//栈顶
	Top int
	//模拟栈
	arr [5]int
}

func (s *Stack) Push(val int) (err error) {
	if s.Top == s.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}

	s.Top++
	s.arr[s.Top] = val
	return

}

func (s *Stack) List() {
	if s.Top == -1 {
		fmt.Println("stack empty")
		return
	}

	curTop := s.Top

	for i := curTop; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, s.arr[i])
	}

}

func (s *Stack) Pop() (val int, err error) {
	if s.Top == -1 {
		fmt.Println("stack empty")
		return 0, errors.New("stack empty")
	}

	val = s.arr[s.Top]
	s.Top--
	return val, nil

}

func main() {
	s := &Stack{
		MaxTop: 5,
		Top:    -1,
		arr:    [5]int{},
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	s.List()

	val, err := s.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)

	s.List()

}
