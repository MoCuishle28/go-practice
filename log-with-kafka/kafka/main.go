package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// kafka生产者例子

func main() {
	// Kafka生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机分区 （分区的负载均衡）
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	// 生存者对象
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}

	defer client.Close()

	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx_log"
	msg.Value = sarama.StringEncoder("this is a good test, my message is good")

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed,", err)
		return
	}

	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
