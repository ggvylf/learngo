package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	//初始化配置
	config := sarama.NewConfig()
	//是否需要ack
	config.Producer.RequiredAcks = sarama.WaitForAll
	//分区,这里配置的是随机
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//返回
	config.Producer.Return.Successes = true

	//初始化生产者
	addr := []string{"192.168.56.11:9092"}
	client, err := sarama.NewSyncProducer(addr, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}

	defer client.Close()

	//生产者消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx_log"
	msg.Value = sarama.StringEncoder("this is a good test, my message is good")

	//发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed,", err)
		return
	}

	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
