package main

import "context"

// 定义事件接收者 接口
type EventReceiver interface {
	// 处理事件函数
	OnEvent(evt Event)
}

// 定义收集器接口
type Collector interface {
	// 初始化
	Init(evtRecevier EventReceiver) error

	// 启动收集器
	Start(agtCtx context.Context) error

	// 停止
	Stop() error

	//销毁
	Destroy() error
}
