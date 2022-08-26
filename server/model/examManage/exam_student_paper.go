// 自动生成模板ExamStudentPaper
package examManage

import (
	"github.com/prl26/exam-system/server/global"
)

// ExamStudentPaper 结构体
type ExamStudentPaper struct {
	global.GVA_MODEL
	PaperId    *int   `json:"paperId" form:"paperId" gorm:"column:paper_id;comment:试卷id;size:32;"`
	QuestionId *int   `json:"questionId" form:"questionId" gorm:"column:question_id;comment:题目id;size:32;"`
	StudentId  *int   `json:"studentId" form:"studentId" gorm:"column:student_id;comment:学生id;size:32;"`
	Answer     string `json:"answer" form:"answer" gorm:"column:answer;comment:题目答案;size:255;"`
}

// TableName ExamStudentPaper 表名
func (ExamStudentPaper) TableName() string {
	return "exam_student_paper"
}
