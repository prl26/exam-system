package questionBank

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/questionBank/enum/questionType"
	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/vo/request"
	questionBankResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
	"github.com/prl26/exam-system/server/model/teachplan"
	request2 "github.com/prl26/exam-system/server/model/teachplan/request"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
	"strconv"
	"time"
)

//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/prl26/exam-system/server/model/common/response"
//	"github.com/prl26/exam-system/server/model/enum/questionType"
//	"strconv"
//)
//
////
////import (
////	"github.com/gin-gonic/gin"
////	"github.com/prl26/exam-system/server/model/common/response"
////	"github.com/prl26/exam-system/server/model/enum/questionType"
////	"strconv"
////)
////
/////**
////
//// * @Author: AloneAtWar
////
//// * @Date:   2022/8/26 16:42
////
//// * @Note:
////
//// **/
////
//

type QuestionBankApi struct{}

// //
// //// FindQuestionsByKnowledgeId 通过KnowledgeId 获取所有题目
// ////func (q *QuestionBankApi) FindQuestionsByKnowledgeId(c *gin.Context) {
// ////	query := c.Query("id")
// ////	idInt, err := strconv.Atoi(query)
// ////	if err != nil {
// ////		response.FailWithMessage(err.Error(), c)
// ////		return
// ////	}
// ////	idUint := uint(idInt)
// ////	questions := questionBankService.FindQuestions(idUint)
// ////	response.OkWithData(questions, c)
// ////	return
// ////}
// //
// //func (q*QuestionBankApi) FindJudgesByChapterId(c *gin.Context){
// //	query := c.Query("id")
// //	idInt, err := strconv.Atoi(query)
// //	if err != nil {
// //		response.FailWithMessage(err.Error(), c)
// //		return
// //	}
// //	idUint := uint(idInt)
// //}
// //
// ////func(q*QuestionBankApi) FindQuestions(c*gin.Context){
// ////	lessonId , err := strconv.Atoi(c.Query("lessonId"))
// ////	if err != nil {
// ////		response.FailWithMessage(err.Error(), c)
// ////		return
// ////	}
// ////
// ////}
// //
var (
	lessonService   = service.ServiceGroupApp.BasicdataApiGroup.LessonService
	practiceService = service.ServiceGroupApp.QuestionBankServiceGroup.PracticeService
)

func (*QuestionBankApi) BeginPractice(c *gin.Context) {
	query := c.Query("lessonId")
	idInt, err := strconv.Atoi(query)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	lessonId := uint(idInt)
	detail, err := lessonService.FindLessonDetail(lessonId, false)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	studentId := utils.GetStudentId(c)
	go func() {
		practiceService.UpdatePracticeRecord(lessonId, studentId)
		now := time.Now()
		ip := c.ClientIP()
		r := &teachplan.PracticeRecord{
			LessonId:  lessonId,
			StudentId: studentId,
			BeginTime: now,
			BeginIp:   ip,
		}
		practiceService.CreatePracticeRecord(r)
	}()
	response.OkWithData(detail, c)
}

func (q *QuestionBankApi) FindQuestionsByChapterId(c *gin.Context) {
	questionT, err := strconv.Atoi(c.Query("questionType"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	switch questionType.QuestionType(questionT) {
	case questionType.JUDGE:
		q.FindJudge(c)
	case questionType.PROGRAM:
		q.FindProgram(c)
	case questionType.SINGLE_CHOICE:
		q.FindSingleChoice(c)
	case questionType.SUPPLY_BLANK:
		q.FindSupplyBlank(c)
	case questionType.MULTIPLE_CHOICE:
		q.FindMultiChoice(c)
	default:
		questionBankResp.CheckHandle(c, fmt.Errorf("题型输入错误"))
	}
}

func (q *QuestionBankApi) FindJudge(c *gin.Context) {
	var pageInfo questionBankReq.JudgePracticeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := questionBankService.FindJudgeList(pageInfo.JudgePracticeCriteria, pageInfo.PageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("获取失败:%s", err.Error()))
	} else {
		questionBankResp.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (q *QuestionBankApi) FindSingleChoice(c *gin.Context) {
	var pageInfo questionBankReq.MultipleChoicePracticeList
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := questionBankService.FindMultipleChoiceList(pageInfo.MultiplePracticeCriteria, pageInfo.PageInfo, false); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("获取失败:%s", err.Error()))
	} else {
		questionBankResp.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (q *QuestionBankApi) FindMultiChoice(c *gin.Context) {
	var pageInfo questionBankReq.MultipleChoicePracticeList
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := questionBankService.FindMultipleChoiceList(pageInfo.MultiplePracticeCriteria, pageInfo.PageInfo, true); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("获取失败:%s", err.Error()))
	} else {
		questionBankResp.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
func (q *QuestionBankApi) FindSupplyBlank(c *gin.Context) {
	var pageInfo questionBankReq.QuestionBankSupplyBlankPracticeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := questionBankService.FindSupplyBlankList(pageInfo.SupplyBlankPracticeCriteria, pageInfo.PageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("获取失败:%s", err.Error()))
	} else {
		questionBankResp.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (q *QuestionBankApi) FindProgram(c *gin.Context) {
	var pageInfo questionBankReq.ProgramPracticeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := questionBankService.FindProgramList(pageInfo.ProgramPracticeCriteria, pageInfo.PageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("获取失败:%s", err.Error()))
	} else {
		questionBankResp.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

//{"c":"#include \u003cstdio.h\u003e\n\nint main() {\n   printf(\"Hello World!\\n\");\n   return 0;\n}\n//更多请阅读：https://www.yiibai.com/c_examples/hello_world_program_in_c.html\n\n"}

func (*QuestionBankApi) FindHistoryAnswer(c *gin.Context) {
	var r request2.History
	_ = c.ShouldBindJSON(&r)
	verify := utils.Rules{
		"QuestionType": {utils.NotEmpty()},
	}
	if err := utils.Verify(r, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	studentId := utils.GetStudentId(c)
	answer := practiceService.FindHistoryAnswer(r.QuestionType, r.Ids, studentId)
	response.OkWithData(answer, c)
	return
}
