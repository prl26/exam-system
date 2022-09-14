// 自动生成模板ProgrammLanguageMerge
package questionBank

import (
	"github.com/prl26/exam-system/server/global"
)

// ProgrammLanguageMerge 结构体
type ProgrammLanguageMerge struct {
	global.GVA_MODEL
	LanguageId      *int   `json:"languageId" form:"languageId" gorm:"column:language_id;comment:;"`
	ProgrammId      *int   `json:"programmId" form:"programmId" gorm:"column:programm_id;comment:;"`
	DefaultCode     string `json:"defaultCode" form:"defaultCode" gorm:"column:default_code;comment:;"`
	ReferenceAnswer string `json:"referenceAnswer" form:"referenceAnswer" gorm:"column:reference_answer;comment:;"`
}

// TableName ProgrammLanguageMerge 表名
func (ProgrammLanguageMerge) TableName() string {
	return "les_questionBank_programm_language_merge"
}
