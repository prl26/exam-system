package request

import "github.com/prl26/exam-system/server/model/common/request"

type TeachClassSituation struct {
	request.PageInfo
	LessonId     uint `form:"lessonId"`
	TeachClassId uint `form:"teachClassId"`
}

type StudentSituation struct {
	request.PageInfo
	LessonId  uint `form:"lessonId"`
	StudentId uint `form:"studentId"`
}

type SituationDetail struct {
	request.PageInfo
	RecordId uint `form:"recordId"`
}
