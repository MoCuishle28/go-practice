package main

import(
	"fmt"
	"strings"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)


var (
	secKillConf = SecKillConf{}
)


type RedisConf struct {
	redisAddr string
	redisMaxIdle int
	redisMaxActive int
	redisIdleTimeout int
}


// 在内存中保存配置
type SecKillConf struct {
	redisConf RedisConf
	etcdAddr string
	logPath string
	logLevel string
	secProductInfo []SecProductInfoConf 	// 秒杀商品信息
}


// json 反序列化必须首字母大写
type SecProductInfoConf struct {
	ProductId int64 	 	// 商品ID
	StartTime int64 		// 抢购开始时间
	EndTime int64
	Status int64 			// 抢购状态
	Total int64 			// 总量
	Left int64 				// 剩余量
}


func loadSecInfo() (err error) {

	productId_list := strings.Split(beego.AppConfig.String("product_id"), ",")
	startTime_list := strings.Split(beego.AppConfig.String("start_time"), ",")
	endTime_list := strings.Split(beego.AppConfig.String("end_time"), ",")
	status_list := strings.Split(beego.AppConfig.String("status"), ",")
	total_list := strings.Split(beego.AppConfig.String("total"), ",")
	left_list := strings.Split(beego.AppConfig.String("left"), ",")

	secKillConf.secProductInfo = make([]SecProductInfoConf, len(productId_list))
	for i, _ := range secKillConf.secProductInfo {
		secKillConf.secProductInfo[i].ProductId, _ = strconv.ParseInt(productId_list[i], 10, 64)
		secKillConf.secProductInfo[i].StartTime, _ = strconv.ParseInt(startTime_list[i], 10, 64)
		secKillConf.secProductInfo[i].EndTime, _ = strconv.ParseInt(endTime_list[i], 10, 64)
		secKillConf.secProductInfo[i].Status, _ = strconv.ParseInt(status_list[i], 10, 64)
		secKillConf.secProductInfo[i].Total, _ = strconv.ParseInt(total_list[i], 10, 64)
		secKillConf.secProductInfo[i].Left, _ = strconv.ParseInt(left_list[i], 10, 64)
	}

	for _, v := range secKillConf.secProductInfo {
		fmt.Println("%v", v)
	}
	return
}


// 初始化配置
func initConfig() (err error) {
	// 获取 app.config 中的配置项
	secKillConf.logPath = beego.AppConfig.String("logs_path")
	secKillConf.logLevel = beego.AppConfig.String("log_level")

	redisAddr := beego.AppConfig.String("redis_addr")
	etcdAddr := beego.AppConfig.String("etcd_addr")
	logs.Debug("read config:%v", redisAddr)
	logs.Debug("read config:%v", etcdAddr)

	secKillConf.redisConf.redisAddr = redisAddr
	secKillConf.etcdAddr = etcdAddr

	if len(redisAddr) == 0 || len(etcdAddr) == 0 {
		err = fmt.Errorf("init config failed, redis[%s] or etcd[%s] config is null", redisAddr, etcdAddr)
		return 
	}

	redisMaxIdle, err := beego.AppConfig.Int("redis_max_idle")
	if err != nil {
		err = fmt.Errorf("init redis config failed, redis_max_idle:%v; err:%v", redisMaxIdle, err)
		return 
	}

	redisMaxActive, err := beego.AppConfig.Int("redis_max_active")
	if err != nil {
		err = fmt.Errorf("init redis config failed, redis_max_active:%v; err:%v", redisMaxActive, err)
		return 
	}

	redisIdleTimeout, err := beego.AppConfig.Int("redis_idle_timeout")
	if err != nil {
		err = fmt.Errorf("init redis config failed, redis_max_timeout:%v; err:%v", redisIdleTimeout, err)
		return 
	}

	secKillConf.redisConf.redisMaxIdle = redisMaxIdle
	secKillConf.redisConf.redisMaxActive = redisMaxActive
	secKillConf.redisConf.redisIdleTimeout = redisIdleTimeout

	err = loadSecInfo()
	if err != nil {
		err = fmt.Errorf("load secProductInfo failed, err:%v", err)
		return
	}
	return
}