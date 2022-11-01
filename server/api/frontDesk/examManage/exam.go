package examManage

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/service/examManage"
	"go.uber.org/zap"
)

type ExamApi struct {
}

var examService = examManage.ExamService{}

func (examApi *ExamApi) FindLessons(c *gin.Context) {
	var studentId uint
	_ = c.ShouldBindQuery(&studentId)
	if nameOfLessons, err := examService.FindTeachClass(studentId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"nameOfLessons": nameOfLessons}, c)
	}
}

//func (examApi *ExamApi) FindLessons(c *gin.Context) {
//
//}
