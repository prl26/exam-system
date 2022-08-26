package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/lessondata"
)

type ArticleResourcesSearch struct {
	lessondata.ArticleResources
	request.PageInfo
}
