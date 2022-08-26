package request

import (
	"exam-system/model/common/request"
	"exam-system/model/lessondata"
)

type KnowledgeSearch struct {
	lessondata.Knowledge
	request.PageInfo
}
