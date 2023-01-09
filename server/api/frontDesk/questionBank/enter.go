package questionBank

import "github.com/prl26/exam-system/server/service"

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/30 21:59

 * @Note:

 **/

type ApiGroup struct {
	QuestionBankApi
	OjApi
	TargetApi
}

var (
	questionBankService = service.ServiceGroupApp.QuestionBankServiceGroup.QuestionBankService
)
