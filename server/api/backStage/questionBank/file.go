package questionBank

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/service"
	"go.uber.org/zap"
)

var FrontSystemService = service.ServiceGroupApp.SystemServiceGroup.SystemService

func UploadFile(c *gin.Context) {
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	file, _, err := FrontSystemService.UploadTitleFile(header)
	if err != nil {
		global.GVA_LOG.Error("上传图片失败!", zap.Error(err))
		response.FailWithMessage("上传图片失败!", c)
		return
	}
	response.OkWithDetailed(file, "上传成功", c)
}
