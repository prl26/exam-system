package examManage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/examManage"
	"github.com/prl26/exam-system/server/model/examManage/request"
	questionBankResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
)

type DraftPaperApi struct {
}

var DraftPaperService = service.ServiceGroupApp.ExammanageServiceGroup.DraftPaperService

func (draftPaperApi *DraftPaperApi) CreatePaperDraft(c *gin.Context) {
	var examPaper examManage.ExamPaperDraft
	_ = c.ShouldBindJSON(&examPaper)
	verify := utils.Rules{
		"Name":     {utils.NotEmpty()},
		"LessonId": {utils.NotEmpty()},
	}
	if err := utils.Verify(examPaper, verify); err != nil {
		response.CheckHandle(c, err)
		return
	} else {
		userId := utils.GetUserID(c)
		examPaper.UserId = &userId
		_ = c.ShouldBindJSON(&examPaper)
		if err := DraftPaperService.CreateExamPaperDraft(examPaper); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("试卷创建失败", c)
		} else {
			response.OkWithData(gin.H{
				"status": "创建成功",
			}, c)
		}
	}
}
func (draftPaperApi *DraftPaperApi) DeleteExamPaperDraft(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := DraftPaperService.DeleteExamPaperDraft(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("批量删除失败:%s", err.Error()))
	} else {
		questionBankResp.OkWithMessage("批量删除成功", c)
	}
}
func (draftPaperApi *DraftPaperApi) UpdateExamPaperDraft(c *gin.Context) {
	var examPaper examManage.ExamPaperDraft
	_ = c.ShouldBindJSON(&examPaper)
	if err := DraftPaperService.UpdateExamPaperDraft(examPaper); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
func (draftPaperApi *DraftPaperApi) FindExamPaperDraft(c *gin.Context) {
	var examPaper examManage.ExamPaperDraft
	_ = c.ShouldBindQuery(&examPaper)
	if reexamPaper, err := DraftPaperService.GetExamPaperDraft(examPaper.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{
			"reexamPaper": reexamPaper,
		}, c)
	}
}
func (draftPaperApi *DraftPaperApi) GetExamPaperDraftList(c *gin.Context) {
	var pageInfo request.DraftPaperSearch
	_ = c.ShouldBindQuery(&pageInfo)
	userId := utils.GetUserID(c)
	authorityId := utils.GetUserAuthorityID(c)
	fmt.Println(pageInfo)
	if list, total, err := DraftPaperService.GetPaperDraftInfoList(pageInfo, userId, authorityId); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
func (draftPaperApi *DraftPaperApi) ConvertDraftToPaper(c *gin.Context) {
	var info request.ConvertDraft
	_ = c.ShouldBindJSON(&info)
	userId := utils.GetUserID(c)
	if IsOk, err := DraftPaperService.ConvertDraftCheck(info); err != nil {
		response.FailWithMessage("统计分数出错了", c)
	} else if IsOk == false {
		response.FailWithMessage("试卷总分应该为100分", c)
	} else {
		if paperId, err := DraftPaperService.ConvertDraftToPaper(info, userId); err != nil {
			response.FailWithMessage("通过草稿生成试卷失败", c)
		} else {
			response.OkWithData(gin.H{
				"status":  "创建成功",
				"paperId": paperId,
			}, c)
		}
	}
}
