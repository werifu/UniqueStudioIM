package model

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"sync"
	"thchat/pkg/logging"
	"time"
)

type SessionManager struct {
	cookieName	string
	mu 			sync.Mutex
	provider	Provider
	maxLifeTime int64
}



// 作为session的底层存储接口（比如硬盘、内存等）
type Provider interface {
	SessionInit(sid string)	(Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

//多种底层支持
var (
	providers = make(map[string]Provider)
)

type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) 		interface{}
	Delete(key interface{}) 	error
	SessionID() 				string
}

// 注册新的底层存储，只能注册一次
func Register(name string, provider Provider){
	if provider == nil {
		panic("provider为空")
	}
	if _, ok := providers[name]; ok {
		panic("provider注册两次：" + name)
	}
	providers[name] = provider
}

// 生成唯一sessionID
func (sm *SessionManager) GenSessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// 为每个请求判定session，有则返回原来的session并更新，没有则创建session
func (sm *SessionManager) SessionStart(c *gin.Context) Session {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	var session Session
	cookie, err := c.Request.Cookie(sm.cookieName)
	if err != nil || cookie.Value == ""{	// 没有cookie
		sid := sm.GenSessionID()
		session, _ = sm.provider.SessionInit(sid)
		cookie := http.Cookie{Name: sm.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(sm.maxLifeTime)}
		http.SetCookie(c.Writer, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = sm.provider.SessionRead(sid)
	}
	return session
}

// 销毁用户session（如登出等操作）
func (sm *SessionManager) SessionDestroy(c *gin.Context) {
	cookie, err := c.Request.Cookie(sm.cookieName)
	if err != nil || cookie.Value == "" {
		return
	}
	sm.mu.Lock()
	defer sm.mu.Unlock()

	err = sm.provider.SessionDestroy(cookie.Value)
	if err != nil {
		logging.Error(err)
		return
	}
	expiration := time.Now()
	cookie = &http.Cookie{Name: sm.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
	http.SetCookie(c.Writer, cookie)
}

//
func (sm *SessionManager) GC() {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.provider.SessionGC(sm.maxLifeTime)
	time.AfterFunc(time.Duration(sm.maxLifeTime), func(){ sm.GC() })
}

