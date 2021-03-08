package lib

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"strings"
	"sync"
)

type LimiterCache struct {
	data sync.Map
}

var IpCache *LimiterCache

func init() {
	IpCache = &LimiterCache{}
}

// 获取客户端ip
func ClientIp(r *http.Request) string {
	ip := strings.TrimSpace(strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

func IpLimiter(cap,rate int64) func(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(handler gin.HandlerFunc) gin.HandlerFunc {
		return func(c *gin.Context) {
			//ip := ClientIp(c.Request)
			ip := c.Request.RemoteAddr // ip+port测试
			var limiter *Bucket
			if v,ok := IpCache.data.Load(ip) ; ok {
				limiter = v.(*Bucket)
			} else {
				// 内存存ip的限流桶，可以用redis代替
				limiter = NewBucket(cap,rate)
				IpCache.data.Store(ip,limiter)
			}

			if limiter.CanGetToken() {
				handler(c)
			} else {
				c.AbortWithStatusJSON(429,gin.H{"message":"too many requests"})
			}
		}
	}
}
