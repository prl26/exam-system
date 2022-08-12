// 自动生成模板TeachClass
package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// TeachClass 结构体
type TeachClass struct {
      global.GVA_MODEL
      Course_id  *int `json:"course_id" form:"course_id" gorm:"column:course_id;comment:课程id;"`
      Term_id  *int `json:"term_id" form:"term_id" gorm:"column:term_id;comment:学期id;"`
      Has_more  *int `json:"has_more" form:"has_more" gorm:"column:has_more;comment:是否还有其他专业的学生;"`
      Belong_class_id  *int `json:"belong_class_id" form:"belong_class_id" gorm:"column:belong_class_id;comment:所属班级id;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:教学班名称;"`
      Teacher_id  *int `json:"teacher_id" form:"teacher_id" gorm:"column:teacher_id;comment:所属老师id;"`
      Attendance_proportion  *float64 `json:"attendance_proportion" form:"attendance_proportion" gorm:"column:attendance_proportion;comment:考勤得分占比;"`
      Learn_resource_proportion  *float64 `json:"learn_resource_proportion" form:"learn_resource_proportion" gorm:"column:learn_resource_proportion;comment:学习资源得分占比;"`
      Final_exam_proportion  *float64 `json:"final_exam_proportion" form:"final_exam_proportion" gorm:"column:final_exam_proportion;comment:期末考得分占比;"`
      Procedure_exam_proportion  *float64 `json:"procedure_exam_proportion" form:"procedure_exam_proportion" gorm:"column:procedure_exam_proportion;comment:过程化考核得分;"`
}


// TableName TeachClass 表名
func (TeachClass) TableName() string {
  return "bas_teach_class"
}

