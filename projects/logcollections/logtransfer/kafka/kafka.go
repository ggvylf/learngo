package kafka

import (
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/ggvylf/learngo/projects/logcollections/logtransfer/es"
)

type LogData struct {
	data string `json:"data"`
}

func Init(addrs []string, topic string) (err error) {
	//初始化consumer
	consumer, err := sarama.NewConsumer(addrs, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
			//发送给es
			ld:=LogData{
				data:string(msg.Value)
			}
			//把消息的内容写入ld中
			err = json.Unmarshal(mgs.Value, ld)
			if err != nil {
				fmt.Println("unmarahal failed ,err=", err)
				continue
			}
			es.SendToES(topic, ld)

		}(pc)
	}
}
