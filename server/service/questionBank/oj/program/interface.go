package program

import (
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	"time"
)

type IProgramService interface {
	Compile(code string) (string, *time.Time, error)
	Execute(fileId string, input string, programmLimit questionBankBo.LanguageLimit) (string, *questionBankBo.ExecuteSituation, error)
	Check(code string, limit questionBankBo.LanguageLimit, cases questionBankBo.ProgramCases) ([]*questionBankBo.Submit, uint, error) //用于给后台检查代码
}
