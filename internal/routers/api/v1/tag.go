package v1

import (
	"ginblog_backend/pkg/app"
	"ginblog_backend/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {}
func (t Tag) List(c *gin.Context) {
	// test
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}
func (t Tag) Update(c *gin.Context) {}
func (t Tag) Delete(c *gin.Context) {}
func (t Tag) Create(c *gin.Context) {}
