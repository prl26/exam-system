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
	File           *multipart.FileHeader `json:"file" form:"file"`
	ClassId        uint                  `json:"classId" form:"classId"`
	CollegeId      uint                  `json:"collegeId"  form:"collegeId"`
	ProfessionalId uint                  `json:"professionalId"  form:"professionalId"`
}
