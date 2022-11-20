package questionBank

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/prl26/exam-system/server/api"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/26 20:26

 * @Note:

 **/

type OjRouter struct {
}



func (s *OjRouter) InitBackgroundOjRouter(Router *gin.RouterGroup) {
	programmRouter := Router.Group("oj/program")
	programmApi :=v1.ApiGroupApp.BackStage.QuestionBankApiGroup.OjApi
	{
		programmRouter.POST("compile", programmApi.Compile) //编译
		programmRouter.POST("execute", programmApi.Execute) //运行
	}

}
