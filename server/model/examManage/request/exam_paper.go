package request

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/examManage"
	"github.com/prl26/exam-system/server/model/questionBank/enum/questionType"
)

type ExamPaperSearch struct {
	EpSearch
	request.PageInfo
}
type PaperDistribution struct {
	PlanId uint `json:"planId" form:"planId"`
}
type EpSearch struct {
	PlanId     int    `json:"planId" form:"planId" gorm:"column:plan_id;comment:考试计划id;size:32;"`
	Name       string `json:"name" form:"name" gorm:"column:name;comment:试卷名称;size:64;"`
	TemplateId int    `json:"templateId" form:"templateId" gorm:"column:template_id;comment:试卷模板Id;size:32;"`
	TermId     uint   `json:"termId" form:"termId"`
	LessonId   uint   `json:"lessonId" form:"lessonId"`
	UserId     uint   `json:"user_id" form:"userId" gorm:"column:user_id;comment:创建id;size:32;"`
}
type ExamPaperBySelf struct {
	global.GVA_MODEL
	PlanId    *int                            `json:"planId" form:"planId" gorm:"column:plan_id;comment:考试计划id;size:32;"`
	Name      string                          `json:"name" form:"name" gorm:"column:name;comment:试卷名称;size:64;"`
	UserId    *uint                           `json:"user_id" form:"userId" gorm:"column:user_id;comment:创建id;size:32;"`
	PaperItem []examManage.PaperQuestionMerge `json:"paperItem"`
}
type ExamPaperItem struct {
	QuestionId   *uint                      `json:"questionId" form:"paperId" gorm:"column:question_id;comment:题目id;size:32;"`
	Score        *int                       `json:"score" form:"score" gorm:"column:score;comment:所占分值;size:8;"`
	QuestionType *questionType.QuestionType `json:"questionType" form:"paperId" gorm:"column:question_type;comment:题目类型;size:8;"`
	ProblemType  *int                       `json:"problemType" form:"problemType" gorm:"column:problem_type;comment:难度;size:8;"`
}
