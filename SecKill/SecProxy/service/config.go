package service

import(
	"sync"
)


type RedisConf struct {
	RedisAddr string
	RedisMaxIdle int
	RedisMaxActive int
	RedisIdleTimeout int
}


// 在内存中保存配置
type SecKillConf struct {
	RedisConf RedisConf
	EtcdAddr string
	LogPath string
	LogLevel string
	SecProductInfoMap map[int64]*SecProductInfoConf 	// 秒杀商品信息 key:id value:info conf
	RwSecProductLock sync.RWMutex						// product map 专用锁
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
