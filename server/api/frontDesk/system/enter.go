package system

import "github.com/prl26/exam-system/server/service"

type ApiGroup struct {
	BaseApi
	SystemApi
}

var (
	jwtService        = service.ServiceGroupApp.SystemServiceGroup.JwtService
	userService       = service.ServiceGroupApp.SystemServiceGroup.UserService
	teachClassService = service.ServiceGroupApp.BasicdataApiGroup.TeachClassService
	termService       = service.ServiceGroupApp.BasicdataApiGroup.TermService
	lessonService     = service.ServiceGroupApp.BasicdataApiGroup.LessonService
)
