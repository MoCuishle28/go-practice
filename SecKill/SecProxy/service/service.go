package service

import(
	"time"
	"fmt"
	"crypto/md5"

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


// 处理秒杀逻辑
func SecKill(req *SecRequest) (data []map[string]interface{}, code int, err error) {
	secKillConf.RwSecProductLock.RLock()
	defer secKillConf.RwSecProductLock.RUnlock()

	err = userCheck(req)
	if err != nil {
		code = ErrUserCheckAuthFailed
		logs.Warn("userid[%d] invalid, check failed, req:%+v", req.UserId, req)
		return 
	}

	// 判别是非为恶意访问
	err = antiSpam(req)
	if err != nil {
		code = ErrUserServiceBusy
		logs.Warn("userid[%d] busy, req:%+v", req.UserId, req)
	}
	return
}


// 校验密钥
func userCheck(req *SecRequest) (err error) {
	// 拼接
	authData := fmt.Sprintf("%d:%s", req.UserId, secKillConf.CookieSecretKey)
	// 返回 authData 的 MD5 校验和 ("%x"-> 十六进制格式化)
	authSign := fmt.Sprintf("%x", md5.Sum([]byte(authData)))
	if authSign != req.UserAuthSign {
		err = fmt.Errorf("invalid cookie user auth")
	}
	return
}


// 返回 data 是一个 map 的数组 (返回商品列表)
func SecInfoList() (data []map[string]interface{}, code int, err error) {
	secKillConf.RwSecProductLock.RLock()
	defer secKillConf.RwSecProductLock.RUnlock()

	for _, v := range secKillConf.SecProductInfoMap {
		item, _, err := SecInfoById(int(v.ProductId))
		if err != nil {
			logs.Error("get product_id[%d] failed, err:%v", v.ProductId, err)
			continue
		}
		data = append(data, item)
	}
	return
}


func SecInfo(productId int) (data []map[string]interface{}, code int, err error) {
	item, code, err := SecInfoById(productId)
	if err != nil {
		return 
	}
	data = append(data, item)
	return
}


func SecInfoById(productId int) (data map[string]interface{}, code int, err error) {
	// 加读锁 RLock() (Lock()是写锁)
	secKillConf.RwSecProductLock.RLock()
	defer secKillConf.RwSecProductLock.RUnlock()
	v, ok := secKillConf.SecProductInfoMap[int64(productId)]
	if !ok {
		code = ErrNotFoundProductId
		err = fmt.Errorf("not found productId:%d", productId)
		return
	}

	start := false
	end := false
	status := "success"
	// 防止客户端时间不一致 拿到 true 就同时开始
	// 以服务器时间为准 判断是否开始

	// 还没开始
	now := time.Now().Unix()
	if now - v.StartTime < 0 {
		start = false
		end = false
		status = "sec kill is not start"
	}

	// 已经开始
	if now - v.StartTime > 0 {
		start = true
	}

	if now - v.EndTime > 0 {
		start = false
		end = true
		status = "sec kill end"
	}

	if v.Status == ProductStatusSaleOut || v.Status == ProductStatusForceSaleOut {
		start = false
		end = true
		status = "product sale out"
	}

	data = make(map[string]interface{}, 1)
	data["product_id"] = productId
	data["start"] = start
	data["end"] = end
	data["status"] = status
	return
}