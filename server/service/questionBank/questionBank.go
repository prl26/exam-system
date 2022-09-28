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

type QuestionBankService struct {
	judge          questionBank.Judge
	program        questionBank.Programm
	supplyBlank    questionBank.SupplyBlank
	multipleChoice questionBank.MultipleChoice
	chapterMerge   questionBank.ChapterMerge
}

func (c *QuestionBankService) FindQuestions(chapterId uint) *ojResp.QuestionBank {
	var result ojResp.QuestionBank
	global.GVA_DB.Table(c.judge.TableName()).Joins("join "+c.chapterMerge.TableName()+" on "+c.chapterMerge.TableName()+".question_id="+c.judge.TableName()+".id").Where("chapter_id = ? and question_type = ? and can_practice=?", chapterId, 1, true).Find(&result.Judges)
	global.GVA_DB.Model(&c.program).Joins("join "+c.chapterMerge.TableName()+" on "+c.chapterMerge.TableName()+".question_id="+c.program.TableName()+".id").Where("chapter_id = ? and question_type = ? and can_practice=?", chapterId, 2, true).Preload("LanguageSupports").Find(&result.Programms)
	global.GVA_DB.Model(&c.supplyBlank).Joins("join "+c.chapterMerge.TableName()+" on "+c.chapterMerge.TableName()+".question_id="+c.supplyBlank.TableName()+".id").Where("chapter_id = ? and question_type = ? and can_practice=?", chapterId, 3, true).Find(&result.SupplyBlanks)
	global.GVA_DB.Model(&c.multipleChoice).Joins("join "+c.chapterMerge.TableName()+" on "+c.chapterMerge.TableName()+".question_id="+c.multipleChoice.TableName()+".id").Where("chapter_id = ? and question_type = ? and can_practice=?", chapterId, 4, true).Preload("Options").Find(&result.MultipleChoices)
	return &result
}
