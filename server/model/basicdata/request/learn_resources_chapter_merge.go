package request

import (
	"github.com/prl26/exam-system/server/model/basicdata"
	"github.com/prl26/exam-system/server/model/common/request"
)

type LearnResourcesChapterMergeSearch struct {
	basicdata.LearnResourcesChapterMerge
	request.PageInfo
}
