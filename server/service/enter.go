package service

import (
	"github.com/prl26/exam-system/server/service/basicdata"
	"github.com/prl26/exam-system/server/service/examManage"
	"github.com/prl26/exam-system/server/service/frontDesk"
	"github.com/prl26/exam-system/server/service/lesson"
	"github.com/prl26/exam-system/server/service/lessondata"
	"github.com/prl26/exam-system/server/service/oj"
	"github.com/prl26/exam-system/server/service/questionBank"
	"github.com/prl26/exam-system/server/service/system"
	"github.com/prl26/exam-system/server/service/teachplan"
)

type ServiceGroup struct {
	SystemServiceGroup       system.ServiceGroup
	LessonServiceGroup       lesson.ServiceGroup
	BasicdataApiGroup        basicdata.ServiceGroup
	LessondataServiceGroup   lessondata.ServiceGroup
	TeachplanServiceGroup    teachplan.ServiceGroup
	ExammanageServiceGroup   examManage.ServiceGroup
	QuestionBankServiceGroup questionBank.ServiceGroup
	OjServiceServiceGroup    oj.ServiceGroup
	FrontServiceGroup        frontDesk.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
