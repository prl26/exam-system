package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/teachplan"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TeachAttendanceSearch struct{
    teachplan.TeachAttendance
    request.PageInfo
}
