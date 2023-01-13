package request

import "github.com/prl26/exam-system/server/model/common/request"

type TeachClassSituation struct {
	request.PageInfo
	LessonId     uint
	TeachClassId uint
}

type StudentSituation struct {
	request.PageInfo
	LessonId  uint
	StudentId uint
}

type SituationDetail struct {
	request.PageInfo
	RecordId uint
}
