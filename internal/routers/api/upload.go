package api

import (
	"fmt"
	"ginblog_backend/global"
	"ginblog_backend/internal/service"
	"ginblog_backend/pkg/app"
	"ginblog_backend/pkg/convert"
	"ginblog_backend/pkg/errcode"
	"ginblog_backend/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

// @Summary 上传文件
// @Produce json
// @Param file body string true "文件"
// @Param type body int false "文件类型"
// @Success 200 {object} nil "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /upload/file [post]
func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		errmsg := fmt.Sprintf("/upload/file [post] c.Request.FormFile err = %v", err)
		global.Logger.Error(c, errmsg)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errmsg))
		return
	}
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		errmsg := fmt.Sprintf("/upload/file [post] svc.UploadFile err = %v", err)
		global.Logger.Error(c, errmsg)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(errmsg))
		return
	}
	response.ToResponse(gin.H{
		"accessURL": fileInfo.AccessURL,
	})
}
