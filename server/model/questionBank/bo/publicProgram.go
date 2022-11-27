package bo

import (
	"encoding/json"
	"github.com/prl26/exam-system/server/model/questionBank/enum"
	questionBankError "github.com/prl26/exam-system/server/model/questionBank/error"
	questionBankPo "github.com/prl26/exam-system/server/model/questionBank/po"
)

type LanguageSupport struct {
	LanguageId enum.LanguageType `json:"languageId" form:"languageId" gorm:"column:language_id;comment:;"`
	LanguageLimit
}

type LanguageLimit struct {
	StrictMemoryLimit *int `json:"strictMemoryLimit,omitempty" form:"strictMemoryLimit" gorm:"column:strict_memory_limit;comment:;"`
	MemoryLimit       *int `json:"memoryLimit,omitempty" form:"memoryLimit" gorm:"column:memory_limit"`
	CpuLimit          *int `json:"cpuLimit,omitempty" form:"cpuLimit" gorm:"column:cpu_limit;comment:;"`
	ClockLimit        *int `json:"clockLimit,omitempty" form:"clockLimit" gorm:"column:clock_limit;comment:;"`
	StackLimit        *int `json:"stackLimit,omitempty" form:"stackLimit" gorm:"column:stack_limit;comment:;"`
	ProcLimit         *int `json:"procLimit,omitempty" form:"procLimit" gorm:"column:proc_limit;comment:;"`
	CpuRateLimit      *int `json:"cpuRateLimit,omitempty" form:"cpuRateLimit" gorm:"column:cpu_rate_limit;comment:;"`
	CpuSetLimit       *int `json:"cpuSetLimit,omitempty" form:"cpuSetLimit" gorm:"column:cpu_set_limit;comment:;"`
}

type ProgramCase struct {
	Name   string `json:"name" form:"name" gorm:"column:name;comment:;"`
	Score  uint   `json:"score" form:"score" gorm:"column:score;comment:;"`
	Input  string `json:"input" form:"input" gorm:"column:input;comment:;"`
	Output string `json:"output" form:"output" gorm:"column:output;comment:;"`
}

type DefaultCode struct {
	LanguageId enum.LanguageType `json:"languageId" form:"languageId" gorm:"column:language_id;comment:;"`
	Code       string            `json:"code"`
}

type ReferenceAnswer struct {
	LanguageId enum.LanguageType `json:"languageId" form:"languageId" gorm:"column:language_id;comment:;"`
	Code       string            `json:"code"`
}

type LanguageSupports []*LanguageSupport
type ProgramCases []*ProgramCase
type DefaultCodes []*DefaultCode
type ReferenceAnswers []*ReferenceAnswer

func (s *ProgramCases) Serialize() (string, error) {
	jsons, err := json.Marshal(s)
	var sum uint
	for _, programCase := range *s {
		sum += programCase.Score
	}
	if sum != 100 {
		return "", questionBankError.ScoreError
	}
	return string(jsons), err
}

func (s *ProgramCases) Deserialize(str string) error {
	err := json.Unmarshal([]byte(str), &s)
	return err
}

func (s *LanguageSupports) Serialize() (string, error) {
	table := make(map[string]*LanguageLimit)
	for _, support := range *s {
		name, err := support.LanguageId.GetLanguageName()
		if err != nil {
			return "", err
		}
		table[name] = &support.LanguageLimit
	}
	jsons, err := json.Marshal(table)
	return string(jsons), err
}

func (s *DefaultCodes) Serialize() (string, error) {
	table := make(map[string]string)
	for _, support := range *s {
		name, err := support.LanguageId.GetLanguageName()
		if err != nil {
			return "", err
		}
		table[name] = support.Code
	}
	jsons, err := json.Marshal(table)
	return string(jsons), err
}

func (s *DefaultCodes) Deserialization(str string) error {
	table := make(map[string]string)
	err := json.Unmarshal([]byte(str), &table)
	if err != nil {
		return err
	}
	*s = make([]*DefaultCode, len(table))
	i := 0
	for k, support := range table {
		(*s)[i] = &DefaultCode{}
		err := (*s)[i].LanguageId.ToLanguageId(k)
		if err != nil {
			return err
		}
		(*s)[i].Code = support
	}
	return nil
}

func (s *ReferenceAnswers) Serialize() (string, error) {
	table := make(map[string]string)
	for _, support := range *s {
		name, err := support.LanguageId.GetLanguageName()
		if err != nil {
			return "", err
		}
		table[name] = support.Code
	}
	jsons, err := json.Marshal(table)
	return string(jsons), err
}

func (s *ReferenceAnswers) Deserialization(str string) error {
	table := make(map[string]string)
	err := json.Unmarshal([]byte(str), &table)
	if err != nil {
		return err
	}
	*s = make([]*ReferenceAnswer, len(table))
	i := 0
	for k, support := range table {
		(*s)[i] = &ReferenceAnswer{}
		err := (*s)[i].LanguageId.ToLanguageId(k)
		if err != nil {
			return err
		}
		(*s)[i].Code = support
	}
	return nil
}
func (s *LanguageSupports) Deserialization(str string) error {
	table := make(map[string]*LanguageLimit)
	err := json.Unmarshal([]byte(str), &table)
	if err != nil {
		return err
	}
	*s = make([]*LanguageSupport, len(table))
	i := 0
	for k, support := range table {
		(*s)[i] = &LanguageSupport{}
		err := (*s)[i].LanguageId.ToLanguageId(k)
		if err != nil {
			return err
		}
		(*s)[i].LanguageLimit = *support
	}
	return nil
}

type PublicProgramSearchCriteria struct {
	questionBankPo.SimpleModel
}
