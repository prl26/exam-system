package examManage

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/examManage"
	"github.com/prl26/exam-system/server/model/teachplan"
	"github.com/prl26/exam-system/server/utils"
	"time"
)

type ExamStatusService struct {
}

//查找那些教学班即将进行考试,并将对应教学班的学生拉入redis黑名单中
func (examStatusServices *ExamStatusService) GaSStudentsOfExam() (students []uint, err error) {
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
func (student_paper_status *ExamStatusService) GetStatus(StudentId uint, PlanId uint) (status examManage.StudentPaperStatus, err error) {
	err = global.GVA_DB.Table("student_paper_status").Where("student_id = ? and plan_id = ?", StudentId, PlanId).Find(&status).Error
	if err != nil {
		return
	}
	return
}
