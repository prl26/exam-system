package response

import "github.com/prl26/exam-system/server/model/teachplan"

type ExamPlanRp struct {
	teachplan.ExamPlan
	ItemName
}
type ItemName struct {
	TermName       string `json:"termName" form:"termName"`
	LessonName     string `json:"lessonName" form:"lessonName"`
	TeachClassName string `json:"teachClassName" form:"teachClassName"`
}
