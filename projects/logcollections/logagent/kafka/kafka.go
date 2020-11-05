package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

var (
	client sarama.SyncProducer
)

func Init(addrs []string, err error) {
	config := sarama.NewConfig()
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("kafka producer closed,err=", err)
		return
	}
	return

}

func SendToKafka(topic, data string) {
	//构造消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message to kafka failed,err=", err)
		return
	}
	fmt.Printf("pid=%v offset=%v\n", pid, offset)
}
