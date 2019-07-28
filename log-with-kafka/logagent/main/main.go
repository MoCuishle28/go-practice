package main

import(
	"fmt"
	"github.com/astaxie/beego/logs"
	"Go-practice/log-with-kafka/logagent/tailf"
	"Go-practice/log-with-kafka/logagent/kafka"
)


func main() {
	// 加载配置文件
	filename := "../conf/logagent.conf"
	err := loadConf("ini", filename)
	if err != nil {
		fmt.Println("load conf failed, err:%v", err)
		panic("load conf failed")
		return
	}

	// 初始化日志组件
	err = initLogger()
	if err != nil {
		fmt.Println("load logger failed, err:%v", err)
		panic("load logger failed")
		return
	}


	err = tailf.InitTail(appConfig.collectConf, appConfig.chanSize)
	if err != nil {
		logs.Error("init tail failed, err:%v", err)
		return
	}

	err = kafka.InitKafka(appConfig.kafkaAddr)
	if err != nil {
		logs.Error("init kafka failed, err:%v", err)
		return
	}

	logs.Debug("init success!") //成功加载配置并能写入日志

	err = serverRun()	// 开始正式的逻辑处理
	if err != nil {
		logs.Error("serverRun failed, err:%v", err)
		return
	}

	logs.Info("program exited!")
}