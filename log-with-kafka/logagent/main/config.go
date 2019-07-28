package main


import(
	"fmt"
	"errors"
	"github.com/astaxie/beego/config"
	"Go-practice/log-with-kafka/logagent/tailf"
)

var (
	appConfig *Config
)

type Config struct {
	logLevel string
	logPath string

	chanSize int
	kafkaAddr string
	collectConf []tailf.CollectConf
}




func loadCollectConf(conf config.Configer) (err error) {
	var cc tailf.CollectConf
	cc.LogPath = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid collect::log_path")
		return err
	}

	cc.Topic = conf.String("collect::topic")
	if len(cc.Topic) == 0 {
		err = errors.New("invalid collect::topic")
		return err
	}

	appConfig.collectConf = append(appConfig.collectConf, cc)
	return
}


func loadConf(confType, filename string) (err error) {

	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return err
	}

	appConfig = &Config{}
	appConfig.logLevel = conf.String("log::log_level")
	if len(appConfig.logLevel) == 0 {
		appConfig.logLevel = "debug"
	}

	appConfig.logPath = conf.String("log::log_path")
	if len(appConfig.logPath) == 0 {
		appConfig.logPath = "./logs"
	}

	appConfig.chanSize, err = conf.Int("collect::chan_size")
	if err != nil {
		appConfig.chanSize = 100
	}

	appConfig.kafkaAddr = conf.String("kafka::server_addr")
	if len(appConfig.kafkaAddr) == 0 {
		err = fmt.Errorf("invalid kafka addr")
		return
	}

	err = loadCollectConf(conf)
	if err != nil {
		fmt.Println("load collect conf failed, err:%v", err)
		return err
	}

	return 
}