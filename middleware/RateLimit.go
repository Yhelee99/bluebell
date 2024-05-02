package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {

	//新建一个令牌桶
	rl := ratelimit.NewBucket(fillInterval, cap) //指定放令牌的次数及令牌桶容量
	return func(c *gin.Context) {
		if rl.TakeAvailable(cap) == 0 { //表示从令牌桶中取cap个令牌，TakeAvailable如果取失败返回0，且该方法不会阻塞
			//取不到令牌就返回阻塞的响应
			c.JSON(http.StatusOK, gin.H{
				"msg": "rate limit!",
			})
			c.Abort()
		}
		//取到令牌就放行
		c.Next()
	}
}
