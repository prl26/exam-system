package examManage

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/examManage/request"
	"github.com/prl26/exam-system/server/service/frontDesk/frontExam"
	"go.uber.org/zap"
)

type ExamApi struct {
}

var examService frontExam.ExamService

func (examApi *ExamApi) FindExamPlans(c *gin.Context) {
	var teachClassId uint
	_ = c.ShouldBindQuery(&teachClassId)
	if examPlans, err := examService.FindExamPlans(teachClassId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"examPlans": examPlans}, c)
	}
}
func (examApi *ExamApi) GetExamPaper(c *gin.Context) {
	var examComing request.ExamComing
	_ = c.ShouldBindJSON(&examComing)
}
