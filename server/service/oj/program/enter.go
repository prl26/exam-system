package program

import (
	ojResp "github.com/prl26/exam-system/server/model/oj/response"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBankEnum "github.com/prl26/exam-system/server/model/questionBank/enum"
	"github.com/prl26/exam-system/server/service/oj/program/cLanguage"
	"github.com/prl26/exam-system/server/service/oj/program/common"
)

type ProgramService struct {
	cLanguage.CLanguageService
	common.CommonService
}

func (s *ProgramService) CheckProgram(id uint, code string, languageId questionBankEnum.LanguageType) ([]*ojResp.Submit, uint, error) {
	program, err := s.CommonService.FindProgram(id)
	if err != nil {
		return nil, 0, err
	}
	support := questionBankBo.LanguageSupport{}
	err = support.Deserialize(program.LanguageSupports, languageId)
	if err != nil {
		return nil, 0, err
	}
	cases := questionBankBo.ProgramCases{}
	err = cases.Deserialize(program.ProgramCases)
	if err != nil {
		return nil, 0, err
	}
	switch languageId {
	case questionBankEnum.C_LANGUAGE:
		result, sum, err := s.CLanguageService.Check(code, support.LanguageLimit, cases)
		if err != nil {
			return nil, 0, err
		}
		return result, sum, nil
	default:
		return nil, 0, err
	}
}
