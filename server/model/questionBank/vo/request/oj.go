package request

import (
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBankEnum "github.com/prl26/exam-system/server/model/questionBank/enum"
)

type Compile struct {
	Code       string                        `json:"code"`
	LanguageId questionBankEnum.LanguageType `json:"languageId"`
}

type Execute struct {
	FileId                       string                        `json:"fileId"`     //文件ID
	LanguageId                   questionBankEnum.LanguageType `json:"languageId"` //此是用于一些特殊语言的编译运行，例如 PY
	Input                        string                        `json:"input"`      //文件标准输入
	questionBankBo.LanguageLimit                               //各种限制
}
