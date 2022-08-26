package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/lessondata"
)

type ResourcesTestSearch struct {
	lessondata.ResourcesTest
	request.PageInfo
}
