package v1

import (
	"ginblog_backend/model"
	"ginblog_backend/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int


// 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreatUser(&data)
	}
	//if code == errmsg.ERROR_USERNAME_USED{
	//	code = errmsg.ERROR_USERNAME_USED
	//}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// 查询单个用户

// 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")

	switch  {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0{
		pageNum = 1
	}

	data, total := model.GetUsers(username,pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":code,
		"data":data,
		"total":total,
		"message":errmsg.GetErrMsg(code),
	})
}

// 编辑用户
func EditUser(c *gin.Context) {
	// todo
}

// 删除用户

func DeleteUser(c *gin.Context) {
	// todo

}
