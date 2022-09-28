package basicdata

//教学班与学生关联表
type StudentAndTeachClass struct {
	StudentId    uint `json:"studentId" gorm:"column:student_id;comment:学生id"`
	TeachClassId uint `json:"teachClassId" gorm:"column:teach_class_id; comment:教学班id"`
}

// TableName Student 表名
func (StudentAndTeachClass) TableName() string {
	return "bas_student_class_teachclasses"
}
