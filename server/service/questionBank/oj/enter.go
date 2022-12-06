package oj

import (
	"github.com/prl26/exam-system/server/service/questionBank/oj/judge"
	"github.com/prl26/exam-system/server/service/questionBank/oj/multipleChoice"
	"github.com/prl26/exam-system/server/service/questionBank/oj/program"
	"github.com/prl26/exam-system/server/service/questionBank/oj/supplyBlank"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/24 12:20

 * @Note:

 **/

type OjService struct {
	supplyBlank.SupplyBlankService
	judge.JudgeService
	multipleChoice.MultipleChoiceService
	program.ProgramService
}
