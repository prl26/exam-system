package response

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/26 16:52

 * @Note:	获取题库的接口返回

 **/

type QuestionBank struct {
	MultipleChoices []*ApiMultipleChoice `json:"multipleChoices"`
	SupplyBlanks    []*ApiSupplyBlank    `json:"supplyBlanks"`
	Programms       []*ApiProgramm       `json:"programms"`
	Judges          []*ApiJudge          `json:"judges"`
}

type Base struct {
	Id          uint   `json:"id"`
	ProblemType *int   `json:"problemType" form:"problemType" gorm:"column:problem_type;comment:试卷难度;"`
	Describe    string `json:"describe" form:"describe" gorm:"column:describe;comment:;"`
	Title       string `json:"title" form:"title" gorm:"column:title;comment:;"`
}

type ApiJudge struct {
	Base
}

type ApiSupplyBlank struct {
	Base
	Num *int `json:"num" form:"num" gorm:"column:num;comment:可填项;"`
}

func (ApiLanguageSupport) TableName() string {
	return "les_questionBank_programm_language_merge"
}

type ApiProgramm struct {
	Base
	LanguageSupports []ApiLanguageSupport `json:"languageSupports" gorm:"foreignKey:ProgrammId"`
}

type ApiOption struct {
	Describe         string `json:"describe" form:"describe" gorm:"column:describe;comment:描述;"`
	Orders           *int   `json:"orders" form:"Order" gorm:"orders"`
	MultipleChoiceId uint   `json:"multipleChoiceId" form:"multipleChoiceId" gorm:"column:multiple_choice_id;comment:选择题id;"`
}

type ApiLanguageSupport struct {
	LanguageId  *int   `json:"languageId" form:"languageId" gorm:"column:language_id;comment:;"`
	ProgrammId  *int   `json:"programmId" form:"programmId" gorm:"column:programm_id;comment:;"`
	DefaultCode string `json:"defaultCode" form:"defaultCode" gorm:"column:default_code;comment:;"`
}

func (ApiOption) TableName() string {
	return "les_questionBank_options"
}

type ApiMultipleChoice struct {
	Base
	MostOptions *int `json:"mostOptions" form:"mostOptions" gorm:"column:most_options;comment:最多可选项;"`
}
