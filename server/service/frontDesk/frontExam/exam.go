package frontExam

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/examManage"
	"github.com/prl26/exam-system/server/model/examManage/request"
	"github.com/prl26/exam-system/server/model/teachplan"
)

type ExamService struct {
}

func (examService *ExamService) FindExamPlans(teachClassId uint) (examPlans []teachplan.ExamPlan, err error) {
	err = global.GVA_DB.Where("teach_class_id = ?", teachClassId).Find(&examPlans).Error
	return
}

func (examService *ExamService) GetExamPapers(examComing request.ExamComing) (examPaper examManage.ExamPaper, err error) {

	return
}
