package main

import(
	"fmt"
	"strings"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"Go-practice/SecKill/SecProxy/service"
)


var (
	secKillConf = service.SecKillConf{}
)


func loadSecInfo() (err error) {

	productId_list := strings.Split(beego.AppConfig.String("product_id"), ",")
	startTime_list := strings.Split(beego.AppConfig.String("start_time"), ",")
	endTime_list := strings.Split(beego.AppConfig.String("end_time"), ",")
	status_list := strings.Split(beego.AppConfig.String("status"), ",")
	total_list := strings.Split(beego.AppConfig.String("total"), ",")
	left_list := strings.Split(beego.AppConfig.String("left"), ",")

	secKillConf.SecProductInfoMap = make(map[int64]*service.SecProductInfoConf, len(productId_list))
	// 初始化加不加锁其实都行 
	// 不过如果做成实时修改配置的 必须加锁 
	// 提升性能办法：先不加锁地读取变化数据到临时 map 再加锁地直接赋值 map 给 secProductInfoMap (这样数据量大时，主线程不会一直处于阻塞状态)
	secKillConf.RwSecProductLock.Lock()
	for i, id := range productId_list {
		id, _ := strconv.ParseInt(id, 10, 64)
		secKillConf.SecProductInfoMap[id] = &service.SecProductInfoConf{}
		secKillConf.SecProductInfoMap[id].ProductId, _ = strconv.ParseInt(productId_list[i], 10, 64)
		secKillConf.SecProductInfoMap[id].StartTime, _ = strconv.ParseInt(startTime_list[i], 10, 64)
		secKillConf.SecProductInfoMap[id].EndTime, _ = strconv.ParseInt(endTime_list[i], 10, 64)
		secKillConf.SecProductInfoMap[id].Status, _ = strconv.ParseInt(status_list[i], 10, 64)
		secKillConf.SecProductInfoMap[id].Total, _ = strconv.ParseInt(total_list[i], 10, 64)
		secKillConf.SecProductInfoMap[id].Left, _ = strconv.ParseInt(left_list[i], 10, 64)
	}
	secKillConf.RwSecProductLock.Unlock()

	fmt.Println("init secProductInfoMap:")
	for k, v := range secKillConf.SecProductInfoMap {
		fmt.Printf("%d:%+v\n", k, v)
	}
	return
}


// 初始化配置
func initConfig() (err error) {
	// 获取 app.config 中的配置项
	secKillConf.LogPath = beego.AppConfig.String("logs_path")
	secKillConf.LogLevel = beego.AppConfig.String("log_level")

	redisAddr := beego.AppConfig.String("redis_addr")
	etcdAddr := beego.AppConfig.String("etcd_addr")
	logs.Debug("read config:%v", redisAddr)
	logs.Debug("read config:%v", etcdAddr)

	secKillConf.RedisConf.RedisAddr = redisAddr
	secKillConf.EtcdAddr = etcdAddr

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

	secKillConf.RedisConf.RedisMaxIdle = redisMaxIdle
	secKillConf.RedisConf.RedisMaxActive = redisMaxActive
	secKillConf.RedisConf.RedisIdleTimeout = redisIdleTimeout

	err = loadSecInfo()
	if err != nil {
		err = fmt.Errorf("load secProductInfoMap failed, err:%v", err)
		return
	}
	return
}