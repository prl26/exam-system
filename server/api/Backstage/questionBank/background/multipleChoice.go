package questionBank

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/enum/questionType"
	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/request"
	questionBankResp "github.com/prl26/exam-system/server/model/questionBank/response"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
)

type MultipleChoiceApi struct {
}

var multipleChoiceService = service.ServiceGroupApp.QuestionBankServiceGroup.MultipleChoiceService

// Create 创建选择题
func (choiceApi *MultipleChoiceApi) Create(c *gin.Context) {
	var req questionBankReq.MultipleChoiceCreate
	_ = c.ShouldBindJSON(&req)
	req.ID = 0
	verify := utils.Rules{
		"ProblemType": {utils.NotEmpty()},
		"CanPractice": {utils.NotEmpty()},
		"CanExam":     {utils.NotEmpty()},
		"Title":       {utils.NotEmpty()},
		"Describe":    {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if req.MostOptions > len(req.Options) {
		response.FailWithMessage("可选择项数大于总的选项个数", c)
	}
	if len(req.Options) != 0 {
		verify := utils.Rules{
			"Describe": {utils.NotEmpty()},
			"Orders":   {utils.NotEmpty()},
		}
		var max uint
		for i := 0; i < len(req.Options); i++ {
			if err := utils.Verify(req.Options[i], verify); err != nil {
				response.FailWithMessage(fmt.Sprintf("%d Option:%s", i+1, err.Error()), c)
				return
			}
			if req.Options[i].Orders > max {
				max = req.Options[i].Orders
			}
		}
		if max != uint(len(req.Options)) {
			response.FailWithMessage(fmt.Sprintf("Order 参数输入错误"), c)
		}
	}
	if err := multipleChoiceService.Create(&req.MultipleChoice, req.ChapterSupport); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除选择题
func (choiceApi *MultipleChoiceApi) Delete(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := multipleChoiceService.Delete(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新选择题
func (choiceApi *MultipleChoiceApi) Update(c *gin.Context) {
	var req questionBankReq.MultipleChoiceUpdate
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"Id": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if len(req.Options) != 0 {
		verify := utils.Rules{
			"Describe": {utils.NotEmpty()},
			"Orders":   {utils.NotEmpty()},
		}
		var max uint
		for i := 0; i < len(req.Options); i++ {
			req.Options[i].MultipleChoiceId = req.ID
			if err := utils.Verify(req.Options[i], verify); err != nil {
				response.FailWithMessage(fmt.Sprintf("%d Option:%s", i+1, err.Error()), c)
				return
			}
			if req.Options[i].Orders > max {
				max = req.Options[i].Orders
			}
		}
		if max != uint(len(req.Options)) {
			response.FailWithMessage(fmt.Sprintf("Order 参数输入错误"), c)
		}
	}
	if err := multipleChoiceService.Update(req.MultipleChoice); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDetail  获取选择题详细
func (choiceApi *MultipleChoiceApi) FindDetail(c *gin.Context) {
	var req questionBankReq.DetailFind
	_ = c.ShouldBindQuery(&req)
	verify := utils.Rules{
		"Id": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	resp := questionBankResp.MultipleChoiceDetail{}
	if err := multipleChoiceService.FindDetail(&resp.MultipleChoice, req.Id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	if resp.MultipleChoice.ID != 0 {
		if err := questionBankService.FindCourseSupport(&resp.CourseSupport, req.Id, questionType.MULTIPLE_CHOICE); err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OkWithData(resp, c)
	} else {
		response.FailWithMessage("该选择题不存在", c)
		return
	}
}

// FindList  分页查找选择题
func (choiceApi *MultipleChoiceApi) FindList(c *gin.Context) {
	var pageInfo questionBankReq.MultipleChoiceFindList
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := multipleChoiceService.FindList(pageInfo); err != nil {
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
