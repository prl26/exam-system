package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/lessondata"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ArticleResourcesSearch struct{
    lessondata.ArticleResources
    request.PageInfo
}
