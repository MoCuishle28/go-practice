package mySession

import (
	"fmt"
	"sync"
	"encoding/base64"
	"math/rand"
)


// 抽象出一个Provider接口，用以表征session管理器底层存储结构
type Provider interface {
	SessionInit(sid string) (Session, error)	// Session的初始化
	SessionRead(sid string) (Session, error)	// 返回sid所代表的Session变量 不存在则以 sid 创建新的 session
	SessionDestroy(sid string) error 			// 销毁sid对应的Session变量
	SessionGC(maxLifeTime int64)				// 根据maxLifeTime来删除过期的数据
}


type Session interface {
	Set(key, value interface{}) error // set session value
	Get(key interface{}) interface{}  // get session value
	Delete(key interface{}) error     // delete session value
	SessionID() string                // back current sessionID
}


// 定义一个全局的session管理器
type Manager struct {
	cookieName  string     // private cookiename
	lock        sync.Mutex // protects session
	provider    Provider
	maxLifeTime int64
}


func NewManager(provideName, cookieName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxLifeTime: maxLifeTime}, nil
}


var provides = make(map[string]Provider)

// Register makes a session provide available by the provided name.
// If Register is called twice with the same name or if driver is nil, it panics.
func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provider " + name)
	}
	provides[name] = provider
}


// 生成全局唯一ID
func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}