package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
)

type PaperTemplateSearch struct {
	PtSearch
	request.PageInfo
}
type PtSearch struct {
	LessonId int    `json:"lessonId" form:"lessonId"`
	Name     string `json:"name" form:"name"`
}
