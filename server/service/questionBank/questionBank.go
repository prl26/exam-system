package questionBank

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBank "github.com/prl26/exam-system/server/model/questionBank/po"
	questionBankVoResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
)

//
//import (
//	"fmt"
//	"github.com/prl26/exam-system/server/global"
//	"github.com/prl26/exam-system/server/model/enum/questionType"
//	ojResp "github.com/prl26/exam-system/server/model/oj/response"
//	"github.com/prl26/exam-system/server/model/questionBank"
//	"github.com/prl26/exam-system/server/model/questionBank/response"
//	questionBankResp "github.com/prl26/exam-system/server/model/questionBank/response"
//	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/vo/request"
//)
//
///**
//
// * @Author: AloneAtWar
//
// * @Date:   2022/8/26 18:48
//
// * @Note:	关于题库查询的前台接口
//
// **/
//

type QuestionBankService struct {
	judge          questionBank.Judge
	program        questionBank.Program
	supplyBlank    questionBank.SupplyBlank
	multipleChoice questionBank.MultipleChoice
}

//

func (service *QuestionBankService) FindJudgeList(criteria questionBankBo.JudgePracticeCriteria, info request.PageInfo) (list []questionBankBo.PracticeModel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&questionBank.Judge{})
	db = db.Where("is_check", 1)
	db = db.Where("can_practice = ?", 1)
	db = db.Where("chapter_id =?", criteria.ChapterId)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Model(&questionBank.Judge{}).Limit(limit).Offset(offset).Find(&list).Error
	return list, total, err
}

func (service *QuestionBankService) FindSupplyBlankList(criteria questionBankBo.SupplyBlankPracticeCriteria, info request.PageInfo) (list []questionBankVoResp.SupplyBlankPractice, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&questionBank.SupplyBlank{})
	db = db.Where("is_check", 1)
	db = db.Where("can_practice = ?", 1)
	db = db.Where("chapter_id =?", criteria.ChapterId)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return list, total, err
}

func (service *QuestionBankService) FindMultipleChoiceList(criteria questionBankBo.MultiplePracticeCriteria, info request.PageInfo, isIndefinite bool) (list []questionBankVoResp.MultipleChoicePractice, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&questionBank.MultipleChoice{})
	db = db.Where("is_check", 1)
	if isIndefinite {
		db = db.Where("is_indefinite=1")
	} else {
		db = db.Where("is_indefinite=0")
	}
	db = db.Where("can_practice = ?", 1)
	db = db.Where("chapter_id =?", criteria.ChapterId)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return list, total, err
}

func (service *QuestionBankService) FindProgramList(criteria questionBankBo.ProgramPracticeCriteria, info request.PageInfo) (list []*questionBankVoResp.ProgramPractice, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&questionBank.Program{})
	db = db.Where("is_check", 1)
	db = db.Where("can_practice = ?", 1)
	db = db.Where("chapter_id =?", criteria.ChapterId)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var l []*questionBankBo.ProgramPractice
	err = db.Limit(limit).Offset(offset).Find(&l).Error
	list = make([]*questionBankVoResp.ProgramPractice, len(l))
	for i, practice := range l {
		list[i] = new(questionBankVoResp.ProgramPractice)
		list[i].Convert(practice)
	}
	return list, total, err
}

