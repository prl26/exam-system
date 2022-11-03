package common

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/service/frontDesk"
	"go.uber.org/zap"
)

type CommonApi struct {
}

var commonService = frontDesk.CommonService{}

//获取该学生所在的每个教学班
func (commonApi *CommonApi) FindLessons(c *gin.Context) {
	var studentId request.GetByStudentId
	_ = c.ShouldBindQuery(&studentId)
	if teachClassAndLesson, err := commonService.FindTeachClass(studentId.StudentId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"nameOfLessons": teachClassAndLesson}, c)
	}
}
