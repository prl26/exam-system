package request

import (
	"github.com/prl26/exam-system/server/model/basicdata"
	"github.com/prl26/exam-system/server/model/common/request"
)

type ChapterSearch struct {
	basicdata.Chapter
	request.PageInfo
}
