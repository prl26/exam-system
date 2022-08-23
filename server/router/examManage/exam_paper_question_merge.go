package examManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PaperQuestionMergeRouter struct {
}

// InitPaperQuestionMergeRouter 初始化 PaperQuestionMerge 路由信息
func (s *PaperQuestionMergeRouter) InitPaperQuestionMergeRouter(Router *gin.RouterGroup) {
	试卷题目表Router := Router.Group("试卷题目表").Use(middleware.OperationRecord())
	试卷题目表RouterWithoutRecord := Router.Group("试卷题目表")
	var 试卷题目表Api = v1.ApiGroupApp.ExammanageApiGroup.PaperQuestionMergeApi
	{
		试卷题目表Router.POST("createPaperQuestionMerge", 试卷题目表Api.CreatePaperQuestionMerge)   // 新建PaperQuestionMerge
		试卷题目表Router.DELETE("deletePaperQuestionMerge", 试卷题目表Api.DeletePaperQuestionMerge) // 删除PaperQuestionMerge
		试卷题目表Router.DELETE("deletePaperQuestionMergeByIds", 试卷题目表Api.DeletePaperQuestionMergeByIds) // 批量删除PaperQuestionMerge
		试卷题目表Router.PUT("updatePaperQuestionMerge", 试卷题目表Api.UpdatePaperQuestionMerge)    // 更新PaperQuestionMerge
	}
	{
		试卷题目表RouterWithoutRecord.GET("findPaperQuestionMerge", 试卷题目表Api.FindPaperQuestionMerge)        // 根据ID获取PaperQuestionMerge
		试卷题目表RouterWithoutRecord.GET("getPaperQuestionMergeList", 试卷题目表Api.GetPaperQuestionMergeList)  // 获取PaperQuestionMerge列表
	}
}
