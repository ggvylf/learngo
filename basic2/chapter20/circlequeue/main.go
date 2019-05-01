package main

import (
	"errors"
	"fmt"
	"os"
)

//使用结构体管理队列
type CircleQueue struct {
	maxSize int
	array   [5]int
	head    int
	tail    int
}

func (cq *CircleQueue) Push(val int) (err error) {
	if cq.IsFull() {
		return errors.New("queue is full")
	}
	//tail不包含最后一个元素，最后一个元素就是保留的标志位
	cq.array[cq.tail] = val
	cq.tail = (cq.tail + 1) % cq.maxSize

	return
}

func (cq *CircleQueue) Pop() (val int, err error) {
	if cq.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	val = cq.array[cq.head]
	cq.head = (cq.head + 1) % cq.maxSize
	return

}

func (cq *CircleQueue) ListQueue() {
	size := cq.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}

	tempHead := cq.head

	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\n", tempHead, cq.array[i])
		tempHead = (tempHead + 1) % cq.maxSize

	}
	fmt.Println()

}

func (cq *CircleQueue) IsFull() bool {
	return (cq.tail+1)%cq.maxSize == cq.head

}

func (cq *CircleQueue) IsEmpty() bool {
	return cq.tail == cq.head
}

func (cq *CircleQueue) Size() int {
	return (cq.tail + cq.maxSize - cq.head) % cq.maxSize
}

func main() {
	cq := &CircleQueue{
		maxSize: 4,
		head:    0,
		tail:    0,
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
			err := cq.Push(val)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("add ok!")
			}

		case 2:
			val, err := cq.Pop()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("get val=", val)
			}

		case 3:
			cq.ListQueue()
		case 4:
			os.Exit(0)
		}
	}

}
