package main

import (
	"errors"
	"fmt"
	"os"
)

//使用结构体管理队列
type Queue struct {
	maxSize int
	array   [5]int
	head    int
	tail    int
}

func (q *Queue) AddQqueue(val int) (err error) {

	//先判断队列是否满
	if q.tail == q.maxSize-1 {
		return errors.New("queue is full")
	}

	//增加后tail要移动
	q.tail++
	//数据放入数组
	q.array[q.tail] = val
	return
}

//显示队列
func (q *Queue) ShowQueue() {
	for i := q.head + 1; i <= q.tail; i++ {
		fmt.Printf("array[%d]=%d\n", i, q.array[i])
	}
	fmt.Println("--------------------")
}

func (q *Queue) Getqueue() (val int, err error) {
	//判断是否队列为空
	if q.head == q.tail {
		return -1, errors.New("array is empty")
	}
	//取出后head要移动
	q.head++

	val = q.array[q.head]
	return

}

func main() {
	q := &Queue{
		maxSize: 5,
		head:    -1,
		tail:    -1,
	}

	var key int
	var val int
	for {
		fmt.Println("------------------")
		fmt.Println("1.add key to queue")
		fmt.Println("2.get key")
		fmt.Println("3.show queue")
		fmt.Println("4.exit")
		fmt.Println("------------------")
		fmt.Scanln(&key)

		switch key {
		case 1:
			fmt.Println("input val:")
			fmt.Scanln(&val)
			err := q.AddQqueue(val)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("add ok!")
			}

		case 2:
			val, err := q.Getqueue()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("get val=", val)
			}

		case 3:
			q.ShowQueue()
		case 4:
			os.Exit(0)
		}
	}

}
