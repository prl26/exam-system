// 自动生成模板ExamStudentPaper
package examManage

import (
	"github.com/prl26/exam-system/server/global"
)

// ExamStudentPaper 结构体
type ExamStudentPaper struct {
	global.GVA_MODEL
	PaperId      *uint  `json:"paperId" form:"paperId" gorm:"column:paper_id;comment:试卷id;size:32;"`
	QuestionId   *uint  `json:"questionId" form:"questionId" gorm:"column:question_id;comment:题目id;size:32;"`
	StudentId    *uint  `json:"studentId" form:"studentId" gorm:"column:student_id;comment:学生id;size:32;"`
	Answer       string `json:"answer" form:"answer" gorm:"column:answer;comment:该生题目答案;size:16000;"`
	PlanId       *uint  `json:"planId" form:"planId" gorm:"column:plan_id;comment:教学计划id;size:32"`
	Score        *uint  `json:"score" form:"score" gorm:"column:score;comment:本题分值;size:8;"`
	QuestionType *int   `json:"questionType" form:"paperId" gorm:"column:question_type;comment:题目类型;size:8;"`
	ProblemType  *int   `json:"problemType" form:"problemType" gorm:"column:problem_type;comment:难度;size:8;"`
	GotScore     *uint  `json:"gotScore" form:"gotScore" gorm:"column:got_score;comment:该生得分"`
}

// TableName ExamStudentPaper 表名
func (ExamStudentPaper) TableName() string {
	return "exam_student_paper"
}

type CommitExamPaper struct {
	StudentId            uint                   `json:"studentId" form:"studentId"`
	PlanId               uint                   `json:"planId" form:"planId"`
	PaperId              uint                   `json:"paperId" form:"paperId"`
	MultipleChoiceCommit []MultipleChoiceCommit `json:"multipleChoiceCommit"`
	JudgeCommit          []JudgeCommit          `json:"judgeCommit"`
	BlankCommit          []BlankCommit          `json:"blankCommit"`
}
type MultipleChoiceCommit struct {
	MergeId    uint  `json:"mergeId"`
	QuestionId uint  `json:"questionId" form:"questionId"`
	Answers    []int `json:"answers" form:"answers"`
}
type JudgeCommit struct {
	MergeId    uint   `json:"mergeId"`
	QuestionId uint   `json:"questionId" form:"questionId"`
	Answer     string `json:"answer" form:"answer"`
}
type BlankCommit struct {
	MergeId    uint   `json:"mergeId"`
	QuestionId uint   `json:"questionId" form:"questionId"`
	Answer     string `json:"answer" form:"answer"`
}
