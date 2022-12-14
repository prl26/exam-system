package questionBank

import (
	"context"
	"fmt"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBank "github.com/prl26/exam-system/server/model/questionBank/po"
	questionBankVoResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
	"time"
)

type TargetService struct {
}

func (service *TargetService) Create(RangTopic *questionBank.Target) error {
	return global.GVA_DB.Create(RangTopic).Error
}

func (service *TargetService) DeleteQuestionBankRangTopic(RangTopic questionBank.Target) (err error) {
	err = global.GVA_DB.Delete(&RangTopic).Error
	return err
}

func (service *TargetService) Delete(ids request.IdsReq) error {
	return global.GVA_DB.Delete(&[]questionBank.Target{}, "id in ?", ids.Ids).Error
}

func (service *TargetService) Update(RangTopic questionBank.Target) (err error) {
	err = global.GVA_DB.Updates(&RangTopic).Error
	return err
}

func (service *TargetService) GetQuestionBankRangTopic(id uint) (RangTopic questionBank.Target, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&RangTopic).Error
	return
}

func (service *TargetService) FindTargetList(criteria questionBankBo.TargetSearchCriteria, info request.PageInfo) (list []questionBankVoResp.TargetSimple, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&questionBank.Target{})
	if criteria.LessonId != 0 {
		db = db.Where("lesson_id=?", criteria.LessonId)
	}
	if criteria.IsCheck != nil {
		db = db.Where("is_check=?", criteria.IsCheck)
	}
	if criteria.ProblemType != 0 {
		db = db.Where("problem_type = ?", criteria.ProblemType)
	}
	if criteria.Title != "" {
		db = db.Where("title like ?", "%"+criteria.Title+"%")
	}
	if criteria.CanExam != nil {
		db = db.Where("can_exam = ?", criteria.CanExam)
	}
	if criteria.CanPractice != nil {
		db = db.Where("can_practice = ?", criteria.CanPractice)
	}
	if criteria.ChapterId != 0 {
		db = db.Where("chapter_id =?", criteria.ChapterId)
	}
	if criteria.KnowledgeId != 0 {
		db = db.Where("knowledge_id=?", criteria.KnowledgeId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return list, total, err
}

func (service *TargetService) FindDetail(id uint, more bool) (RangTopic *questionBankBo.TargetDetail, err error) {
	RangTopic = &questionBankBo.TargetDetail{}
	db := global.GVA_DB
	if more {
		db = db.Preload("Chapter").Preload("Knowledge")
	}
	err = db.Model(&questionBank.Target{}).First(RangTopic, id).Error
	return
}

func (service *TargetService) FindTargetPracticeList(knowledge questionBankBo.TargetPracticeCriteria, info request.PageInfo) (list []*questionBankVoResp.TargetSimplePractice, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&questionBank.Target{})
	db = db.Where("is_check=?", 1)
	db = db.Where("can_practice=?", 1)
	if knowledge.KnowledgeId != 0 {
		db = db.Where("knowledge_id=?", knowledge.KnowledgeId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return list, total, err
}

func (service *TargetService) GetByteCode(id uint) *questionBankBo.TargetByteCode {
	code := questionBankBo.TargetByteCode{}
	global.GVA_DB.Model(&questionBank.Target{}).Where("id=?", id).First(&code)
	if code.Id == 0 {
		return nil
	}
	return &code
}

func (service *TargetService) PracticeRecord(studentId uint, targetId uint, address string) {
	global.GVA_REDIS.Set(context.Background(), fmt.Sprintf("targetPractice:%d:%d", studentId, targetId), address, 7*24*time.Hour)
}

func (service *TargetService) QueryPracticeRecord(studentId uint, targetId uint) (string, bool) {
	address, err := global.GVA_REDIS.Get(context.Background(), fmt.Sprintf("targetPractice:%d:%d", studentId, targetId)).Result()
	if err != nil {
		return "", false
	}
	return address, true
}
