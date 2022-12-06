//agent用来加载多个collector，并负责调用collector

package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
)

const (
	// agent的运行状态
	Waiting = iota
	Running
)

var WrongStateError = errors.New("can not do operation in the current state")

// 错误集
type CollectorsError struct {
	CollectorErrors []error
}

// 拼接错误信息
func (ce CollectorsError) Error() string {
	var strs []string
	for _, err := range ce.CollectorErrors {
		strs = append(strs, err.Error())
	}
	return strings.Join(strs, ";")

}

// 事件结构体
type Event struct {
	Source  string
	Content string
}

type Agent struct {
	// 收集器的map集合
	collectors map[string]Collector
	// 事件队列
	evtBuf chan Event
	cancel context.CancelFunc
	ctx    context.Context
	// collector的状态
	state int
}

// 创建agent
func NewAgent(sizeEventBuf int) *Agent {
	agt := Agent{
		collectors: map[string]Collector{},
		evtBuf:     make(chan Event, sizeEventBuf),
		state:      Waiting,
	}
	return &agt
}

// 注册收集器
func (agt *Agent) RegisterCollector(name string, collector Collector) error {
	// agent的状态不是等待状态，禁止向agent注册收集器
	if agt.state != Waiting {
		return WrongStateError
	}
	// 把收集器添加到收集器列表
	agt.collectors[name] = collector

	// 初始化 收集器
	return collector.Init(agt)
}

// 启动收集器
func (agt *Agent) startCollectors() error {
	var err error
	var errs CollectorsError
	var mutex sync.Mutex

	// 遍历收集器列表
	for name, collector := range agt.collectors {
		go func(name string, collector Collector, ctx context.Context) {
			defer func() {
				mutex.Unlock()
			}()

			// 启动收集器
			err = collector.Start(ctx)
			mutex.Lock()

			// 把启动失败的collector的错误日志添加到日志集中
			if err != nil {
				errs.CollectorErrors = append(errs.CollectorErrors, errors.New(name+":"+err.Error()))
			}
		}(name, collector, agt.ctx)

	}
	// 返回错误日志集
	return errs
}

// 停止所有收集器
func (agt *Agent) stopCollectors() error {
	var err error
	var errs CollectorsError
	for name, collector := range agt.collectors {
		if err = collector.Stop(); err != nil {
			errs.CollectorErrors = append(errs.CollectorErrors, errors.New(name+":"+err.Error()))
		}
	}
	return errs
}

// 销毁所有收集器
func (agt *Agent) destroyCollectors() error {
	var err error
	var errs CollectorsError
	for name, collector := range agt.collectors {
		if err = collector.Destroy(); err != nil {
			errs.CollectorErrors = append(errs.CollectorErrors, errors.New(name+":"+err.Error()))
		}
	}
	return errs
}

// 批量事件打印
func (agt *Agent) EventProcessGroutine() {

	//10个事件作为一个组
	var evtSeg [10]Event
	for {
		for i := 0; i < 10; i++ {
			select {
			// 从事件队列里读取事件
			case evtSeg[i] = <-agt.evtBuf:

			// 等待上下文信号
			case <-agt.ctx.Done():
				return
			}
		}
		fmt.Println(evtSeg)
	}
}

// 启动agent
func (agt *Agent) Start() error {
	if agt.state != Waiting {
		return WrongStateError
	}
	agt.state = Running

	//传递上下文
	agt.ctx, agt.cancel = context.WithCancel(context.Background())

	//并发打印事件Event
	go agt.EventProcessGroutine()

	//批量启动收集器
	return agt.startCollectors()

}

// 停止agent
func (agt *Agent) Stop() error {
	// 状态错误
	if agt.state != Running {
		return WrongStateError
	}
	//变更状态
	agt.state = Waiting

	// 停止收集器中运行的任务
	agt.cancel()

	//停止所有收集器
	return agt.stopCollectors()

}

// 销毁agent
func (agt *Agent) Destroy() error {
	// 运行状态的不能销毁
	if agt.state != Waiting {
		return WrongStateError
	}
	return agt.destroyCollectors()
}

// 事件写入事件队列
func (agt *Agent) OnEvent(evt Event) {
	agt.evtBuf <- evt
}
