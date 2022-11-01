package frontDesk

import "github.com/prl26/exam-system/server/service/frontDesk/frontExam"

type ServiceGroup struct {
	CommonService
	frontExam.ServiceGroup
}
