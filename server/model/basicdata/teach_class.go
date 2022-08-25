// 自动生成模板TeachClass
package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// TeachClass 结构体
type TeachClass struct {
	global.GVA_MODEL
	CourseId                 *int      `json:"courseId" form:"courseId" gorm:"column:course_id;comment:课程id;"`
	TermId                   *int      `json:"termId" form:"termId" gorm:"column:term_id;comment:学期id;"`
	Name                     string    `json:"name" form:"name" gorm:"column:name;comment:教学班名称;"`
	TeacherId                *int      `json:"teacherId" form:"teacherId" gorm:"column:teacher_id;comment:后台user;"`
	AttendanceProportion     *int      `json:"attendanceProportion" form:"attendanceProportion" gorm:"column:attendance_proportion;comment:考勤得分占比;"`
	LearnResourcesProportion *int      `json:"learnResourcesProportion" form:"learn_resourcesProportion" gorm:"column:learn_resources_proportion;comment:学习资源得分占比;"`
	FinalExamProportion      *int      `json:"finalExamProportion" form:"finalExamProportion" gorm:"column:final_exam_proportion;comment:期末考得分占比;"`
	ProcedureExamProportion  *int      `json:"procedureExamProportion" form:"procedureExamProportion" gorm:"column:procedure_exam_proportion;comment:过程化考核得分占比;"`
	HomeworkProportion       *int      `json:"homeworkProportion" form:"homeworkProportion" gorm:"column:homework_proportion;comment:作业得分占比;"`
	Student                  []Student `gorm:"many2many:bas_student_teachClass;"`
}

// TableName TeachClass 表名
func (TeachClass) TableName() string {
	return "bas_teach_class"
}
