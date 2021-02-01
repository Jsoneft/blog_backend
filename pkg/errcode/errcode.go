package errcode

import (
	"fmt"
	"net/http"
)

var (
	Success                   = NewError(0, "Success!")
	ServerError               = NewError(10000000, "服务器内部错误")
	InvalidParams             = NewError(10000001, "入参错误")
	NotFound                  = NewError(10000002, "找不到")
	UnauthorizedAuthNotExist  = NewError(10000003, "鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "鉴权失败，Token错误")
	UnauthorizedTokenTimeout  = NewError(10000005, "鉴权失败，Token超时")
	UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败，Token生成失败")

	TooManyRequests = NewError(10000007, " 请求过多")

	// Tag类
	ErrorGetTagListFail = NewError(20000000, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(20000001, "创建标签失败")
	ErrorDeleteTagFail  = NewError(20000002, "删除标签失败")
	ErrorUpdateTagFail  = NewError(20000003, "更新标签失败")
	ErrorCountTagFail   = NewError(20000004, "统计标签失败")

	// Article类
	ErrorGetArticlesFail   = NewError(30000000, "获取文章列表失败")
	ErrorGetArticleFail    = NewError(30000001, "获取该文章失败")
	ErrorCreateArticleFail = NewError(30000002, "创建文章失败")
	ErrorDeleteArticleFail = NewError(30000003, "删除文章失败")
	ErrorUpdateArticleFail = NewError(30000004, "更新文章失败")
)

type Error struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %v 已经存在请换一个", code))
	}
	codes[code] = msg
	return &Error{
		Code:    code,
		Msg:     msg,
		Details: nil,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码: %d, 错误信息: %s", e.Code, e.Msg)
}

//func (e *Error) Code() int {
//	return e.Code
//}
//
//func (e *Error) Msg() string {
//	return e.Msg
//}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.Msg, args...)
}

//func (e *Error) Details() []string {
//	return e.Details
//}
func (e *Error) WithDetails(details ...string) *Error {
	e.Details = []string{}
	for _, v := range details {
		e.Details = append(e.Details, v)
	}
	return e
}

func (e *Error) StatusCode() int {
	switch e.Code {
	case Success.Code:
		return http.StatusOK
	case ServerError.Code:
		return http.StatusInternalServerError
	case InvalidParams.Code:
		return http.StatusBadRequest
	case NotFound.Code:
		return http.StatusNotFound
	case UnauthorizedAuthNotExist.Code:
		fallthrough
	case UnauthorizedTokenError.Code:
		fallthrough
	case UnauthorizedTokenGenerate.Code:
		fallthrough
	case UnauthorizedTokenTimeout.Code:
		return http.StatusUnauthorized
	case TooManyRequests.Code:
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
