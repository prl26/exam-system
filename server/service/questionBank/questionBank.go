package questionBank

import (
	"fmt"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/enum/questionType"
	ojResp "github.com/prl26/exam-system/server/model/oj/response"
	"github.com/prl26/exam-system/server/model/questionBank"
	"github.com/prl26/exam-system/server/model/questionBank/response"
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
	global.GVA_DB.Table(c.judge.TableName()).Joins("join "+c.chapterMerge.TableName()+" on "+c.chapterMerge.TableName()+".question_id="+c.judge.TableName()+".id").Where("chapter_id = ? and question_type = ? and can_practice=?", chapterId, questionType.JUDGE, true).Find(&result.Judges)
	global.GVA_DB.Model(&c.program).Joins("join "+c.chapterMerge.TableName()+" on "+c.chapterMerge.TableName()+".question_id="+c.program.TableName()+".id").Where("chapter_id = ? and question_type = ? and can_practice=?", chapterId, questionType.PROGRAM, true).Preload("LanguageSupports").Find(&result.Programms)
	global.GVA_DB.Model(&c.supplyBlank).Joins("join "+c.chapterMerge.TableName()+" on "+c.chapterMerge.TableName()+".question_id="+c.supplyBlank.TableName()+".id").Where("chapter_id = ? and question_type = ? and can_practice=?", chapterId, questionType.SUPPLY_BLANK, true).Find(&result.SupplyBlanks)
	global.GVA_DB.Model(&c.multipleChoice).Joins("join "+c.chapterMerge.TableName()+" on "+c.chapterMerge.TableName()+".question_id="+c.multipleChoice.TableName()+".id").Where("chapter_id = ? and question_type = ? and can_practice=?", chapterId, questionType.MULTIPLE_CHOICE, true).Preload("Options").Find(&result.MultipleChoices)
	return &result
}

var languageSupport = `
	select *
	from les_questionbank_programm_language_merge a
	where a.programm_id= ?
	and ISNULL(a.deleted_at)
`

func (c *QuestionBankService) FindCourseSupport(support *[]response.CourseSupport, questionId uint, questionType0 int) error {
	if err := global.GVA_DB.Raw(lessonSupportSql, questionId, questionType.PROGRAM).Find(support).Error; err != nil {
		return fmt.Errorf("无法找到课程支持")
	}
	return nil
}
