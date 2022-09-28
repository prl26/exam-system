package examManage

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/teachplan"
	"github.com/prl26/exam-system/server/utils"
	"time"
)

type ExamStatusService struct {
}

func (examStatusServices *ExamPaperService) GaSStudentsOfExam() (students []uint, err error) {
	var ExamPlan []teachplan.ExamPlan
	var teachClassIds []uint
	err = global.GVA_DB.Where("start_time <= ?", time.Now().Add(10*time.Minute)).Select("teach_class_id").Find(&ExamPlan).Error
	if err != nil {
		return nil, err
	}
	for _, v := range ExamPlan {
		teachClassIds = append(teachClassIds, *v.TeachClassId)
	}
	students, err = utils.GaSStudentsOfTeachClass(teachClassIds)
	if err != nil {
		return nil, err
	}
	return students, err
}
