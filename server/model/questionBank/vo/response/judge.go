package response

import (
	"github.com/prl26/exam-system/server/global"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBankPo "github.com/prl26/exam-system/server/model/questionBank/po"
)

type JudgeSimple struct {
	global.GVA_MODEL
	questionBankPo.SimpleModel
}

type JudgePractice struct {
	questionBankBo.PracticeModel
}
