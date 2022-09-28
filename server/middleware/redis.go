package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/utils"
)

func QuestionBankAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		Id := utils.GetUserID(c)
		if utils.IsExistInRedis(Id) == true {
			c.Abort()
			return
		}
	}
}
