// 自动生成模板TeachClass
package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// TeachClass 结构体
type TeachClass struct {
	global.GVA_MODEL
	Course_id                  *int   `json:"course_id" form:"course_id" gorm:"column:course_id;comment:课程id;"`
	Term_id                    *int   `json:"term_id" form:"term_id" gorm:"column:term_id;comment:学期id;"`
	Name                       string `json:"name" form:"name" gorm:"column:name;comment:教学班名称;"`
	Teacher_id                 *int   `json:"teacher_id" form:"teacher_id" gorm:"column:teacher_id;comment:后台user;"`
	Attendance_proportion      *int   `json:"attendance_proportion" form:"attendance_proportion" gorm:"column:attendance_proportion;comment:考勤得分占比;"`
	Learn_resources_proportion *int   `json:"learn_resources_proportion" form:"learn_resources_proportion" gorm:"column:learn_resources_proportion;comment:学习资源得分占比;"`
	Final_exam_proportion      *int   `json:"final_exam_proportion" form:"final_exam_proportion" gorm:"column:final_exam_proportion;comment:期末考得分占比;"`
	Procedure_exam_proportion  *int   `json:"procedure_exam_proportion" form:"procedure_exam_proportion" gorm:"column:procedure_exam_proportion;comment:过程化考核得分占比;"`
	Homework_proportion        *int   `json:"homework_proportion" form:"homework_proportion" gorm:"column:homework_proportion;comment:作业得分占比;"`
}

// TableName TeachClass 表名
func (TeachClass) TableName() string {
	return "bas_teach_class"
}
