package service


import(
	"fmt"
	"sync"
)

var(
	secLimitMgr = &SecLimitMgr{
		UserLimitMap: make(map[int64]*SecLimit, 10000),
	}
)


type SecLimitMgr struct {
	UserLimitMap map[int64]*SecLimit
	lock sync.Mutex
}


// 控制用户每秒访问
type SecLimit struct {
	count int
	currTime int64 	// 精确到秒的时间戳
}


// 判别用户是否为恶意访问
func antiSpam(req *SecRequest) (err error) {
	secLimitMgr.lock.Lock()
	secLimit, ok := secLimitMgr.UserLimitMap[req.UserId]
	if !ok {
		secLimit = &SecLimit{}
		secLimitMgr.UserLimitMap[req.UserId] = secLimit
	}
	count := secLimit.Count(req.AccessTime.Unix())	// 先计数
	secLimitMgr.lock.Unlock()

	if count > secKillConf.UserSecAccessLimit {
		err = fmt.Errorf("invalid request")
		return
	}
	return
}


func (p *SecLimit) Count(nowTime int64) int {
	// 和当前时间不为同一秒 说明当前秒访问次数为 1
	if p.currTime != nowTime {
		p.count = 1
		p.currTime = nowTime
	} else {
		// 否则是同一秒
		p.count++
	}
	return p.count
}


// 返回当前时间计数
func (p *SecLimit) Check(nowTime int64) int {
	if p.currTime != nowTime {
		return 0
	}
	return p.count
}