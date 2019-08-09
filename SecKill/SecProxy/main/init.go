package main

import(
	"time"
	"encoding/json"

	"github.com/astaxie/beego/logs"
	"github.com/gomodule/redigo/redis"
	"Go-practice/SecKill/SecProxy/service"
)


var (
	redisPool *redis.Pool
)


func initRedis() (err error) {
	redisPool = &redis.Pool{
		// 空闲连接数
		MaxIdle: secKillConf.RedisConf.RedisMaxIdle,
		// 活跃连接数
		MaxActive: secKillConf.RedisConf.RedisMaxActive,
		// 超时 单位是纳秒的 Duration 类型
		IdleTimeout: time.Duration(secKillConf.RedisConf.RedisIdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", secKillConf.RedisConf.RedisAddr)
		},
	}
	// 获取一个链接 检测是否可用
	conn := redisPool.Get()
	defer conn.Close()
	_, err = conn.Do("ping")
	if err != nil {
		logs.Error("ping redis failed; err:%v", err)
		return 
	}
	logs.Info("redis init succ!")
	return
}


func initEtcd() (err error) {
	// TODO
	return
}


func convertLogLevel(level string) int {
	switch(level) {
		case "debug":
			return logs.LevelDebug
		case "warn":
			return logs.LevelWarn
		case "info":
			return logs.LevelInfo
		case "trace":
			return logs.LevelTrace
	}
	return logs.LevelDebug
}


func initLogger() (err error) {
	config := make(map[string]interface{})
	config["filename"] = secKillConf.LogPath
	config["level"] = convertLogLevel(secKillConf.LogLevel)

	configStr, err := json.Marshal(config)
	if err != nil {
		logs.Error("Marshal failed, err:%v", err)
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configStr))
	logs.Info("init logger succ")
	return
}


// 加载秒杀活动、商品配置 (应该通过 etcd 读取)
func loadSecConf() (err error) {
	// 先在 config.go 中通过读取配置文件配置 TODO
	return
}


func initSec() (err error) {
	// 初始化日志
	err = initLogger()
	if err != nil {
		logs.Error("init logger failed; err:%v", err)
		return 
	}

	// 初始化 redis
	err = initRedis()
	if err != nil {
		logs.Error("init redis failed, err:%v", err)
		return 
	}
	// 初始化 etcd
	err = initEtcd()
	if err != nil {
		logs.Error("init etcd failed, err:%v", err)
		return 
	}

	// 加载秒杀活动、商品等配置
	err = loadSecConf()
	if err != nil {
		logs.Error("load conf failed, err:%v", err)
		return 
	}

	service.InitService(&secKillConf)
	logs.Info("init sec succ")
	return
}