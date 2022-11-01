package questionBank

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/questionBank"
	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/request"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/30 22:02

 * @Note:

 **/

type CommonApi struct {
}

var questionBankService = service.ServiceGroupApp.QuestionBankServiceGroup.QuestionBankService

// AddCourseSupport 增加题目的课程支持
func (CommonApi) AddCourseSupport(c *gin.Context) {
	var req questionBankReq.CourseSupportAdd
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"ChapterId":    {utils.NotEmpty()},
		"QuestionId":   {utils.NotEmpty()},
		"QuestionType": {utils.NotEmpty()},
	}
	n := len(req.CourseSupports)
	if n == 0 {
		merges := make([]questionBank.ChapterMerge, n)
		for i, support := range req.CourseSupports {
			if err := utils.Verify(support, verify); err != nil {
				response.FailWithMessage(err.Error(), c)
				return
			}
			merges[i].QuestionId = support.QuestionId
			merges[i].QuestionType = support.QuestionType
			merges[i].ChapterId = support.ChapterId
		}
		if err := questionBankService.AddCourseSupport(merges); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
			return
		}
	}
	response.Ok(c)
}

// DeleteCourseSupport 删除题目的课程支持
func (CommonApi) DeleteCourseSupport(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := questionBankService.DeleteCourseSupport(IDS.Ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// FindQuestionSupport 分页获取题目支持
func (CommonApi) FindQuestionSupport(c *gin.Context) {
	var req questionBankReq.QuestionsSupportFind
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"ChapterId":    {utils.NotEmpty()},
		"QuestionType": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := questionBankService.FindQuestionSupport(req); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", c)
	}
}
