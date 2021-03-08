package lib

import "github.com/gin-gonic/gin"

// 针对url形式为"abs?name=xxx"进行限流，name就是key参数
func ParamLimiter(cap,rate int64,key string) func(handler gin.HandlerFunc) gin.HandlerFunc {
	limiter := NewBucket(cap,rate)
	return func(handler gin.HandlerFunc) gin.HandlerFunc {
		return func(c *gin.Context) {
			if c.Query(key) != "" {
				if limiter.CanGetToken() {
					handler(c)
				}else {
					c.AbortWithStatusJSON(429,gin.H{"message":"too many requests-param"})
				}
			} else {
				handler(c)
			}
		}
	}
}