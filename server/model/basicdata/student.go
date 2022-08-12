// 自动生成模板Student
package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Student 结构体
type Student struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:学生姓名;size:255;"`
      Sex  *int `json:"sex" form:"sex" gorm:"column:sex;comment:性别;"`
      Id_card  string `json:"id_card" form:"id_card" gorm:"column:id_card;comment:身份证号;size:255;"`
      Class_id  *int `json:"class_id" form:"class_id" gorm:"column:class_id;comment:班级id;"`
      Password  string `json:"password" form:"password" gorm:"column:password;comment:密码;size:255;"`
      Class_name  string `json:"class_name" form:"class_name" gorm:"column:class_name;comment:班级名称;size:255;"`
}


// TableName Student 表名
func (Student) TableName() string {
  return "bas_student"
}

