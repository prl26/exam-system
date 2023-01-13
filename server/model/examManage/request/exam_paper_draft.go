package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
)

type ConvertDraft struct {
	PlanId       int    `json:"planId" form:"planId"`
	Name         string `json:"name" form:"name"`
	DraftPaperId uint   `json:"draftPaperId" form:"draftPaperId"`
}
type IdsReq struct {
	Ids []uint `json:"ids" form:"ids"`
}
type DraftPaperSearch struct {
	request.PageInfo
	DraftSearch
}
type DraftSearch struct {
	Name     string `json:"name" form:"name" gorm:"column:name;comment:试卷名称;size:64;"`
	LessonId uint   `json:"lessonId" form:"lessonId"`
	UserId   *uint  `json:"user_id" form:"userId" gorm:"column:user_id;comment:创建id;size:32;"`
}
