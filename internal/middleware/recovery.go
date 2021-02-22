package middleware

import (
	"fmt"
	"ginblog_backend/global"
	"ginblog_backend/pkg/app"
	"ginblog_backend/pkg/email"
	"ginblog_backend/pkg/errcode"
	"github.com/gin-gonic/gin"
	"time"
)

func Recovery() gin.HandlerFunc {

	defaultMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.PassWord,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrame().Errorf(c, "panic recover err: %v", err)
				err := defaultMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("来自ZJX的Panic提醒服务 于:%v 抛出异常", time.Now()),
					fmt.Sprintf("错误信息 err: %v", err),
				)
				if err != nil {
					global.Logger.Panicf(c, "defaultMailer.SendMail err : %v", err)
				}
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		// panic Mail test
		//log.Panicf("panic err = %v", "test")
		c.Next()
	}
}
