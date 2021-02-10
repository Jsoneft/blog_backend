package middleware

import (
	"ginblog_backend/pkg/app"
	"ginblog_backend/pkg/errcode"
	"ginblog_backend/pkg/limiter"
	"github.com/gin-gonic/gin"
)

func RateLimiter(l limiter.LimiterIF) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				resp := app.NewResponse(c)
				resp.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}
	}
}
