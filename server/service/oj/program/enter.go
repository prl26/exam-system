package program

import (
	"github.com/prl26/exam-system/server/model/oj"
	ojResp "github.com/prl26/exam-system/server/model/oj/response"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBankEnum "github.com/prl26/exam-system/server/model/questionBank/enum"
	questionBankError "github.com/prl26/exam-system/server/model/questionBank/error"
	"github.com/prl26/exam-system/server/service/oj/program/cLanguage"
	"github.com/prl26/exam-system/server/service/oj/program/common"
	goLanguage "github.com/prl26/exam-system/server/service/oj/program/go"
	"time"
)

type ProgramService struct {
	cLanguage.CLanguageService
	goLanguage.GoLanguage
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
	case questionBankEnum.GO_LANGUAGE:
		result, sum, err := s.GoLanguage.Check(code, support.LanguageLimit, cases)
		if err != nil {
			return nil, 0, err
		}
		return result, sum, nil
	default:
		return nil, 0, err
	}
}

func (s *ProgramService) Compile(code string, languageId questionBankEnum.LanguageType) (string, *time.Time, error) {
	switch languageId {
	case questionBankEnum.C_LANGUAGE:
		compile, t, err := s.CLanguageService.Compile(code)
		return compile, t, err
	case questionBankEnum.GO_LANGUAGE:
		return s.GoLanguage.Compile(code)
	default:
		return "", nil, questionBankError.NotLanguageSupportError
	}
}

func (s *ProgramService) Execute(languageId questionBankEnum.LanguageType, fileId string, input string, limit questionBankBo.LanguageLimit) (string, *oj.ExecuteSituation, error) {
	switch languageId {
	case questionBankEnum.C_LANGUAGE:
		return s.CLanguageService.Execute(fileId, input, limit)
	case questionBankEnum.GO_LANGUAGE:
		return s.GoLanguage.Execute(fileId, input, limit)
	default:
		return "", nil, questionBankError.NotLanguageSupportError
	}
}

//switch req.LanguageId {
//	case questionBankEnum.C_LANGUAGE:
//		compile, t, err := cService.Execute(req.FileId, req.Input, req.LanguageLimit)
//		if err != nil {
//			questionBankResp.ErrorHandle(c, err)
//			return
//		}
//		questionBankResp.OkWithDetailed(questionBankResp.Execute{
//			Output:           compile,
//			ExecuteSituation: *t,
//		}, "获取运行结果成功", c)
//	default:
//		questionBankResp.CheckHandle(c, fmt.Errorf("编程语言输入错误"))
//	}
