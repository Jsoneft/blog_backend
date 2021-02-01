package v1

import (
	"fmt"
	"ginblog_backend/global"
	"ginblog_backend/internal/service"
	"ginblog_backend/pkg/app"
	"ginblog_backend/pkg/convert"
	"ginblog_backend/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// @Summary 查看指定文章
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [get]
func (a Article) Get(c *gin.Context) {
	param := service.ArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	valid, errs := app.BindAndValid(c, &param)
	response := app.NewResponse(c)
	if !valid {
		errmsg := fmt.Sprintf("/api/v1/articles/{id} [get] app.BindAndValid err = %v", errs)
		global.Logger.Error(errmsg)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errmsg))
		return
	}
	svc := service.New(c.Request.Context())
	articles, err := svc.GetArticle(&param)
	if err != nil {
		errmsg := fmt.Sprintf("/api/v1/articles/{id} [get] svc.GetArticle err = %v", err)
		global.Logger.Error(errmsg)
		response.ToErrorResponse(errcode.ErrorGetArticleFail.WithDetails(errmsg))
		return
	}
	response.ToResponse(articles)
}

// @Summary 获取多个文章
// @Produce json
// @Param name query string false "文章名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [get]
func (a Article) List(c *gin.Context) {
	param := service.ArticleListRequest{}
	valid, errs := app.BindAndValid(c, &param)
	response := app.NewResponse(c)
	if !valid {
		errmsg := fmt.Sprintf("/api/v1/articles [get] app.BindAndValid err = %v", errs)
		global.Logger.Error(errmsg)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errmsg))
		return
	}
	svc := service.New(c.Request.Context())
	pager := app.Pager{
		Page:     app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}
	articles, totalRows, err := svc.GetArticleList(&param, pager)
	if err != nil {
		errmsg := fmt.Sprintf("/api/v1/articles [get] svc.GetArticleList err = %v", err)
		global.Logger.Error(errmsg)
		response.ToErrorResponse(errcode.ErrorGetArticlesFail.WithDetails(errmsg))
		return
	}
	pager.TotalRows = totalRows
	response.ToResponseList(articles, pager.TotalRows)
}

// @Summary 更新文章
// @Produce json
// @Param title body string true "文章名称" minlength(3) maxlength(100)
// @Param desc body string false "文章描述"
// @Param content body string ture "文章内容"
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param cover_image_url body string false "封面图片链接" minlength(3) maxlength(100)
// @Param modified_by body string false "修改者" minlength(3) maxlength(100)
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {
	param := service.UpdateArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	valid, errs := app.BindAndValid(c, &param)
	response := app.NewResponse(c)
	if !valid {
		errmsg := fmt.Sprintf("/api/v1/articles/{id} [put] app.BindAndValid err = %v", errs)
		global.Logger.Error(errmsg)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errmsg))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.UpdateArticle(&param)
	if err != nil {
		errmsg := fmt.Sprintf("/api/v1/articles/{id} [put] svc.UpdateArticle err = %v", err)
		global.Logger.Error(errmsg)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail.WithDetails(errmsg))
		return
	}
	response.ToResponse(gin.H{})
}

// @Summary 删除指定文章
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [delete]
func (a Article) Delete(c *gin.Context) {
	param := service.DeleteArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	valid, errs := app.BindAndValid(c, &param)
	response := app.NewResponse(c)
	if !valid {
		errmsg := fmt.Sprintf("/api/v1/articles/{id} [delete] app.BindAndValid err = %v", errs)
		global.Logger.Error(errmsg)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errmsg))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.DeleteArticle(&param)
	if err != nil {
		errmsg := fmt.Sprintf("/api/v1/articles/{id} [delete] svc.DeleteArticle err = %v", err)
		global.Logger.Error(errmsg)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail.WithDetails(errmsg))
		return
	}
	response.ToResponse(gin.H{})
}

// @Summary 创建文章
// @Produce json
// @Param tag_id body string true "标签ID"
// @Param title body string true "文章标题"
// @Param desc body string false "文章简述"
// @Param cover_image_url body string true "封面图片地址"
// @Param content body string true "文章内容"
// @Param created_by body int true "创建者"
// @Param state body int false "状态"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}
	valid, errs := app.BindAndValid(c, &param)
	response := app.NewResponse(c)
	if !valid {
		errmsg := fmt.Sprintf("/api/v1/articles [post] app.BindAndValid err = %v", errs)
		global.Logger.Error(errmsg)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errmsg))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CreateArticle(&param)
	if err != nil {
		errmsg := fmt.Sprintf("/api/v1/articles [post] svc.CreateArticle err = %v", err)
		global.Logger.Error(errmsg)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail.WithDetails(errmsg))
		return
	}
	response.ToResponse(gin.H{})
}
