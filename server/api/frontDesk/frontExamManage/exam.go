package frontExamManage

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	request2 "github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/examManage"
	"github.com/prl26/exam-system/server/model/examManage/request"
	"github.com/prl26/exam-system/server/service/frontDesk/frontExam"
	"go.uber.org/zap"
)

type ExamApi struct {
}

var examService frontExam.ExamService

// 查询该学生 某个教学班 下所有的教学计划
func (examApi *ExamApi) FindExamPlans(c *gin.Context) {
	var teachClassId request2.GetByTeachClassId
	_ = c.ShouldBindQuery(&teachClassId)
	if examPlans, err := examService.FindExamPlans(teachClassId.TeachClassId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"examPlans": examPlans}, c)
	}
}

// 学生进入考试时获取试卷内容
func (examApi *ExamApi) GetExamPaper(c *gin.Context) {
	var examComing request.ExamComing
	_ = c.ShouldBindJSON(&examComing)
	if examPaper, err := examService.GetExamPapers(examComing); err != nil {
		global.GVA_LOG.Error("查询考试试卷失败", zap.Error(err))
		response.FailWithMessage("查询考试试卷失败", c)
	} else {
		response.OkWithData(gin.H{"examPaper": examPaper}, c)
	}
}

//提交试卷内容
func (examApi *ExamApi) CommitExamPaper(c *gin.Context) {
	var ExamCommit examManage.CommitExamPaper
	_ = c.ShouldBindJSON(&ExamCommit)
	if err := examService.CommitExamPapers(ExamCommit); err != nil {
		global.GVA_LOG.Error("试卷提交失败", zap.Error(err))
		response.FailWithMessage("试卷提交试卷失败", c)
	} else {
		response.OkWithData(gin.H{"examPaper": ExamCommit}, c)

	}
}
