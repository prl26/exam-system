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
type ExamPlanRp1 struct {
	Plan   teachplan.ExamPlan `json:"plan"`
	Status PlanStatus         `json:"status"`
}
type PlanStatus struct {
	IsBegin          int `json:"isBegin" form:"isBegin"`
	IsCommit         int `json:"isCommit" form:"isCommit"`
	IsFinishPreExams int `json:"isFinishPreExams"`
}
