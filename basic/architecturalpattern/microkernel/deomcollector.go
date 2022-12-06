// 创建一个collector

package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// 收集器结构体
type DemoCollector struct {
	eventReceiver EventReceiver
	agtCtx        context.Context
	stopChan      chan struct{}
	name          string
	content       string
}

// 创建收集器实例
func NewDemoCollect(name string, content string) *DemoCollector {
	return &DemoCollector{
		stopChan: make(chan struct{}),
		name:     name,
		content:  content,
	}
}

// 初始化
func (c *DemoCollector) Init(evtReceiver EventReceiver) error {
	fmt.Println("init collector", c.name)
	c.eventReceiver = evtReceiver
	return nil
}

// 启动收集器
func (c *DemoCollector) Start(agtCtx context.Context) error {
	fmt.Println("start collect", c.name)
	for {
		select {
		// agent的取消信号
		case <-agtCtx.Done():
			// 往stop通道里写入数据
			c.stopChan <- struct{}{}
			break
		default:
			time.Sleep(50 * time.Millisecond)
			// 写入日志
			c.eventReceiver.OnEvent(Event{c.name, c.content})
		}
	}
}

// 停止收集器
func (c *DemoCollector) Stop() error {
	fmt.Println("stop collect", c.name)
	select {
	// 收到停止通道的消息
	case <-c.stopChan:
		return nil
	// 超时 停止失败
	case <-time.After(1 * time.Second):
		return errors.New("failed to stop for timeout")

	}
}

// 销毁收集器
func (c *DemoCollector) Destroy() error {
	fmt.Println(c.name, "released resources.")
	return nil
}
