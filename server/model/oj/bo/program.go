package bo

import "github.com/prl26/exam-system/server/model/oj"

type Program struct {
	ProgramCases     string `json:"programCases"`
	LanguageSupports string `json:"languageSupport"`
}

type Submit struct {
	//Id string `json:"id"`		// 用例ID
	Name  string `json:"name"`  // 用例名称
	Score uint   `json:"score"` // 用例得分
	oj.ExecuteSituation
}
