package request

import (
	"exam-system/model/common/request"
	"exam-system/model/lessondata"
)

type ResourcesTestSearch struct {
	lessondata.ResourcesTest
	request.PageInfo
}
