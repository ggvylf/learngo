package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack struct {

	//最大存放的个数
	MaxTop int
	//栈顶
	Top int
	//模拟栈
	arr [20]int
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

func (s *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

func (s *Stack) Calc(num1, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("oper is wrong")
	}

	return res
}

func (s *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}

func main() {

	numstack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}

	operstack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}

	exp := "38+62-10"

	index := 0

	num1 := 0
	num2 := 0
	oper := 0
	res := 0

	keepNum := ""

	for {
		//拿到一个字符
		ch := exp[index : index+1]

		//把拿到的字符转成int类型的ascii码
		temp := int([]byte(ch)[0])

		//是否是运算符
		if operstack.IsOper(temp) {

			//符号栈是空栈直接入栈
			if operstack.Top == -1 {
				operstack.Push(temp)
			} else {
				//判断temp和符号栈top的优先级
				if operstack.Priority(operstack.arr[operstack.Top]) >= operstack.Priority(temp) {
					//弹出数字和运算符
					num1, _ = numstack.Pop()
					num2, _ = numstack.Pop()
					oper, _ = operstack.Pop()

					//计算结果并重新入数字栈,把当前的符号入符号栈
					res = operstack.Calc(num1, num2, oper)
					numstack.Push(res)
					operstack.Push(temp)
				} else {
					//优先级相同，入符号栈
					operstack.Push(temp)
				}
			}

		} else {

			//多位数字拼接
			keepNum += ch

			//处理多位数字
			if index == len(exp)-1 {
				//数字直接入数字栈,//把ascii值改成int
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numstack.Push(int(val))
			} else {
				//判断index的下一位是不是操作符，如果不是那就是数字，需要拼在ch后边
				if operstack.IsOper(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numstack.Push(int(val))
					//清空keepNum
					keepNum = ""
				}

			}

		}

		//判断index是否移动到exp末尾
		if index == len(exp)-1 {
			break

		} else {
			//移动index
			index++
		}

	}

	//清空栈
	for {
		if operstack.Top == -1 {
			break
		} else {
			num1, _ = numstack.Pop()
			num2, _ = numstack.Pop()
			oper, _ = operstack.Pop()

			//计算结果并重新入数字栈,把当前的符号入符号栈
			res = operstack.Calc(num1, num2, oper)
			numstack.Push(res)
		}
	}

	res, _ = numstack.Pop()
	fmt.Printf("exp=%v,res=%d\n", exp, res)

}
