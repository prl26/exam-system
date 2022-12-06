package response

import (
	"github.com/gin-gonic/gin"
	questionBankError "github.com/prl26/exam-system/server/model/questionBank/error"
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
	NOT_FIND         = 404
	ClientError      = 400
	ServiceError     = 500
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

func CheckHandle(c *gin.Context, err error) {
	Result(ClientError, nil, err.Error(), c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}
func ErrorHandle(c *gin.Context, err error) {
	switch err {
	case questionBankError.ScoreError:
		Result(ClientError, nil, err.Error(), c)
	case questionBankError.NotLanguageSupportError:
		Result(ClientError, nil, err.Error(), c)
	default:
		if e, ok := err.(questionBankError.CompileError); ok {
			Result(compilationError, nil, e.Error(), c)
		} else {
			Result(ServiceError, nil, err.Error(), c)
		}
	}
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func NotFind(c *gin.Context) {
	Result(NOT_FIND, nil, "无法找到该数据", c)
}
