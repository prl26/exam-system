package service

import (
	"github.com/prl26/exam-system/server/service/basicdata"
	"github.com/prl26/exam-system/server/service/examManage"
	"github.com/prl26/exam-system/server/service/lessondata"
	"github.com/prl26/exam-system/server/service/questionBank"
	"github.com/prl26/exam-system/server/service/system"
	"github.com/prl26/exam-system/server/service/teachplan"
)

type ServiceGroup struct {
	SystemServiceGroup       system.ServiceGroup
	BasicdataApiGroup        basicdata.ServiceGroup
	LessondataServiceGroup   lessondata.ServiceGroup
	TeachplanServiceGroup    teachplan.ServiceGroup
	ExammanageServiceGroup   examManage.ServiceGroup
	QuestionBankServiceGroup questionBank.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
