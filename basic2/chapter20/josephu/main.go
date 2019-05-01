package main

import "fmt"

type Boy struct {
	no   int
	next *Boy
}

func AddBoy(sum int) *Boy {

	first := &Boy{}
	curBoy := &Boy{}

	if sum < 1 {
		fmt.Println("sum至少大于2")
		return first
	}

	for i := 1; i <= sum; i++ {
		boy := &Boy{
			no: i,
		}

		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.next = first
		} else {
			curBoy.next = boy
			curBoy = boy
			curBoy.next = first

		}
	}
	return first
}

func ListBoy(first *Boy) {
	if first.next == nil {
		fmt.Println("empty")
		return
	}

	curBoy := first
	for {
		fmt.Printf("[no=%d]--->", curBoy.no)

		if curBoy.next == first {
			fmt.Printf("tail")
			break
		}
		curBoy = curBoy.next
	}
	fmt.Println()

}

func PlayGame(first *Boy, startNo, countNum int) {
	if first.next == nil {
		fmt.Println("empty")
		return
	}

	// if startNo <= BoysCount {
	// 	fmt.Println("startNo超过总数"))
	// 	return
	// }

	//tail指向最后的元素
	tail := first

	for {
		if tail.next == first {
			break
		}
		tail = tail.next
	}

	//first移动到startNo，循环的次数=startN-1
	for i := 1; i <= startNo-1; i++ {
		//tail跟随first移动
		first = first.next
		tail = tail.next

	}

	//数的次数=countNmu-1

	for {
		//开始数
		for i := 1; i <= countNum-1; i++ {
			//移动first和tail
			first = first.next
			tail = tail.next

		}
		fmt.Printf("no=%d is out\n", first.no)

		//将数到的no移出去
		first = first.next
		tail.next = first

		if tail == first {
			fmt.Println("只有一个了")
			break
		}

	}
	fmt.Println("最后的no=", first.no)
}

func main() {
	first := AddBoy(500)

	ListBoy(first)

	PlayGame(first, 20, 31)

}
