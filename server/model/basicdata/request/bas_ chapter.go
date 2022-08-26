package request

import (
	"exam-system/model/basicdata"
	"exam-system/model/common/request"
)

type ChapterSearch struct {
	basicdata.Chapter
	request.PageInfo
}
