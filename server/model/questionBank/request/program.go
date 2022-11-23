package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/questionBank"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/28 13:50

 * @Note:

 **/

type ProgramCreate struct {
	questionBank.Programm
	LanguageSupports []*LanguageSupport `json:"languageSupports"`
	LessonSupports   []*LessonSupport   `json:"LessonSupportSupports"`
}

type ProgramFindList struct {
	questionBank.Programm
	request.PageInfo
}

type ProgramDetailEdit struct {
	Id uint `json:"id"`
	questionBank.BasicModel
}

type ProgramCaseFind struct {
	ProgramId  uint `json:"programId"`
	LanguageId int  `json:"languageId"`
}

type ProgramCaseAdd struct {
	ProgramId  uint `json:"programId"`
	LanguageId int  `json:"languageId"`
	Cases      []struct {
		Name   string `json:"name" form:"name" gorm:"column:name;comment:;"`
		Score  uint   `json:"score" form:"score" gorm:"column:score;comment:;"`
		Input  string `json:"input" form:"input" gorm:"column:input;comment:;"`
		Output string `json:"output" form:"output" gorm:"column:output;comment:;"`
		questionBank.ProgrammLimit
	}
}

type ProgramCaseEdit struct {
	Cases []questionBank.ProgrammCase `json:"cases"`
}

type LanguageSupport struct {
	LanguageId      int    `json:"languageId"`
	DefaultCode     string `json:"defaultCode" form:"defaultCode" gorm:"column:default_code;comment:;"`
	ReferenceAnswer string `json:"referenceAnswer" form:"referenceAnswer" gorm:"column:reference_answer;comment:;"`
	Cases           []struct {
		Name   string `json:"name" form:"name" gorm:"column:name;comment:;"`
		Score  uint   `json:"score" form:"score" gorm:"column:score;comment:;"`
		Input  string `json:"input" form:"input" gorm:"column:input;comment:;"`
		Output string `json:"output" form:"output" gorm:"column:output;comment:;"`
		questionBank.ProgrammLimit
	}
}
type LanguageSupportAdd struct {
	ProgramId uint `json:"programId"`
	LanguageSupport
}

type LanguageSupportEdit struct {
	Id              uint
	ProgramId       uint   `json:"programId" form:"programId" gorm:"column:programm_id;comment:;"`
	DefaultCode     string `json:"defaultCode" form:"defaultCode" gorm:"column:default_code;comment:;"`
	ReferenceAnswer string `json:"referenceAnswer" form:"referenceAnswer" gorm:"column:reference_answer;comment:;"`
}

type LanguageSupportDelete struct {
	LanguageIds []int `json:"languageIds"`
	ProgramId   uint  `json:"programId"`
}
