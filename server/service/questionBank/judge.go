package questionBank

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/enum/questionType"
	"github.com/prl26/exam-system/server/model/questionBank"
	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/request"
	"gorm.io/gorm"
)

type JudgeService struct {
}

// Create 创建QuestionBankJudge记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *JudgeService) Create(judge *questionBank.Judge, lessonSupports []*questionBankReq.LessonSupport) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(judge).Error; err != nil {
			return err
		}
		courseSupport := buildCourseSupport(lessonSupports, judge.ID, questionType.JUDGE)
		if err := tx.Create(&courseSupport).Error; err != nil {
			return err
		}
		return nil
	})
}

// DeleteQuestionBankJudge 删除QuestionBankJudge记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *JudgeService) DeleteQuestionBankJudge(judge questionBank.Judge) (err error) {
	err = global.GVA_DB.Delete(&judge).Error
	return err
}

// Delete 批量删除QuestionBankJudge记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *JudgeService) Delete(ids request.IdsReq) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&[]questionBank.Judge{}, "id in ?", ids.Ids).Error; err != nil {
			return err
		}
		if err := tx.Delete(&[]questionBank.ChapterMerge{}, "id in ? and question_type=?", ids.Ids, questionType.JUDGE).Error; err != nil {
			return err
		}
		return nil
	})
}

// Update 更新QuestionBankJudge记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *JudgeService) Update(judge questionBank.Judge) (err error) {
	err = global.GVA_DB.Save(&judge).Error
	return err
}

// GetQuestionBankJudge 根据id获取QuestionBankJudge记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *JudgeService) GetQuestionBankJudge(id uint) (judge questionBank.Judge, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&judge).Error
	return
}

// FindJudgeList 分页获取QuestionBankJudge记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *JudgeService) FindJudgeList(info questionBankReq.QuestionBankJudgeSearch) (list []questionBank.JudgeView, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&questionBank.Judge{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ProblemType != 0 {
		db = db.Where("problem_type = ?", info.ProblemType)
	}
	if info.Title != "" {
		db = db.Where("title like ?", "%"+info.Title+"%")
	}
	if info.CanExam != nil {
		db = db.Where("can_exam = ?", info.CanExam)
	}
	if info.CanPractice != nil {
		db = db.Where("can_practice = ?", info.CanPractice)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return list, total, err
}

func (service *JudgeService) FindDetail(judge *questionBank.Judge, id uint) error {
	return global.GVA_DB.Where("id=?", id).Find(judge).Error
}
