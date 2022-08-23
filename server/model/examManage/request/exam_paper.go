package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/examManage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ExamPaperSearch struct{
    examManage.ExamPaper
    request.PageInfo
}
