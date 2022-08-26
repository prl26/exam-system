package request

import (
	"exam-system/model/basicdata"
	"exam-system/model/common/request"
)

type LessonSearch struct {
	basicdata.Lesson
	request.PageInfo
}
