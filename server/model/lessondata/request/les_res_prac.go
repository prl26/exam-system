package request

import (
	"exam-system/model/common/request"
	"exam-system/model/lessondata"
)

type ResourcePracticeSearch struct {
	lessondata.ResourcePractice
	request.PageInfo
}
