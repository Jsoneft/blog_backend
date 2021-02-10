package middleware

import (
	"context"
	"ginblog_backend/global"
	"github.com/gin-gonic/gin"
)

func ContextTimeout() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, global.AppSetting.DefaultContextTimeout)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
