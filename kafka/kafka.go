package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"llvvlv00.org/logtransfer/es"
)

// 初始化消费者，从kafka取数据发往es
func Init(addrs []string, topic string)(err error) {
	consumer, err := sarama.NewConsumer(addrs, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err: %v\n", err)
		return
	}

	partitionList, err := consumer.Partitions(topic) //根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition: err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList {	//遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d, err: %v\n", partition, err)
			return err
		}
		defer pc.AsyncClose()

		//异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
				//直接发给es
				ld := es.LogData{
					Topic:topic,
					Data:string(msg.Value),
				}
				es.SendToESChan(&ld)	//优化，放到一个channel中
			}
		}(pc)
	}
	return
}
