package response

import (
	"github.com/gin-gonic/gin"
	exception "github.com/prl26/exam-system/server/model/oj/error"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR            = 7
	SUCCESS          = 0
	compilationError = 10000
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func ErrorHandle(c *gin.Context, err error) {
	switch err.(type) {
	case exception.CompileError:
		Result(compilationError, nil, err.Error(), c)
	default:
		Result(ERROR, nil, err.Error(), c)
	}
}
