package response

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/questionBank"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/28 13:37

 * @Note:

 **/

type Programms struct {
	Programs []*questionBank.Programm
}

type ProgramDetail struct {
	Program         questionBank.Programm `json:"program"`
	CourseSupport   []CourseSupport       `json:"courseSupport"`
	LanguageSupport []LanguageSupport     `json:"languageSupport"`
}

type LanguageSupport struct {
	global.GVA_MODEL
	LanguageId      int    `json:"languageId" gorm:"column:language_id"`
	ReferenceAnswer string `json:"referenceAnswer" gorm:"column:reference_answer"`
	DefaultCode     string `json:"defaultCode" gorm:"column:default_code;comment:;"`
}
