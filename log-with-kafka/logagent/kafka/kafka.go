package kafka


import(
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var (
	client sarama.SyncProducer
)


func InitKafka(addr string) (err error) {
	// Kafka生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机分区 （分区的负载均衡）
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	// 生存者对象
	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		fmt.Println("init kafka producer failed, err:", err)
		logs.Error("init kafka producer failed, err:", err)
		return
	}
	logs.Debug("init kafka succ")
	return
}


func SendToKafka(data, topic string) (err error) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed, err:%v data:%v topic:%v\n", err, data, topic)
		logs.Error("send message failed, err:%v data:%v topic:%v", err, data, topic)
		return
	}

	fmt.Printf("send succ, pid:%v offset:%v topic:%v data:%v\n", pid, offset, topic, data)
	logs.Debug("send succ, pid:%v offset:%v topic:%v data:%v", pid, offset, topic, data)
	return
}