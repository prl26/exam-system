package languageType

import (
	questionBankError "github.com/prl26/exam-system/server/model/questionBank/error"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/28 13:10

 * @Note:

 **/

type LanguageType uint

const (
	C_LANGUAGE = LanguageType(1 + iota)
	JAVA
	GO_LANGUAGE
)

const (
	c_language_name    = "c"
	go_language_name   = "go"
	java_language_name = "java"
)

var languageNameToType map[string]LanguageType = map[string]LanguageType{
	c_language_name:    C_LANGUAGE,
	go_language_name:   GO_LANGUAGE,
	java_language_name: JAVA,
}

var languageTypeToName map[LanguageType]string = map[LanguageType]string{
	C_LANGUAGE:  c_language_name,
	GO_LANGUAGE: go_language_name,
	JAVA:        java_language_name,
}

func (t LanguageType) GetLanguageName() (string, error) {
	if v, ok := languageTypeToName[t]; ok {
		return v, nil
	} else {
		return "", questionBankError.NotLanguageSupportError
	}
}

func (t *LanguageType) ToLanguageId(str string) error {
	if v, ok := languageNameToType[str]; ok {
		*t = v
		return nil
	} else {
		return questionBankError.NotLanguageSupportError
	}
}
