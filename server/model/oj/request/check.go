package request

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/26 19:49

 * @Note:

 **/

type CheckSupplyBlank struct {
	Id      uint     `json:"id"`
	Answers []string `json:"answers"`
}

type CheckMultipleChoice struct {
	Id      uint     `json:"id"`
	Answers []string `json:"answers"`
}

type CheckJudge struct {
	Id     uint `json:"id"`
	Answer bool `json:"answers"`
}

type CheckProgramm struct {
	Id         uint   `json:"id"`
	Code       string `json:"code"`
	LanguageId uint   `json:"language_id"`
}
