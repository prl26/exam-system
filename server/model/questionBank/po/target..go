package po

import "github.com/prl26/exam-system/server/global"

type Target struct {
	global.GVA_MODEL
	TargetModel
	CourseSupport
}

type TargetModel struct {
	BasicModel
	Code     string `json:"code" form:"code"`
	ByteCode string `json:"byteCode" form:"byteCode"`
}

func (Target) TableName() string {
	return "les_questionBank_target"
}
