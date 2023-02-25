package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBankPo "github.com/prl26/exam-system/server/model/questionBank/po"
	"mime/multipart"
)

type TargetCreate struct {
	questionBankPo.Target
}

type TargetSearch struct {
	questionBankBo.TargetSearchCriteria
	request.PageInfo
}

type TargetPracticeSearch struct {
	questionBankBo.TargetPracticeCriteria
	request.PageInfo
}

type TargetExcel struct {
	File     *multipart.FileHeader `json:"file" form:"file"`
	LessonId uint                  `json:"lessonId" form:"lessonId"`
}

type RankingList struct {
	LessonId uint `json:"lessonId" form:"lessonId"`
	request.PageInfo
}
