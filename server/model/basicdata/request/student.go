package request

import (
	"github.com/prl26/exam-system/server/model/basicdata"
	"github.com/prl26/exam-system/server/model/common/request"
	"mime/multipart"
)

type StudentSearch struct {
	basicdata.Student
	request.PageInfo
}

type StudentExcel struct {
	File     *multipart.FileHeader `json:"file" form:"file"`
	ClassId  uint                  `json:"classId" form:"classId"`
	TermId   uint                  `json:"termId" form:"termId"`
	CourseId uint                  `json:"courseId" form:"courseId"`
	// 不要学院 专业了 学生基础管理只暴露 根据学号查询就行
	//CollegeId      uint                  `json:"collegeId"  form:"collegeId"`
	//ProfessionalId uint                  `json:"professionalId"  form:"professionalId"`
}
