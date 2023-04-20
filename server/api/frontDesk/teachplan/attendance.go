package teachplan

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/model/common/response"
	teachplanReq "github.com/prl26/exam-system/server/model/teachplan/request"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"strings"
	"time"
)

type AttendanceApi struct {
}

var teachplanService = service.ServiceGroupApp.TeachplanServiceGroup.TeachAttendanceService

func (a AttendanceApi) Attendance(c *gin.Context) {
	var req teachplanReq.Attendance
	_ = c.ShouldBindJSON(&req)

	// escape
	req.Code = strings.ReplaceAll(req.Code, " ", "+")

	crypto := utils.Decryption(req.Code)
	ExpirationTimeStr := ""
	AttendanceId := 0
	split := strings.Split(crypto, ",")
	_, err := fmt.Sscanf(split[0], "%d", &AttendanceId)
	if err != nil {
		response.FailWithMessage("Code错误，无法签到", c)
		return
	}
	ExpirationTimeStr = split[1]
	now := time.Now()
	toTime := utils.StringToTime(ExpirationTimeStr)
	if now.After(toTime) {
		response.CheckHandle(c, fmt.Errorf("签到已过期，请下课后找任课老师确认!"))
		return
	}
	// 对于IP进行检查
	ip, _ := c.RemoteIP()
	attendance, err := teachplanService.Attendance(req.StudentId, ip.String(), uint(AttendanceId), 1)
	if err != nil {
		response.FailWithMessage("服务器错误", c)
		return
	}
	if attendance == 0 {
		response.CheckHandle(c, fmt.Errorf("输入学号错误,或不再该班级"))
		return
	}
	response.OkWithMessage("签到成功！", c)
}
