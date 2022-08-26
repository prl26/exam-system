package request

import (
	"exam-system/model/basicdata"
	"exam-system/model/common/request"
)

type LearnResourcesChapterMergeSearch struct {
	basicdata.LearnResourcesChapterMerge
	request.PageInfo
}
