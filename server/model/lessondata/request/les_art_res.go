package request

import (
	"exam-system/model/common/request"
	"exam-system/model/lessondata"
)

type ArticleResourcesSearch struct {
	lessondata.ArticleResources
	request.PageInfo
}
