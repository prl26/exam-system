package request

import (
	"exam-system/model/common/request"
	"exam-system/model/lessondata"
)

type VideoResourcesSearch struct {
	lessondata.VideoResources
	request.PageInfo
}
