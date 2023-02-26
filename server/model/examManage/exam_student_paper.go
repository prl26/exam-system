// 自动生成模板ExamStudentPaper
package examManage

import (
	"github.com/prl26/exam-system/server/global"
	questionBankEnum "github.com/prl26/exam-system/server/model/questionBank/enum/languageType"
	"github.com/prl26/exam-system/server/model/questionBank/enum/questionType"
)

// ExamStudentPaper 结构体
type ExamStudentPaper struct {
	global.GVA_MODEL
	PaperId      *uint                      `json:"paperId" form:"paperId" gorm:"column:paper_id;comment:试卷id;size:32;"`
	QuestionId   *uint                      `json:"questionId" form:"questionId" gorm:"column:question_id;comment:题目id;size:32;"`
	StudentId    *uint                      `json:"studentId" form:"studentId" gorm:"column:student_id;comment:学生id;size:32;"`
	Answer       string                     `json:"answer" form:"answer" gorm:"column:answer;comment:该生题目答案;size:16000;"`
	PlanId       *uint                      `json:"planId" form:"planId" gorm:"column:plan_id;comment:教学计划id;size:32"`
	Score        *float64                   `json:"score" form:"score" gorm:"column:score;comment:本题分值;size:8;"`
	QuestionType *questionType.QuestionType `json:"questionType" form:"paperId" gorm:"column:question_type;comment:题目类型;size:8;"`
	ProblemType  *int                       `json:"problemType" form:"problemType" gorm:"column:problem_type;comment:难度;size:8;"`
	GotScore     *float64                   `json:"gotScore" form:"gotScore" gorm:"column:got_score;comment:该生得分l;size:4"`
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
type CommitExamPaper1 struct {
	StudentId            uint                   `json:"studentId" form:"studentId"`
	PlanId               uint                   `json:"planId" form:"planId"`
	PaperId              uint                   `json:"paperId" form:"paperId"`
	MultipleChoiceCommit []MultipleChoiceCommit `json:"multipleChoiceCommit"`
	JudgeCommit          []JudgeCommit          `json:"judgeCommit"`
	BlankCommit          []BlankCommit          `json:"blankCommit"`
	ProgramCommit        []CommitProgram        `json:"programCommit" form:"programCommit"`
}
type ReExecExamPaper struct {
	StudentId            uint                    `json:"studentId" form:"studentId"`
	PlanId               uint                    `json:"planId" form:"planId"`
	PaperId              uint                    `json:"paperId" form:"paperId"`
	MultipleChoiceCommit []MultipleChoiceCommit1 `json:"multipleChoiceCommit"`
	JudgeCommit          []JudgeCommit1          `json:"judgeCommit"`
	BlankCommit          []BlankCommit1          `json:"blankCommit"`
}

type MultipleChoiceCommit1 struct {
	Id         uint   `json:"id"`
	QuestionId uint   `json:"questionId" form:"questionId"`
	Answer     string `json:"answer" form:"answer"`
}
type JudgeCommit1 struct {
	Id         uint `json:"id"`
	QuestionId uint `json:"questionId" form:"questionId"`
	Answer     bool `json:"answer" form:"answer"`
}
type BlankCommit1 struct {
	Id         uint   `json:"id"`
	QuestionId uint   `json:"questionId" form:"questionId"`
	Answer     string `json:"answer" form:"answer"`
}
type MultipleChoiceCommit struct {
	MergeId    uint     `json:"mergeId"`
	QuestionId uint     `json:"questionId" form:"questionId"`
	Answer     []string `json:"answer" form:"answer"`
}
type JudgeCommit struct {
	MergeId    uint `json:"mergeId"`
	QuestionId uint `json:"questionId" form:"questionId"`
	Answer     bool `json:"answer" form:"answer"`
}
type BlankCommit struct {
	MergeId    uint     `json:"mergeId"`
	QuestionId uint     `json:"questionId" form:"questionId"`
	Answer     []string `json:"answer" form:"answer"`
}

type CommitProgram struct {
	PlanId     uint                          `json:"planId" form:"planId"`
	MergeId    uint                          `json:"mergeId"`
	QuestionId uint                          `json:"questionId" form:"questionId"`
	StudentId  uint                          `json:"studentId" form:"studentId"`
	Code       string                        `json:"code"`
	LanguageId questionBankEnum.LanguageType `json:"languageId"`
}
type QuesNum struct {
	StudentId uint `json:"studentId"`
	Num       uint `json:"num"`
}
type QuesScore struct {
	StudentId uint `json:"studentId"`
	Score     int  `json:"score"`
}
