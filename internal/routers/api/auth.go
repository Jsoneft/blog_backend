package api

import (
	"fmt"
	"ginblog_backend/global"
	"ginblog_backend/internal/service"
	"ginblog_backend/pkg/app"
	"ginblog_backend/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// @Summary JWT认证
// @Produce json
// @Param app_key path string true "App唯一标识符"
// @Param app_secret path string true "App密钥"
// @Success 200 {object} nil "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	param := &service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errmsg := fmt.Sprintf("/auth [get] c.Request.FormFile err = %v", errs)
		global.Logger.Error(c, errmsg)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errmsg))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(param)
	if err != nil {
		errmsg := fmt.Sprintf("/auth [get] svc.CheckAuth err = %v", err)
		global.Logger.Error(c, errmsg)
		response.ToErrorResponse(errcode.UnauthorizedTokenError.WithDetails(errmsg))
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		errmsg := fmt.Sprintf("/auth [get] app.GenerateToken err = %v", err)
		global.Logger.Error(c, errmsg)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate.WithDetails(errmsg))
		return
	}
	response.ToResponse(gin.H{
		"token": token,
	})
}