//func (c *QuestionBankService) FindQuestions(chapterId uint) *ojResp.QuestionBank {
//	var result ojResp.QuestionBank
//	global.GVA_DB.Model(&c.judge).Joins("join "+c.chapterMerge.TableName()+" on "+c.chapterMerge.TableName()+".question_id="+c.judge.TableName()+".id").Where("knowledge_id = ? and question_type = ? and can_practice=?", chapterId, questionType.JUDGE, true).Find(&result.Judges)
//	global.GVA_DB.Model(&c.program).Joins("join "+c.chapterMerge.TableName()+" on "+c.chapterMerge.TableName()+".question_id="+c.program.TableName()+".id").Where("knowledge_id = ? and question_type = ? and can_practice=?", chapterId, questionType.PROGRAM, true).Preload("LanguageSupports").Find(&result.Programms)
//	global.GVA_DB.Model(&c.supplyBlank).Joins("join "+c.chapterMerge.TableName()+" on "+c.chapterMerge.TableName()+".question_id="+c.supplyBlank.TableName()+".id").Where("knowledge_id = ? and question_type = ? and can_practice=?", chapterId, questionType.SUPPLY_BLANK, true).Find(&result.SupplyBlanks)
//	global.GVA_DB.Model(&c.multipleChoice).Joins("join "+c.chapterMerge.TableName()+" on "+c.chapterMerge.TableName()+".question_id="+c.multipleChoice.TableName()+".id").Where("knowledge_id = ? and question_type = ? and can_practice=?", chapterId, questionType.SINGLE_CHOICE, true).Find(&result.MultipleChoices)
//	return &result
//}
//func (c QuestionBankService) FindJudges(chapterId uint) (result []*ojResp.ApiJudge) {
//	global.GVA_DB.Model(&c.judge).Joins("join "+c.chapterMerge.TableName()+" on "+c.chapterMerge.TableName()+".question_id="+c.judge.TableName()+".id").Where("knowledge_id = ? and question_type = ? and can_practice=?", chapterId, questionType.JUDGE, true).Find(&result)
//	return
//}
//
//func (c QuestionBankService) FindProgramms(chapterId uint) (result []*ojResp.ApiProgramm) {
//	global.GVA_DB.Model(&c.program).Joins("join "+c.chapterMerge.TableName()+" on "+c.chapterMerge.TableName()+".question_id="+c.program.TableName()+".id").Where("knowledge_id = ? and question_type = ? and can_practice=?", chapterId, questionType.PROGRAM, true).Preload("LanguageSupports").Find(&result)
//	return
//}
//
//func (c QuestionBankService) FindSupplyBlank(chapterId uint) (result []*ojResp.ApiSupplyBlank) {
//	global.GVA_DB.Model(&c.supplyBlank).Joins("join "+c.chapterMerge.TableName()+" on "+c.chapterMerge.TableName()+".question_id="+c.supplyBlank.TableName()+".id").Where("knowledge_id = ? and question_type = ? and can_practice=?", chapterId, questionType.SUPPLY_BLANK, true).Find(&result)
//	return
//}
//
//func (c QuestionBankService) FindChoice(chapterId uint, isMultiple bool) (result []*ojResp.ApiJudge) {
//	db := global.GVA_DB.Model(&c.multipleChoice).Joins("join " + c.chapterMerge.TableName() + " on " + c.chapterMerge.TableName() + ".question_id=" + c.multipleChoice.TableName() + ".id")
//	if isMultiple {
//		db.Where("most_options>1")
//	} else {
//		db.Where("most_options=1")
//	}
//	db.Where("knowledge_id = ? and question_type = ? and can_practice=? ", chapterId, questionType.SINGLE_CHOICE, true).Find(&result)
//	return
//}
//
//var languageSupport = `
//	select *
//	from les_questionbank_programm_language_merge a
//	where a.programm_id= ?
//	and ISNULL(a.deleted_at)
//`
//
//func (c *QuestionBankService) FindCourseSupport(support *[]response.CourseSupport, questionId uint, questionType0 int) error {
//	if err := global.GVA_DB.Raw(lessonSupportSql, questionId, questionType0).Find(support).Error; err != nil {
//		return fmt.Errorf("无法找到课程支持")
//	}
//	return nil
//}
//
//func (c *QuestionBankService) DeleteCourseSupport(ids []int) error {
//	return global.GVA_DB.delete(&[]questionBank.ChapterMerge{}, ids).Error
//}
//
//func (c *QuestionBankService) AddCourseSupport(merges []questionBank.ChapterMerge) error {
//	return global.GVA_DB.Create(&merges).Error
//}
//
//func (c *QuestionBankService) FindQuestionSupport(req questionBankReq.QuestionsSupportFind) (list []questionBankResp.QuestionSupport, total int64, err error) {
//	limit := req.PageSize
//	offset := req.PageSize * (req.Page - 1)
//	db := global.GVA_DB
//	switch req.QuestionType {
//	case questionType.PROGRAM:
//		db = db.Model(&c.program).Joins("join "+c.chapterMerge.TableName()+" on "+c.chapterMerge.TableName()+".question_id="+c.program.TableName()+".id").Where("chapter_id = ? and question_type = ?", req.ChapterId, questionType.PROGRAM)
//	case questionType.SINGLE_CHOICE:
//		db = db.Model(&c.multipleChoice).Joins("join "+c.chapterMerge.TableName()+" on "+c.chapterMerge.TableName()+".question_id="+c.multipleChoice.TableName()+".id").Where("chapter_id = ? and question_type = ?", req.ChapterId, questionType.SINGLE_CHOICE)
//	case questionType.SUPPLY_BLANK:
//		db = db.Model(&c.supplyBlank).Joins("join "+c.chapterMerge.TableName()+" on "+c.chapterMerge.TableName()+".question_id="+c.supplyBlank.TableName()+".id").Where("chapter_id = ? and question_type = ?", req.ChapterId, questionType.SUPPLY_BLANK)
//	case questionType.JUDGE:
//		db = db.Model(&c.judge).Joins("join "+c.chapterMerge.TableName()+" on "+c.chapterMerge.TableName()+".question_id="+c.judge.TableName()+".id").Where("chapter_id = ? and question_type = ?", req.ChapterId, questionType.JUDGE)
//	default:
//		return nil, 0, fmt.Errorf("unknown question type")
//	}
//	if req.ProblemType != 0 {
//		db.Where("problem_type=?", req.ProblemType)
//	}
//	if req.CanPractice != nil {
//		db.Where("can_practice=?", req.CanPractice)
//	}
//	if req.Title != "" {
//		db.Where("title Like ?", "%"+req.Title+"%")
//	}
//	err = db.QuestionCount(&total).Error
//	if err != nil {
//		return
//	}
//	err = db.Limit(limit).Offset(offset).Find(&list).Error
//	return
//}
//
//
