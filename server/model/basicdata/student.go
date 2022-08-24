// 自动生成模板Student
package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Student 结构体
type Student struct {
	global.GVA_MODEL
	Name            string `json:"name" form:"name" gorm:"column:name;comment:姓名;"`
	Sex             string `json:"sex" form:"sex" gorm:"column:sex;comment:性别;"`
	Id_card         string `json:"id_card" form:"id_card" gorm:"column:id_card;comment:身份证号;"`
	Password        string `json:"password" form:"password" gorm:"column:password;comment:密码;"`
	College_id      *int   `json:"college_id" form:"college_id" gorm:"column:college_id;comment:学院id;"`
	Professional_id *int   `json:"professional_id" form:"professional_id" gorm:"column:professional_id;comment:专业id;"`
	Class_id        *int   `json:"class_id" form:"class_id" gorm:"column:class_id;comment:班级id;"`
}

// TableName Student 表名
func (Student) TableName() string {
	return "bas_student"
}
