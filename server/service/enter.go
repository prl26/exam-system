package service

import (
	"exam-system/service/basicdata"
	"exam-system/service/examManage"
	"exam-system/service/lesson"
	"exam-system/service/lessondata"
	"exam-system/service/questionBank"
	"exam-system/service/system"
	"exam-system/service/teachplan"
)

type ServiceGroup struct {
	SystemServiceGroup       system.ServiceGroup
	LessonServiceGroup       lesson.ServiceGroup
	BasicdataApiGroup        basicdata.ServiceGroup
	LessondataServiceGroup   lessondata.ServiceGroup
	TeachplanServiceGroup    teachplan.ServiceGroup
	ExammanageServiceGroup   examManage.ServiceGroup
	QuestionBankServiceGroup questionBank.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
