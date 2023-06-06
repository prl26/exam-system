package bo

type OjHelper struct {
	LessonId         uint   `json:"LessonId"`
	ProgramCases     string `json:"programCases"`
	LanguageSupports string `json:"languageSupport"`
}

type Submit struct {
	//ID string `json:"id"`		// 用例ID
	Name         string `json:"name"`  // 用例名称
	Score        uint   `json:"score"` // 用例得分
	ActualOutput string `json:"actualOutput"`
	AnswerOutput string `json:"answerOutput"`
	ExecuteSituation
}

type ExecuteSituation struct {
	ResultStatus uint `json:"resultStatus"` //  结果码
	//ResultStatusStr string `json:"resultStatusStr"` // 结果状态
	ExitStatus int  `json:"exitStatus"` // 程序返回值
	Time       uint `json:"time"`       // 程序运行 CPU 时间，单位纳秒
	Memory     uint `json:"memory"`     // 程序运行内存，单位 byte
	Runtime    uint `json:"runtime"`    // 程序运行现实时间，单位纳秒
}
