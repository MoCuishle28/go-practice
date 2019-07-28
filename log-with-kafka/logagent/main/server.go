package main


import(
	"github.com/astaxie/beego/logs"
	"Go-practice/log-with-kafka/logagent/tailf"
	"Go-practice/log-with-kafka/logagent/kafka"
	"time"
	// "fmt"
)


func serverRun() (err error) {
	for {
		msg := tailf.GetOneLine()
		err = sendToKafka(msg)
		if err != nil {
			logs.Error("send to kafka failed, err:%v", err)
			time.Sleep(1*time.Second)
			continue
		}
	}	
	return
}


func sendToKafka(msg *tailf.TextMsg) (err error) {
	// fmt.Printf("read msg:%s, topic:%s\n", msg.Msg, msg.Topic)
	err = kafka.SendToKafka(msg.Msg, msg.Topic)
	return
}