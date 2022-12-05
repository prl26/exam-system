package oj

import (
	"github.com/prl26/exam-system/server/service/oj/judge"
	"github.com/prl26/exam-system/server/service/oj/multipleChoice"
	"github.com/prl26/exam-system/server/service/oj/program"
	"github.com/prl26/exam-system/server/service/oj/program/common"
	"github.com/prl26/exam-system/server/service/oj/supplyBlank"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/24 12:20

 * @Note:

 **/

type ServiceGroup struct {
	supplyBlank.SupplyBlankService
	judge.JudgeService
	multipleChoice.MultipleChoiceService
	program.ProgramService
	common.CommonService
}
