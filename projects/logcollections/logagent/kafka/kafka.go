package kafka

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

var (
	client      sarama.SyncProducer
	logDataChan chan *logData
)

type logData struct {
	topic string
	data  string
}

func Init(addrs []string, maxSize int) (err error) {
	//初始化client
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("kafka producer closed,err=", err)
		return
	}

	//初始化通道
	logDataChan := make(chan *logData, maxSize)

	//取出数据发送到kafka
	go sendToKafka()
	return
}

func sendToKafka() {
	//构造消息
	for {

		select {
		case ld := <-loadDataChan:
			msg := &sarama.ProducerMessage{}
			msg := ld.topic
			msg.Value = sarama.StringEncoder(ld.data)

			//发送消息
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send message to kafka failed,err=", err)
				return
			}
			fmt.Printf("pid=%v offset=%v\n", pid, offset)
		default:
			time.Sleep(500 * time.Millisecond)
		}

	}

}

func SendToChan(topic, data string) {
	msg = &logData{
		topic: topic,
		data:  data,
	}

	//把消息写入chan中
	logDataChan <- msg

}
