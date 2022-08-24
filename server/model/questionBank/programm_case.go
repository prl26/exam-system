// 自动生成模板QuestionBankProgrammCase
package questionBank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// QuestionBankProgrammCase 结构体
type ProgrammCase struct {
	global.GVA_MODEL
	ProgrammId        string `json:"programm_id" form:"programm_id" gorm:"column:programm_id;comment:;"`
	Name              string `json:"name" form:"name" gorm:"column:name;comment:;"`
	Score             *int   `json:"score" form:"score" gorm:"column:score;comment:;"`
	LanguageId        *int   `json:"languageId" form:"languageId" gorm:"column:language_id;comment:;"`
	CpuLimit          *int   `json:"cpuLimit" form:"cpuLimit" gorm:"column:cpu_limit;comment:;"`
	ClockLimit        *int   `json:"clockLimit" form:"clockLimit" gorm:"column:clock_limit;comment:;"`
	StackLimit        *int   `json:"stackLimit" form:"stackLimit" gorm:"column:stack_limit;comment:;"`
	ProcLimit         *int   `json:"procLimit" form:"procLimit" gorm:"column:proc_limit;comment:;"`
	CpuRateLimit      *int   `json:"cpuRateLimit" form:"cpuRateLimit" gorm:"column:cpu_rate_limit;comment:;"`
	CpuSetLimit       *int   `json:"cpuSetLimit" form:"cpuSetLimit" gorm:"column:cpu_set_limit;comment:;"`
	StrictMemoryLimit *int   `json:"strictMemoryLimit" form:"strictMemoryLimit" gorm:"column:strict_memory_limit;comment:;"`
	Input_type        *int   `json:"input_type" form:"input_type" gorm:"column:input_type;comment:;"`
	Input             string `json:"input" form:"input" gorm:"column:input;comment:;"`
	Out_type          *int   `json:"out_type" form:"out_type" gorm:"column:out_type;comment:;"`
	Output            *int   `json:"output" form:"output" gorm:"column:output;comment:;"`
}

// TableName QuestionBankProgrammCase 表名
func (ProgrammCase) TableName() string {
	return "ProgrammCase"
}
