package questionBank

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/response"
	questionBankResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
	"github.com/prl26/exam-system/server/model/teachplan"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type TargetApi struct {
}

var targetService = service.ServiceGroupApp.QuestionBankServiceGroup.TargetService

func (*TargetApi) BeginPractice(c *gin.Context) {
	query := c.Query("lessonId")
	idInt, err := strconv.Atoi(query)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	lessonId := uint(idInt)
	detail, err := lessonService.FindLessonDetail(lessonId, true)

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

func (*TargetApi) FindTargetByKnowledgeId(c *gin.Context) {
	//query := c.Query("id")
	//idInt, err := strconv.Atoi(query)
	//if err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	//knowledgeId := uint(idInt)
	//targetService.FindTargetByKnowledgeId(knowledgeId)
}

func (*TargetApi) FindTargetDetail(c *gin.Context) {
	query := c.Query("id")
	idInt, err := strconv.Atoi(query)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	targetId := uint(idInt)
	detail, err := targetService.FindDetail(targetId)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("获取失败:%s", err.Error()))
		return
	} else {
		q := &questionBankResp.TargetDetail{
			TargetDetail:      detail,
			IsGenerateAddress: false,
			Address:           "",
		}
		questionBankResp.OkWithDetailed(q, "获取成功", c)
		return
	}

}
