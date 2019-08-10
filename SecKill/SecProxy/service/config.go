package service

import(
	"sync"
	"time"
)


const(
	ProductStatusNormal = 0
	ProductStatusSaleOut = 1
	ProductStatusForceSaleOut = 2
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
	CookieSecretKey string
	UserSecAccessLimit int 								// 每秒访问限制
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


type SecRequest struct {
	ProductId int64
	Source string
	AuthCode string
	SecTime string
	Nance string
	UserId int64
	UserAuthSign string 	// 权限签名 用于验证是否真的登录 (是 userid 和密钥联合生成的一个字符串)
	AccessTime time.Time 	// 访问时间
}