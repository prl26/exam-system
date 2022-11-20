package questionBank

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/enum/questionType"
	"github.com/prl26/exam-system/server/model/questionBank"
	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/request"
	questionBankResp "github.com/prl26/exam-system/server/model/questionBank/response"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
)

type SupplyBlankApi struct {
}

var supplyBlankService = service.ServiceGroupApp.QuestionBankServiceGroup.SupplyBlankService

// Create 创建填空题
func (api *SupplyBlankApi) Create(c *gin.Context) {
	var req questionBankReq.SupplyBlankCreate
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"ProblemType": {utils.NotEmpty()},
		"CanPractice": {utils.NotEmpty()},
		"CanExam":     {utils.NotEmpty()},
		"Title":       {utils.NotEmpty()},
		"Describe":    {utils.NotEmpty()},
		"Answers":     {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	supplyBlank := questionBank.SupplyBlank{}
	supplyBlank.BasicModel = req.BasicModel
	supplyBlank.IsOrder = req.IsOrder
	if a, b, err := req.Answers.GetAnswersAndProportions(); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		supplyBlank.Answer = a
		supplyBlank.Proportion = b
		supplyBlank.Num = len(req.Answers)
	}
	if err := supplyBlankService.Create(&supplyBlank, req.LessonSupports); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除填空题
func (api *SupplyBlankApi) Delete(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := supplyBlankService.Delete(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新填空题
func (api *SupplyBlankApi) Update(c *gin.Context) {
	var req questionBankReq.SupplyBlankUpdate
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"Id": {utils.NotEmpty()},
	}
	supplyBlank := questionBank.SupplyBlank{}
	supplyBlank.ID = req.Id
	supplyBlank.BasicModel = req.BasicModel
	supplyBlank.IsOrder = req.IsOrder
	if a, b, err := req.Answers.GetAnswersAndProportions(); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		supplyBlank.Answer = a
		supplyBlank.Proportion = b
		supplyBlank.Num = len(req.Answers)
	}
	if err := utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := supplyBlankService.UpdateQuestionBankSupplyBlank(supplyBlank); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindList  分页查找填空题
func (api *SupplyBlankApi) FindList(c *gin.Context) {
	var pageInfo questionBankReq.QuestionBankSupplyBlankSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := supplyBlankService.FindList(pageInfo); err != nil {
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

// FindDetail  获取填空题详细
func (api *SupplyBlankApi) FindDetail(c *gin.Context) {
	var req questionBankReq.DetailFind
	_ = c.ShouldBindQuery(&req)
	verify := utils.Rules{
		"Id": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	resp := questionBankResp.SupplyBlankDetail{}
	if err := supplyBlankService.FindDetail(&resp.SupplyBlank, req.Id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	}
	if resp.SupplyBlank.ID == 0 {
		response.FailWithMessage("该填空题不存在", c)
		return
	} else {
		if err := questionBankService.FindCourseSupport(&resp.CourseSupport, resp.SupplyBlank.ID, questionType.SUPPLY_BLANK); err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		}
		response.OkWithData(resp, c)
	}
}
