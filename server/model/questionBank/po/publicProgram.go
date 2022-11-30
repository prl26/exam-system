package po

import (
	"github.com/prl26/exam-system/server/global"
)

type PublicProgram struct {
	global.GVA_MODEL
	ProgramModel
}

func (PublicProgram) TableName() string {
	return "les_questionBank_publicProgram"
}
