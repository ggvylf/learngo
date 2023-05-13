// pipe-filter模式
// 非常适合数据处理和数据分析系统
// filter 封装数据处理的功能
// 松耦合 filter指根数据的格式耦合
// pipe 用于连接filter传递数据，或者是在异步处理中缓冲数据流
// 进程内同步调用的时候，pipe演变胃数据在方法调用之间传递

// 简单的例子
// 传入一个string类型的数据
// SplitFilter 负责根据分隔符分割数据，转换成[]string
// ToIntFilter 负责把[]string类型转换成[]int类型
// SumFilter 负责计算[]int 中元素的和

package main

import "fmt"

type Request interface{}
type Response interface{}

type Filter interface {
	Porcess(data Request) (Response, error)
}

func main() {

	// 创建Filter实例
	spliter := NewSplitFilter(",")
	conventer := NewToIntFilter()
	sumer := NewSumFilter()

	sp := NewStraightPipeline("p1", spliter, conventer, sumer)
	ret, err := sp.Porcess("1,2,3")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)

}
