// micro kernel
// 特点
// 易于扩展
// 错误隔离
// 保持架构一致性

// 要点
// 内核包含公共流程或者是通用逻辑
// 将可变或扩展部分规划为扩展点
// 抽象扩展点行为，定义接口
// 利用插件进行扩展

package main

import (
	"fmt"
	"time"
)

func main() {

	agt := NewAgent(100)
	c1 := NewDemoCollect("c1", "1")
	c2 := NewDemoCollect("c2", "2")
	agt.RegisterCollector("c1", c1)
	agt.RegisterCollector("c2", c2)
	agt.Start()
	fmt.Println(agt.Start())
	time.Sleep(1 * time.Second)
	agt.Stop()
	agt.Destroy()

}
