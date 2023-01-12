// 自动生成模板Student
package basicdata

import (
	"github.com/prl26/exam-system/server/global"
)

// Student 结构体
type Student struct {
	global.GVA_MODEL
	Name             string       `json:"name" form:"name" gorm:"column:name;comment:姓名;"`
	Sex              string       `json:"sex" form:"sex" gorm:"column:sex;comment:性别;"`
	IdCard           string       `json:"id_card" form:"id_card" gorm:"column:id_card;comment:身份证号;"`
	Password         string       `json:"password" form:"password" gorm:"column:password;comment:密码;"`
	CollegeName      string       `json:"collegeName" form:"collegeName" gorm:"column:college_name;comment:学院名称;"`
	ProfessionalName string       `json:"professionalName" form:"professionalName" gorm:"column:professional_name;comment:专业名称;"`
	ClassId          uint         `json:"classId" form:"classId" gorm:"column:class_id;comment:班级id;"`
	TeachClass       []TeachClass `gorm:"many2many:bas_student_teachClass;"`
}

// TableName Student 表名
func (Student) TableName() string {
	return "bas_student"
}
