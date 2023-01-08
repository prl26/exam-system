package program

import (
	"github.com/prl26/exam-system/server/global"
	commonError "github.com/prl26/exam-system/server/model/common/error"
	questionBankEnum "github.com/prl26/exam-system/server/model/questionBank/enum/languageType"

	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBankError "github.com/prl26/exam-system/server/model/questionBank/error"
	questionBank "github.com/prl26/exam-system/server/model/questionBank/po"
	"time"
)

type ProgramService struct {
}

var table = make(map[questionBankEnum.LanguageType]IProgramService)

func Register(languageType questionBankEnum.LanguageType, program IProgramService) {
	table[languageType] = program
}

func (s *ProgramService) Compile(code string, languageId questionBankEnum.LanguageType) (string, *time.Time, error) {
	if service, ok := table[languageId]; ok {
		return service.Compile(code)
	} else {
		return "", nil, questionBankError.NotLanguageSupportError
	}
}

func (s *ProgramService) Execute(languageId questionBankEnum.LanguageType, fileId string, input string, limit questionBankBo.LanguageLimit) (string, *questionBankBo.ExecuteSituation, error) {
	if service, ok := table[languageId]; ok {
		return service.Execute(fileId, input, limit)
	} else {
		return "", nil, questionBankError.NotLanguageSupportError
	}
}

func (s *ProgramService) CheckProgram(id uint, code string, languageId questionBankEnum.LanguageType) ([]*questionBankBo.Submit, uint, uint, error) {
	program, err := s.findOjHelper(id)
	if err != nil {
		return nil, 0, 0, err
	}
	support := questionBankBo.LanguageSupport{}
	err = support.Deserialize(program.LanguageSupports, languageId)
	if err != nil {
		return nil, 0, 0, err
	}
	cases := questionBankBo.ProgramCases{}
	err = cases.Deserialize(program.ProgramCases)
	if err != nil {
		return nil, 0, 0, err
	}
	submits, score, err := table[languageId].Check(code, support.LanguageLimit, cases)
	if err != nil {
		return nil, 0, 0, err
	}
	return submits, score, program.LessonId, nil
}

// 寻找题目所对应的 Oj支持   (测试用例与语言限制)
func (s *ProgramService) findOjHelper(id uint) (*questionBankBo.OjHelper, error) {
	result := &questionBankBo.OjHelper{}
	if r := global.GVA_DB.Model(&questionBank.Program{}).Find(&result, id); r.Error != nil {
		return nil, r.Error
	} else if r.RowsAffected == 0 {
		return nil, commonError.NotFoundError
	} else {
		return result, nil
	}
}
