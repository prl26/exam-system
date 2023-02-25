package teachplan

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api"
)

type AttendanceRouter struct {
}

func (a AttendanceRouter) InitAttendanceRouter(Router *gin.RouterGroup) {
	TeachPlanWithoutRecord := Router.Group("teachplan")
	var attendanceApi = api.ApiGroupApp.FrontDesk.TeachplanApiGroup.AttendanceApi
	{
		TeachPlanWithoutRecord.POST("attendance", attendanceApi.Attendance)

	}
}
