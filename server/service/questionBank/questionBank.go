package questionBank

import (
	"github.com/prl26/exam-system/server/global"
	ojResp "github.com/prl26/exam-system/server/model/oj/response"
	"github.com/prl26/exam-system/server/model/questionBank"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/26 18:48

 * @Note:	关于题库查询的前台接口

 **/

type QuestionBankService struct{}

func (c *QuestionBankService) FindQuestions(chapterId uint) *ojResp.QuestionBank {
	var result ojResp.QuestionBank
	_ = global.GVA_DB.Where("chapter_id=? and can_practice = ?", chapterId, 1).Model(&questionBank.Judge{}).Find(&result.Judges)
	_ = global.GVA_DB.Where("chapter_id=? and can_practice = ?", chapterId, 1).Model(&questionBank.Programm{}).Preload("LanguageSupports").Find(&result.Programms)
	_ = global.GVA_DB.Where("chapter_id=? and can_practice = ?", chapterId, 1).Model(&questionBank.SupplyBlank{}).Find(&result.SupplyBlanks)
	_ = global.GVA_DB.Where("chapter_id=? and can_practice = ?", chapterId, 1).Model(&questionBank.MultipleChoice{}).Preload("Options").Find(&result.MultipleChoices)
	return &result
}
