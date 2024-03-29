package questionBank

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBankPo "github.com/prl26/exam-system/server/model/questionBank/po"
	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/vo/request"
	questionBankResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
	"strconv"
)

type ProgramApi struct {
}

var programService = service.ServiceGroupApp.QuestionBankServiceGroup.ProgramService

// Create 创建编程题
func (p *ProgramApi) Create(c *gin.Context) {
	var req questionBankReq.ProgramCreate
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req.BasicModel, questionBankReq.BaseVerify); err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}
	programPo := questionBankPo.Program{}
	if req.ProgramCases != nil {
		programCaseStr, err := req.ProgramCases.Serialize()
		if err != nil {
			questionBankResp.ErrorHandle(c, err)
			return
		}
		programPo.ProgramCases = &programCaseStr
	} else {
		//questionBankResp.ErrorHandle(c, fmt.Errorf("未输入编程题用例"))
		//return
	}
	if req.LanguageSupports != nil {
		languageSupportStr, brief, err := req.LanguageSupports.Serialize()
		if err != nil {
			questionBankResp.ErrorHandle(c, err)
			return
		}
		programPo.LanguageSupports = &languageSupportStr
		programPo.LanguageSupportsBrief = &brief
	} else {
		//questionBankResp.ErrorHandle(c, fmt.Errorf("未输入语言支持"))
		//return
	}
	if req.DefaultCodes != nil {
		defaultCodeStr, err := req.DefaultCodes.Serialize()
		if err != nil {
			questionBankResp.ErrorHandle(c, err)
			return
		}
		programPo.DefaultCodes = &defaultCodeStr
	}
	if req.ReferenceAnswers != nil {
		referenceAnswerStr, err := req.ReferenceAnswers.Serialize()
		if err != nil {
			questionBankResp.ErrorHandle(c, err)
			return
		}
		programPo.ReferenceAnswers = &referenceAnswerStr
	}
	programPo.BasicModel = req.BasicModel
	programPo.CourseSupport = req.CourseSupport
	if err := programService.Create(&programPo); err != nil {
		questionBankResp.ErrorHandle(c, err)
		return
	} else {
		questionBankResp.OkWithMessage("创建成功", c)
	}
}

func (api *ProgramApi) FindList(c *gin.Context) {
	var pageInfo questionBankReq.ProgramSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := programService.FindList(pageInfo.ProgramSearchCriteria, pageInfo.PageInfo); err != nil {
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

func (api *ProgramApi) FindDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}
	if detail, err := programService.FindDetail(id); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, err)
	} else {
		if detail == nil {
			questionBankResp.NotFind(c)
			return
		} else {
			programDetail := questionBankResp.ProgramDetail{}
			programDetail.GVA_MODEL = detail.GVA_MODEL
			programDetail.BasicModel = detail.BasicModel
			programDetail.CourseSupportPtr = detail.CourseSupportPtr
			programDetail.CourseSupport = detail.CourseSupport
			if detail.ProgramCases != nil {
				if err := programDetail.ProgramCases.Deserialize(*detail.ProgramCases); err != nil {
					global.GVA_LOG.Error(err.Error())
					questionBankResp.ErrorHandle(c, err)
					return
				}
			}
			if detail.LanguageSupports != nil && *detail.LanguageSupports != "" {
				programDetail.LanguageSupports = new(questionBankBo.LanguageSupports)
				if err := programDetail.LanguageSupports.Deserialization(*detail.LanguageSupports); err != nil {
					global.GVA_LOG.Error(err.Error())
					questionBankResp.ErrorHandle(c, err)
					return
				}
			}
			if detail.ReferenceAnswers != nil && *detail.ReferenceAnswers != "" {
				programDetail.ReferenceAnswers = new(questionBankBo.ReferenceAnswers)
				if err := programDetail.ReferenceAnswers.Deserialization(*detail.ReferenceAnswers); err != nil {
					global.GVA_LOG.Error(err.Error())
					questionBankResp.ErrorHandle(c, err)
					return
				}
			}
			if detail.DefaultCodes != nil && *detail.DefaultCodes != "" {
				programDetail.DefaultCodes = new(questionBankBo.DefaultCodes)
				if err := programDetail.DefaultCodes.Deserialization(*detail.DefaultCodes); err != nil {
					global.GVA_LOG.Error(err.Error())
					questionBankResp.ErrorHandle(c, err)
					return
				}
			}
			questionBankResp.OkWithDetailed(programDetail, "获取成功", c)
		}
	}
}

func (api *ProgramApi) Update(c *gin.Context) {
	var req questionBankReq.ProgramUpdate
	_ = c.ShouldBindJSON(&req)
	if req.Id == 0 {
		questionBankResp.CheckHandle(c, fmt.Errorf("请输入修改ID"))
		return
	}
	programPo := questionBankPo.Program{}
	programPo.ID = req.Id
	if req.ProgramCases != nil {
		programCaseStr, err := req.ProgramCases.Serialize()
		if err != nil {
			questionBankResp.ErrorHandle(c, err)
			return
		}
		programPo.ProgramCases = &programCaseStr
	} else {
		// 修改的时候不一定修改编程题用例
		//questionBankResp.ErrorHandle(c, fmt.Errorf("未输入编程题用例"))
		//return
	}
	if req.LanguageSupports != nil {
		languageSupportStr, brief, err := req.LanguageSupports.Serialize()
		if err != nil {
			questionBankResp.ErrorHandle(c, err)
			return
		}
		programPo.LanguageSupports = &languageSupportStr
		programPo.LanguageSupportsBrief = &brief
	} else {
		// 修改的时候不一定修改语言支持
		//questionBankResp.ErrorHandle(c, fmt.Errorf("未输入编程题用例"))
		//return
	}
	if req.DefaultCodes != nil {
		defaultCodeStr, err := req.DefaultCodes.Serialize()
		if err != nil {
			questionBankResp.ErrorHandle(c, err)
			return
		}
		programPo.DefaultCodes = &defaultCodeStr
	}
	if req.ReferenceAnswers != nil {
		referenceAnswerStr, err := req.ReferenceAnswers.Serialize()
		if err != nil {
			questionBankResp.ErrorHandle(c, err)
			return
		}
		programPo.ReferenceAnswers = &referenceAnswerStr
	}
	programPo.BasicModel = req.BasicModel
	programPo.CourseSupport = req.CourseSupport
	if err := programService.Update(&programPo); err != nil {
		questionBankResp.ErrorHandle(c, err)
		return
	} else {
		questionBankResp.OkWithMessage("更新成功", c)
	}
}

func (api *ProgramApi) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}
	if err := programService.Delete([]uint{uint(id)}); err != nil {
		questionBankResp.ErrorHandle(c, err)
		return
	} else {
		questionBankResp.OkWithMessage("删除成功", c)
	}
}

func (api *ProgramApi) Deletes(c *gin.Context) {
	var req request.IdsReq
	_ = c.ShouldBindJSON(&req)
	if err := programService.Delete(req.Ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("批量删除失败:%s", err.Error()))
	} else {
		questionBankResp.OkWithMessage("批量删除成功", c)
	}
}
