package basic

import "github.com/prl26/exam-system/server/service"

type ApiGroup struct{
	LessonApi
}

var(
	lessonService = service.ServiceGroupApp.BasicdataApiGroup.LessonService
)