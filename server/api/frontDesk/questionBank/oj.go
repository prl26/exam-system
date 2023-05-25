package questionBank

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/questionBank/enum/questionType"
	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/vo/request"
	questionBankResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"strconv"
	"strings"
)

type OjApi struct {
}

var (
	judgeService          = service.ServiceGroupApp.QuestionBankServiceGroup.OjService.JudgeService
	programService        = &service.ServiceGroupApp.QuestionBankServiceGroup.OjService.ProgramService
	supplyBlankService    = service.ServiceGroupApp.QuestionBankServiceGroup.OjService.SupplyBlankService
	multipleChoiceService = service.ServiceGroupApp.QuestionBankServiceGroup.OjService.MultipleChoiceService
)

func (*OjApi) CheckJudge(c *gin.Context) {
	var r questionBankReq.CheckJudge
	_ = c.ShouldBindJSON(&r)
	verify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	if err := utils.Verify(r, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	result, lessonId, err := judgeService.Check(r.Id, r.Answer)
	if err != nil {
		response.FailWithMessage("找不到该题目", c)
		return
	}

	studentId := utils.GetStudentId(c)
	go func() {
		//t := practiceService.FindTheLatestRecord(lessonId, studentId)
		var score uint = 0
		if result {
			score = 100
		}
		answer := strconv.FormatBool(r.Answer)

		practiceService.CreatePracticeItem(questionType.JUDGE, r.Id, lessonId, studentId, score, answer)
		practiceService.UpdatePracticeAnswer(questionType.JUDGE, r.Id, lessonId, studentId, score, answer)
	}()
	response.OkWithData(result, c)
	return
}

func (*OjApi) CheckProgram(c *gin.Context) {
	var r questionBankReq.CheckProgramm
	_ = c.ShouldBindJSON(&r)
	verify := utils.Rules{
		"ID":         {utils.NotEmpty()},
		"Code":       {utils.NotEmpty()},
		"LanguageId": {utils.NotEmpty()},
	}
	if err := utils.Verify(r, verify); err != nil {
		questionBankResp.ErrorHandle(c, err)
		return
	}

	program, score, lessonId, err := programService.CheckProgram(r.Id, r.Code, r.LanguageId)
	studentId := utils.GetStudentId(c)
	go func() {
		//t := practiceService.FindTheLatestRecord(lessonId, studentId)
		practiceService.CreatePracticeItem(questionType.PROGRAM, r.Id, lessonId, studentId, score, r.Code)
		practiceService.UpdatePracticeAnswer(questionType.PROGRAM, r.Id, lessonId, studentId, score, r.Code)
	}()
	if err != nil {
		questionBankResp.ErrorHandle(c, err)
		return
	} else {
		response.OkWithData(program, c)
	}
}

func (*OjApi) CheckSupplyBlank(c *gin.Context) {
	var r questionBankReq.CheckSupplyBlank
	_ = c.ShouldBindJSON(&r)
	verify := utils.Rules{
		"ID":      {utils.NotEmpty()},
		"Answers": {utils.NotEmpty()},
	}
	if err := utils.Verify(r, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	result, score, lessonId, err := supplyBlankService.Check(r.Id, r.Answers)
	if err != nil {
		response.FailWithMessage("找不到该题目", c)
		return
	}
	studentId := utils.GetStudentId(c)
	go func() {
		//t := practiceService.FindTheLatestRecord(lessonId, studentId)
		answers := strings.Join(r.Answers, ",")
		practiceService.CreatePracticeItem(questionType.SUPPLY_BLANK, r.Id, lessonId, studentId, uint(score), answers)
		practiceService.UpdatePracticeAnswer(questionType.SUPPLY_BLANK, r.Id, lessonId, studentId, uint(score), answers)
	}()
	response.OkWithData(result, c)
	return
}

func (*OjApi) CheckMultipleChoice(c *gin.Context) {
	var r questionBankReq.CheckMultipleChoice
	_ = c.ShouldBindJSON(&r)
	verify := utils.Rules{
		"ID":      {utils.NotEmpty()},
		"Answers": {utils.NotEmpty()},
	}
	if err := utils.Verify(r, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	result, lessonId, err := multipleChoiceService.Check(r.Id, r.Answers)
	if err != nil {
		response.FailWithMessage("找不到该题目", c)
		return
	}
	t := questionType.SINGLE_CHOICE
	if len(r.Answers) > 1 {
		t = questionType.MULTIPLE_CHOICE
	}
	studentId := utils.GetStudentId(c)
	go func() {
		var score uint = 0
		if result {
			score = 100
		}
		answers := strings.Join(r.Answers, ",")
		practiceService.CreatePracticeItem(t, r.Id, lessonId, studentId, score, answers)
		practiceService.UpdatePracticeAnswer(t, r.Id, lessonId, studentId, score, answers)
	}()
	response.OkWithData(result, c)
	return
}
