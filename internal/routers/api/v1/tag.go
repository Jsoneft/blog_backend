package v1

import (
	"fmt"
	"ginblog_backend/global"
	_ "ginblog_backend/internal/model"
	"ginblog_backend/internal/service"
	"ginblog_backend/pkg/app"
	"ginblog_backend/pkg/errcode"
	_ "ginblog_backend/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// @Summary 获取多个标签
// @Produce json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	// Bind 一定要传指针
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errmsg := fmt.Sprintf(" /api/v1/tags [get]  app.BindAndValid err = %v", errs)
		global.Logger.Error(errmsg)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errmsg))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{
		Page:     app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}
	totalRows, err := svc.CountTag(&service.CountTagRequest{
		Name:  param.Name,
		State: param.State,
	})
	if err != nil {
		errmsg := fmt.Sprintf(" /api/v1/tags [get]  svc.CountTag err = %v", err)
		global.Logger.Error(errmsg)
		response.ToErrorResponse(errcode.ErrorCountTagFail.WithDetails(errmsg))
		return
	}
	pager.TotalRows = totalRows
	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		errmsg := fmt.Sprintf(" /api/v1/tags [get]  svc.GetTagList err = %v", err)
		global.Logger.Error(errmsg)
		response.ToErrorResponse(errcode.ErrorGetTagListFail.WithDetails(errmsg))
		return
	}
	response.ToResponseList(tags, pager.TotalRows)
}

// @Summary 更新标签
// @Produce json
// @Param id path int true "标签ID"
// @Param name body string ture "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string false "修改者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{}
	valid, errs := app.BindAndValid(c, &param)
	response := app.NewResponse(c)
	if !valid {
		errmsg := fmt.Sprintf(" /api/v1/tags/{id} [put]  app.BindAndValid err = %v", errs)
		global.Logger.Error(errmsg)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errmsg))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		errmsg := fmt.Sprintf(" /api/v1/tags/{id} [put]  svc.UpdateTag err = %v", err)
		global.Logger.Error(errmsg)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail.WithDetails(errmsg))
		return
	}
	response.ToResponse(gin.H{})
}

// @Summary 删除标签
// @Produce json
// @Param id path int true "标签ID"
// @Success 200 {object} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	// 参数校验
	param := service.DeleteTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errmsg := fmt.Sprintf("/api/v1/tags/{id} [delete] app.BindAndValid errs = %v", errs)
		global.Logger.Error(errmsg)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errmsg))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		errmsg := fmt.Sprintf("/api/v1/tags/{id} [delete] svc.DeleteTag() errs = %v", err)
		global.Logger.Error(errmsg)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail.WithDetails(errmsg))
		return
	}
	response.ToResponse(gin.H{})
}

// @Summary 新增标签
// @Produce json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string false "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	// 参数校验
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errmsg := fmt.Sprintf("/api/v1/tags [post] app.BindAndValid errs = %v", errs)
		global.Logger.Error(errmsg)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errmsg))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		errmsg := fmt.Sprintf("/api/v1/tags [post] svc.CreateTag err = %v", err)
		global.Logger.Error(errmsg)
		response.ToErrorResponse(errcode.ErrorCreateTagFail.WithDetails(errmsg))
		return
	}
	response.ToResponse(gin.H{})
}
