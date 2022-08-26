package request

import "github.com/prl26/exam-system/server/model/questionBank"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/26 0:22

 * @Note:	用来创建编程题的请求

 **/

type MakeProgramm struct {
	questionBank.Programm
	SupportLanguage []*SupportLanguage `json:"support_language"`
}

type SupportLanguage struct {
	questionBank.ProgrammLanguageMerge
	Cases []*questionBank.ProgrammCase `json:"cases"`
}
