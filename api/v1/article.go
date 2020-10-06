package v1

import (
	"ginblog_backend/model"
	"ginblog_backend/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code = model.CreateArt(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// todo 查询单个文章

// todo 查询文章列表

// 编辑文章
func EditArticle(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Param("id"))
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code = model.EditArticle(cid, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除文章

func DeleteArticle(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteArticle(cid)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
