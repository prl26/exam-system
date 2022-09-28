package questionBank

import (
	"fmt"
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

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/28 13:30

 * @Note:

 **/

type ProgramApi struct {
}

var programmService = service.ServiceGroupApp.QuestionBankServiceGroup.ProgrammService

//FindDetail  获取编程题的详细 信息 需要参数 programmId
func (p *ProgramApi) FindDetail(c *gin.Context) {
	var req questionBankReq.ProgramDetailFind
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"ProgramId": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var resp questionBankResp.ProgrammDetail

	if err := programmService.FindProgramDetail(&resp.Programms, req.ProgramId); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if resp.Programms.ID != 0 {
		if err := programmService.FindLanguageSupport(&resp.LanguageSupport, req.ProgramId); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}

		if err := questionBankService.FindCourseSupport(&resp.CourseSupport, req.ProgramId, questionType.PROGRAM); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	} else {
		response.FailWithMessage("无法找到该编程题", c)
	}

	response.OkWithData(resp, c)
}

// EditProgramDetail 编辑编程题
func (p *ProgramApi) EditProgramDetail(c *gin.Context) {
	var req questionBankReq.ProgramDetailEdit
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"Id":          {utils.NotEmpty()},
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
	programm := questionBank.Programm{}
	programm.BasicModel = req.BasicModel
	programm.ID = req.Id
	if err := programmService.EditProgrammDetail(&programm); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// DeleteProgramm 删除编程题
func (p *ProgramApi) DeleteProgramm(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := programmService.DeleteProgramm(IDS.Ids); err != nil {
		global.GVA_LOG.Error("", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// FindProgrammCases 寻找编程题支持的语言的用例
func (p *ProgramApi) FindProgrammCases(c *gin.Context) {
	var req questionBankReq.ProgramCaseFind
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"ProgramId":  {utils.NotEmpty()},
		"LanguageId": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var resp []questionBank.ProgrammCase
	if err := programmService.FindProgrammCases(&resp, req.ProgramId, req.LanguageId); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(resp, c)
}

//AddProgrammCase  增加编程题用例
func (p *ProgramApi) AddProgrammCase(c *gin.Context) {
	var req questionBankReq.ProgramCaseAdd
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"ProgramId":  {utils.NotEmpty()},
		"LanguageId": {utils.NotEmpty()},
		"Cases":      {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//Score      uint   `json:"score" form:"score" gorm:"column:score;comment:;"`
	//		Input      string `json:"input" form:"input" gorm:"column:input;comment:;"`
	//		Output     string `json:"output" form:"output" gorm:"column:output;comment:;"`
	verify = utils.Rules{
		"Name":   {utils.NotEmpty()},
		"Score":  {utils.NotEmpty()},
		"Output": {utils.NotEmpty()},
	}
	for i, s := range req.Cases {
		if err := utils.Verify(s, verify); err != nil {
			response.FailWithMessage(fmt.Sprintf("第%d个用例出现错误:%s", i+1, err.Error()), c)
		}
		return
	}

	programmCases := make([]questionBank.ProgrammCase, len(req.Cases))
	for i := 0; i < len(programmCases); i++ {
		programmCases[i].ProgrammId = req.ProgramId
		programmCases[i].LanguageId = req.LanguageId
		programmCases[i].ProgrammLimit = req.Cases[i].ProgrammLimit
		programmCases[i].Name = req.Cases[i].Name
		programmCases[i].Score = req.Cases[i].Score
		programmCases[i].Input = req.Cases[i].Input
		programmCases[i].Output = req.Cases[i].Output
	}

	if err := programmService.AddProgrammCase(&programmCases); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// EditProgrammCases 编辑编程题用例
func (p *ProgramApi) EditProgrammCases(c *gin.Context) {
	var req questionBankReq.ProgramCaseEdit
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"Id":     {utils.NotEmpty()},
		"Name":   {utils.NotEmpty()},
		"Score":  {utils.NotEmpty()},
		"Output": {utils.NotEmpty()},
	}
	for i, s := range req.Cases {
		if err := utils.Verify(s, verify); err != nil {
			response.FailWithMessage(fmt.Sprintf("第%d个用例出现错误:%s", i+1, err.Error()), c)
		}
		return
	}
	if err := programmService.EditProgrammCases(req.Cases); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

//DeleteProgrammCases 删除编程题用例
func (p *ProgramApi) DeleteProgrammCases(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := programmService.DeleteProgrammCases(IDS.Ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// AddLanguageSupport  增加语言支持
func (p *ProgramApi) AddLanguageSupport(c *gin.Context) {
	var req questionBankReq.LanguageSupportAdd
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"ProgramId":  {utils.NotEmpty()},
		"LanguageId": {utils.NotEmpty()},
		"Cases":      {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	for i, s := range req.Cases {
		if err := utils.Verify(s, verify); err != nil {
			response.FailWithMessage(fmt.Sprintf("第%d个用例出现错误:%s", i+1, err.Error()), c)
		}
		return
	}

	programmCases := make([]questionBank.ProgrammCase, len(req.Cases))
	for i := 0; i < len(programmCases); i++ {
		programmCases[i].ProgrammId = req.ProgramId
		programmCases[i].LanguageId = req.LanguageId
		programmCases[i].ProgrammLimit = req.Cases[i].ProgrammLimit
		programmCases[i].Name = req.Cases[i].Name
		programmCases[i].Score = req.Cases[i].Score
		programmCases[i].Input = req.Cases[i].Input
		programmCases[i].Output = req.Cases[i].Output
	}

	languageSupport := questionBank.ProgrammLanguageMerge{}
	languageSupport.LanguageId = req.LanguageId
	languageSupport.ProgrammId = req.ProgramId
	languageSupport.DefaultCode = req.DefaultCode
	languageSupport.ReferenceAnswer = req.ReferenceAnswer

	err := programmService.AddLanguageSupport(&languageSupport, &programmCases)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

//EditLanguageSupport  编辑语言支持
func (p *ProgramApi) EditLanguageSupport(c *gin.Context) {
	var req questionBankReq.LanguageSupportEdit
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"id":        {utils.NotEmpty()},
		"ProgramId": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var languageSupport questionBank.ProgrammLanguageMerge
	languageSupport.ID = req.Id
	languageSupport.ProgrammId = req.ProgramId
	languageSupport.DefaultCode = req.DefaultCode
	languageSupport.ReferenceAnswer = req.ReferenceAnswer

	if err := programmService.EditLanguageSupport(&languageSupport); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)

}

// DeleteLanguageSupport 删除语言支持
func (p *ProgramApi) DeleteLanguageSupport(c *gin.Context) {
	var req questionBankReq.LanguageSupportDelete
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"ProgramId": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if len(req.LanguageIds) != 0 {
		if err := programmService.DeleteLanguageSupport(req.ProgramId, req.LanguageIds); err != nil {
			response.FailWithMessage("批量删除失败", c)
		}
	}
	response.Ok(c)
}
