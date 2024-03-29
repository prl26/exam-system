package examManage

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/questionBank/enum/questionType"
	"time"
)

type ExamRecord struct {
	global.GVA_MODEL
	StudentId uint `json:"studentId"`
	PlanId    uint `json:"planId" form:"planId"`
	EnterTime time.Time
	EndTime   time.Time
	Ip        string `json:"ip"` //学生ip
}

func (ExamRecord) TableName() string {
	return "exam_record"
}

type ExamRecordMerge struct {
	global.GVA_MODEL
	PaperId      *uint                      `json:"paperId" form:"paperId" gorm:"column:paper_id;comment:试卷id;size:32;"`
	QuestionId   *uint                      `json:"questionId" form:"questionId" gorm:"column:question_id;comment:题目id;size:32;"`
	StudentId    *uint                      `json:"studentId" form:"studentId" gorm:"column:student_id;comment:学生id;size:32;"`
	Answer       string                     `json:"answer" form:"answer" gorm:"column:answer;comment:该生题目答案;size:16000;"`
	PlanId       *uint                      `json:"planId" form:"planId" gorm:"column:plan_id;comment:教学计划id;size:32"`
	Score        *float64                   `json:"score" form:"score" gorm:"column:score;comment:本题分值;size:8;"`
	QuestionType *questionType.QuestionType `json:"questionType" form:"paperId" gorm:"column:question_type;comment:题目类型;size:8;"`
	ProblemType  *int                       `json:"problemType" form:"problemType" gorm:"column:problem_type;comment:难度;size:8;"`
	GotScore     *float64                   `json:"gotScore" form:"gotScore" gorm:"column:got_score;comment:该生得分"`
	RecordId     *uint                      `json:"recordId" gorm:"record_id"`
}

func (ExamRecordMerge) TableName() string {
	return "exam_record_merge"
}
