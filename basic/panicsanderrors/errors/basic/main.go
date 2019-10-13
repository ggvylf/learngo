package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Myerr struct {
	Time   time.Time
	ErrMsg string
	ErrNo  int
}

func (m *Myerr) Error() string {
	info := fmt.Sprintf("time=%v  type=%s no=%d\n", m.Time, m.ErrMsg, m.ErrNo)
	return info
}

func testpanic(i int) (err error) {
	if i < 0 {
		panic(&Myerr{time.Now(), "不能为负数", 1})
	}

	if i > 50 || i < 5 {
		err = &Myerr{time.Now(), "超出范围", 3}
	}

	return

}

func main() {
	//defer func() {
	//	if r:=recover();r!=nil{
	//		fmt.Println("recover")
	//	}
	//}()
	i, _ := strconv.Atoi(os.Args[1])
	err := testpanic(i)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("程序执行结束")

}
