package examManage

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/prl26/exam-system/server/api/Backstage"
)

type ExamRouter struct {
}

func (e *ExamRouter) InitExamRouter(Router *gin.RouterGroup) {
	examRouterWithoutRecord := Router.Group("exam")
	var examApi = v1.ApiGroupApp.ExammanageApiGroup.ExamGroup.ExamApi
	{
		examRouterWithoutRecord.GET("findExamPaper", examApi.FindLessons) // 根据ID获取课程名
	}
}
