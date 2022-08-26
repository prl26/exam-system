/*
*

	@author: qianyi  2022/8/24 19:18:00
	@note:
*/
package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
)

// 接收 教学班id 和学生id 的结构体
type StuTeachClass struct {
	TeachClassId uint   `json:"teachClassId" from:"teachClassId"`
	StudentIds   []uint `json:"studentIds"`
}

type AddStudentByClass struct {
	TeachClassId uint `json:"teachClassId" from:"teachClassId"`
	ClassId      uint `json:"classId" form:"classId"`
}

type TeachClassStudent struct {
	TeachClassId uint `json:"teachClassId" form:"teachClassId"`
	request.PageInfo
}

type Association struct {
	teachClassID uint `gorm:"column:teach_class_id"`
	studentId    uint `gorm:"column:student_id"`
}
