package examManage

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/questionBank/enum/questionType"
)

type ExamPaperDraft struct {
	global.GVA_MODEL
	Name      string                    `json:"name" form:"name" gorm:"column:name;comment:试卷名称;size:64;"`
	LessonId  uint                      `json:"lessonId" form:"lessonId"`
	UserId    *uint                     `json:"userId" form:"userId" gorm:"column:user_id;comment:创建id;size:32;"`
	PaperItem []DraftPaperQuestionMerge `json:"paperItem" gorm:"foreignKey:DraftPaperId"`
}
type ExamPaperDraft1 struct {
	global.GVA_MODEL
	Name     string `json:"name" form:"name" gorm:"column:name;comment:试卷名称;size:64;"`
	LessonId uint   `json:"lessonId" form:"lessonId"`
	UserId   *uint  `json:"userId" form:"userId" gorm:"column:user_id;comment:创建id;size:32;"`
}

func (ExamPaperDraft1) TableName() string {
	return "exam_paper_draft"
}
func (ExamPaperDraft) TableName() string {
	return "exam_paper_draft"
}

type DraftPaperQuestionMerge struct {
	global.GVA_MODEL
	DraftPaperId *uint                      `json:"draftPaperId" form:"draftPaperId" gorm:"column:draft_paper_id"`
	QuestionId   *uint                      `json:"questionId" form:"paperId" gorm:"column:question_id;comment:题目id;size:32;"`
	Score        *int                       `json:"score" form:"score" gorm:"column:score;comment:所占分值;size:8;"`
	QuestionType *questionType.QuestionType `json:"questionType" form:"paperId" gorm:"column:question_type;comment:题目类型;size:8;"`
	ProblemType  *int                       `json:"problemType" form:"problemType" gorm:"column:problem_type;comment:难度;size:8;"`
}

func (DraftPaperQuestionMerge) TableName() string {
	return "exam_draft_paper_merge"
}
