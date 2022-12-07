// 自动生成模板Score
package teachplan

import (
	"github.com/prl26/exam-system/server/global"
)

// Score 结构体
type Score struct {
	global.GVA_MODEL
	StudentId                *int     `json:"studentId" form:"studentId" gorm:"column:student_id;comment:学生id;size:32;"`
	CourseId                 *int     `json:"courseId" form:"courseId" gorm:"column:course_id;comment:课程Id;size:32;"`
	CourseName               string   `json:"courseName" form:"courseName" gorm:"column:course_name;comment:课程名称;size:16;"`
	TeachClassName           string   `json:"teachClassName" form:"teachClassName" gorm:"column:teach_class_name;comment:教学班名称;size:16;"`
	TeachClassId             *int     `json:"teachClassId" form:"teachClassId" gorm:"column:teach_class_id;comment:教学班id;size:32;"`
	AttendanceProportion     *float64 `json:"attendanceProportion" form:"attendanceProportion" gorm:"column:attendance_proportion;comment:考勤占比;size:8;"`
	AttendanceScore          *float64 `json:"attendanceScore" form:"attendanceScore" gorm:"column:attendance_score;comment:考勤得分;size:8;"`
	LearnResourcesProportion *float64 `json:"learnResourcesProportion" form:"learnResourcesProportion" gorm:"column:learn_resources_proportion;comment:学习资源占比;size:8;"`
	LearnResourcesScore      *float64 `json:"learnResourcesScore" form:"learnResourcesScore" gorm:"column:learn_resources_score;comment:学习资源得分;size:8;"`
	ProcedureScore           *float64 `json:"procedureScore" form:"procedureScore" gorm:"column:procedure_score;comment:过程化考核得分;size:8;"`
	ProcedureProportion      *float64 `json:"procedureProportion" form:"procedureProportion" gorm:"column:procedure_proportion;comment:过程化考核占比;size:8;"`
	ExamScrore               *float64 `json:"examScore" form:"examScore" gorm:"column:exam_score;comment:期末考试成绩;size:8;"`
	ExamProporation          *float64 `json:"examProporation" form:"examProporation" gorm:"column:exam_proporation;comment:期末考试占比;size:8;"`
	PlanId                   *uint    `json:"planId" form:"planId" gorm:"column:plan_id"`
	TermId                   *int     `json:"termId" form:"TermId" gorm:"column:term_id"`
	TotalScore               *float64 `json:"totalScore" form:"totalScore" gorm:"column:total_score"`
}

// TableName Score 表名
func (Score) TableName() string {
	return "tea_score"
}
