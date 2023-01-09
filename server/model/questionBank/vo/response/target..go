package response

import (
	"github.com/prl26/exam-system/server/global"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBank "github.com/prl26/exam-system/server/model/questionBank/po"
)

type TargetSimple struct {
	global.GVA_MODEL
	questionBank.SimpleModel
}

type TargetDetail struct {
	*questionBankBo.TargetDetail
	IsGenerateAddress bool
	Address           string
}

type TargetSimplePractice struct {
	global.GVA_MODEL
	questionBank.BasicModel
}
