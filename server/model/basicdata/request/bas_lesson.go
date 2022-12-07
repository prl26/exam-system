package request

import (
	"github.com/prl26/exam-system/server/model/basicdata"
	"github.com/prl26/exam-system/server/model/common/request"
)

type LessonSearch struct {
	basicdata.Lesson
	request.PageInfo
}
type FrontLessonSearch struct {
	basicdata.Lesson
}
