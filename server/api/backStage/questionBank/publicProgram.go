package questionBank

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	questionBankPo "github.com/prl26/exam-system/server/model/questionBank/po"
	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/vo/request"
	questionBankResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
	"strconv"
)

type PublicProgramApi struct {
}

var publicProgramService = service.ServiceGroupApp.QuestionBankServiceGroup.PublicProgramService

// Create 创建公共编程题
func (p *PublicProgramApi) Create(c *gin.Context) {
	var req questionBankReq.PublicProgramCreate
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req.BasicModel, questionBankReq.BaseVerify); err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}
	programPo := questionBankPo.PublicProgram{}
	if len(req.ProgramCases) != 0 {
		programCaseStr, err := req.ProgramCases.Serialize()
		if err != nil {
			questionBankResp.ErrorHandle(c, err)
			return
		}
		programPo.ProgramCases = programCaseStr
	} else {
		//questionBankResp.ErrorHandle(c, fmt.Errorf("未输入编程题用例"))
		//return
	}
	if len(req.LanguageSupports) != 0 {
		languageSupportStr, err := req.LanguageSupports.Serialize()
		if err != nil {
			questionBankResp.ErrorHandle(c, err)
			return
		}
		programPo.LanguageSupports = languageSupportStr
	} else {
		//questionBankResp.ErrorHandle(c, fmt.Errorf("未输入语言支持"))
		//return
	}
	if len(req.ReferenceAnswers) != 0 {
		languageSupportStr, err := req.ReferenceAnswers.Serialize()
		if err != nil {
			questionBankResp.ErrorHandle(c, err)
			return
		}
		programPo.ReferenceAnswers = languageSupportStr
	}
	if len(req.DefaultCodes) != 0 {
		languageSupportStr, err := req.DefaultCodes.Serialize()
		if err != nil {
			questionBankResp.ErrorHandle(c, err)
			return
		}
		programPo.DefaultCodes = languageSupportStr
	}
	programPo.BasicModel = req.BasicModel
	if err := publicProgramService.Create(&programPo); err != nil {
		questionBankResp.ErrorHandle(c, err)
		return
	} else {
		questionBankResp.OkWithMessage("创建成功", c)
	}
}

func (api *PublicProgramApi) FindList(c *gin.Context) {
	var pageInfo questionBankReq.PublicProgramSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := publicProgramService.FindList(pageInfo.PublicProgramSearchCriteria, pageInfo.PageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, err)
	} else {
		questionBankResp.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (api *PublicProgramApi) FindDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}
	if detail, err := publicProgramService.FindDetail(id); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, err)
	} else {
		if detail == nil {
			questionBankResp.NotFind(c)
			return
		} else {
			programDetail := questionBankResp.PublicProgramDetail{}
			programDetail.GVA_MODEL = detail.GVA_MODEL
			programDetail.BasicModel = detail.BasicModel
			if detail.ProgramCases != "" {
				if err := programDetail.ProgramCases.Deserialize(detail.ProgramCases); err != nil {
					global.GVA_LOG.Error(err.Error())
					questionBankResp.ErrorHandle(c, err)
					return
				}
			}
			if detail.LanguageSupports != "" {
				if err := programDetail.LanguageSupports.Deserialization(detail.LanguageSupports); err != nil {
					global.GVA_LOG.Error(err.Error())
					questionBankResp.ErrorHandle(c, err)
					return
				}
			}
			if detail.ReferenceAnswers != "" {
				if err := programDetail.ReferenceAnswers.Deserialization(detail.ReferenceAnswers); err != nil {
					global.GVA_LOG.Error(err.Error())
					questionBankResp.ErrorHandle(c, err)
					return
				}
			}
			if detail.DefaultCodes != "" {
				if err := programDetail.DefaultCodes.Deserialization(detail.DefaultCodes); err != nil {
					global.GVA_LOG.Error(err.Error())
					questionBankResp.ErrorHandle(c, err)
					return
				}
			}
			questionBankResp.OkWithDetailed(programDetail, "获取成功", c)
		}
	}
}

func (api *PublicProgramApi) Update(c *gin.Context) {
	var req questionBankReq.PublicProgramUpdate
	_ = c.ShouldBindJSON(&req)
	if req.Id == 0 {
		questionBankResp.CheckHandle(c, fmt.Errorf("请输入修改ID"))
		return
	}
	programPo := questionBankPo.PublicProgram{}
	programPo.ID = req.Id
	if len(req.ProgramCases) != 0 {
		programCaseStr, err := req.ProgramCases.Serialize()
		if err != nil {
			questionBankResp.ErrorHandle(c, err)
			return
		}
		programPo.ProgramCases = programCaseStr
	} else {
		// 修改的时候不一定修改编程题用例
		//questionBankResp.ErrorHandle(c, fmt.Errorf("未输入编程题用例"))
		//return
	}
	if len(req.LanguageSupports) != 0 {
		languageSupportStr, err := req.LanguageSupports.Serialize()
		if err != nil {
			questionBankResp.ErrorHandle(c, err)
			return
		}
		programPo.LanguageSupports = languageSupportStr
	} else {
		// 修改的时候不一定修改语言支持
		//questionBankResp.ErrorHandle(c, fmt.Errorf("未输入编程题用例"))
		//return
	}
	if len(req.DefaultCodes) != 0 {
		defaultCodeStr, err := req.DefaultCodes.Serialize()
		if err != nil {
			questionBankResp.ErrorHandle(c, err)
			return
		}
		programPo.DefaultCodes = defaultCodeStr
	}
	if len(req.ReferenceAnswers) != 0 {
		referenceAnswerStr, err := req.ReferenceAnswers.Serialize()
		if err != nil {
			questionBankResp.ErrorHandle(c, err)
			return
		}
		programPo.ReferenceAnswers = referenceAnswerStr
	}
	programPo.BasicModel = req.BasicModel
	if err := publicProgramService.Update(&programPo); err != nil {
		questionBankResp.ErrorHandle(c, err)
		return
	} else {
		questionBankResp.OkWithMessage("更新成功", c)
	}
}

func (api *PublicProgramApi) Migrate(c *gin.Context) {
	var req questionBankReq.PublicProgramMigration
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"Id":        {utils.NotEmpty()},
		"ChapterId": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}
	if err := publicProgramService.Migrate([]uint{req.Id}, req.PublicProgramMigration); err != nil {
		questionBankResp.ErrorHandle(c, err)
		return
	} else {
		questionBankResp.OkWithMessage("迁移成功", c)
	}
}

func (api *PublicProgramApi) Migrates(c *gin.Context) {
	var req questionBankReq.PublicProgramMigrations
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"Ids":       {utils.NotEmpty()},
		"ChapterId": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}
	if err := publicProgramService.Migrate(req.Ids, req.PublicProgramMigration); err != nil {
		questionBankResp.ErrorHandle(c, err)
		return
	} else {
		questionBankResp.OkWithMessage("迁移成功", c)
	}
}

func (api *PublicProgramApi) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}
	if err := publicProgramService.Delete([]uint{uint(id)}); err != nil {
		questionBankResp.ErrorHandle(c, err)
		return
	} else {
		questionBankResp.OkWithMessage("删除成功", c)
	}
}

func (api *PublicProgramApi) Deletes(c *gin.Context) {
	var req request.IdsReq
	_ = c.ShouldBindJSON(&req)
	if err := publicProgramService.Delete(req.Ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("批量删除失败:%s", err.Error()))
	} else {
		questionBankResp.OkWithMessage("批量删除成功", c)
	}
}
