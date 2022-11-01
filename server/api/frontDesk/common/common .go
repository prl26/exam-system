package common

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/service/frontDesk"
	"go.uber.org/zap"
)

type CommonApi struct {
}

var commonService = frontDesk.CommonService{}

func (commonApi *CommonApi) FindLessons(c *gin.Context) {
	var studentId uint
	_ = c.ShouldBindQuery(&studentId)
	if teachClassAndLesson, err := commonService.FindTeachClass(studentId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"nameOfLessons": teachClassAndLesson}, c)
	}
}
