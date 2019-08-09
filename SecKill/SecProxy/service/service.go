package service

import(
	"fmt"

	"github.com/astaxie/beego/logs"
)

// service 写具体逻辑

var(
	secKillConf *SecKillConf
)


func InitService(serviceConf *SecKillConf) {
	secKillConf = serviceConf
	logs.Debug("init service succ, config:%+v", secKillConf)
}


func SecInfo(productId int) (data map[string]interface{}, code int, err error) {
	// 加读锁 RLock() (Lock()是写锁)
	secKillConf.RwSecProductLock.RLock()
	defer secKillConf.RwSecProductLock.RUnlock()
	v, ok := secKillConf.SecProductInfoMap[int64(productId)]
	if !ok {
		code = ErrNotFoundProductId
		err = fmt.Errorf("not found productId:%d", productId)
		return
	}

	data = make(map[string]interface{})
	data["product_id"] = productId
	data["start_time"] = v.StartTime
	data["end_time"] = v.EndTime
	data["status"] = v.Status
	return
}