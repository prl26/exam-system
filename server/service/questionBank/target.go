package questionBank

import (
	"fmt"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	"github.com/prl26/exam-system/server/model/questionBank/enum/questionType"
	questionBank "github.com/prl26/exam-system/server/model/questionBank/po"
	questionBankVoResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
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
	if criteria.ChapterId != 0 {
		db = db.Where("chapter_id =?", criteria.ChapterId)
	} else {
		if criteria.LessonId != 0 {
			db = db.Where("lesson_id=?", criteria.LessonId)
		}
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
	//key := fmt.Sprintf("targetPractice:%d:%d", studentId, targetId)
	//err := global.GVA_REDIS.Set(context.Background(), key, address, 2*time.Hour).Err()
	//if err != nil {
	//	global.GVA_LOG.Info(fmt.Sprintf("保存实例地址MYSQL出错了:%s", err))
	//}
	//go func() {
	err2 := global.GVA_DB.Raw("INSERT INTO target_practice(student_id, question_id, question_address) VALUES (?, ?, ?)\nON DUPLICATE KEY UPDATE question_address = VALUES(question_address)", studentId, targetId, address).Scan(nil).Error
	if err2 != nil {
		global.GVA_LOG.Info(fmt.Sprintf("保存实例地址MYSQL出错了:%s", err2))
	}
	//}()
}
func (service *TargetService) ExamRecord(studentId uint, targetId uint, address string, planId uint) {
	//send := studentId % 1000000
	//ddl := send + 22*3600
	//err := global.GVA_REDIS.Set(context.Background(), fmt.Sprintf("targetExam:%d:%d:%d", studentId, targetId, planId), address, time.Duration(ddl)*time.Second).Err()
	//if err != nil {
	//	global.GVA_LOG.Info(fmt.Sprintf("保存实例地址出错了:%s", err))
	//}
	//global.GVA_LOG.Info(fmt.Sprintf("targetExam:%d:%d:%d", studentId, targetId, planId))
	//global.GVA_LOG.Info(fmt.Sprintf("题目实例地址:%s", address))
	err2 := global.GVA_DB.Raw("INSERT INTO target_exam(student_id, question_id, question_address,plan_id) VALUES (?, ?, ?,?)\nON DUPLICATE KEY UPDATE question_address = VALUES(question_address)", studentId, targetId, address, planId).Scan(nil).Error
	if err2 != nil {
		global.GVA_LOG.Info(fmt.Sprintf("保存实例地址MYSQL出错了:%s", err2))
	}
}
func (service *TargetService) QueryExamRecord(studentId uint, targetId uint, planId uint) (string, bool) {
	//address, err := global.GVA_REDIS.Get(context.Background(), fmt.Sprintf("targetExam:%d:%d:%d", studentId, targetId, planId)).Result()
	//global.GVA_LOG.Info(fmt.Sprintf("targetExam:%d:%d:%d answer:%s", studentId, targetId, planId, address))
	//if err != nil {
	//	return "", false
	//}
	//return address, true
	address := ""
	err := global.GVA_DB.Raw("SELECT question_address\nFROM target_exam \nWHERE student_id = ? AND question_id = ? and plan_id =?", studentId, targetId, planId).Scan(&address).Error
	if err != nil || address == "" {
		return address, false
	}
	return address, true
}
func (service *TargetService) QueryPracticeRecord(studentId uint, targetId uint) (string, bool, error) {
	//key := fmt.Sprintf("targetPractice:%d:%d", studentId, targetId)
	//address, err := global.GVA_REDIS.Get(context.Background(), key).Result()
	//if err != nil {
	//	if err == redis.Nil {
	//		return "", false, nil
	//	}
	//	return "", false, err
	//}
	address := ""
	err := global.GVA_DB.Raw("SELECT question_address\nFROM target_practice\nWHERE student_id = ? AND question_id = ?", studentId, targetId).Scan(&address).Error
	if err != nil || address == "" {
		return address, false, err
	}
	return address, true, err
}

func (service *TargetService) QueryHistory(studentId uint, targetId uint) (int, bool) {
	score := new(int)
	if global.GVA_DB.Raw("SELECT score FROM `tea_practice_answer` where student_id=? and question_type=? and question_id=? LIMIT 1", studentId, questionType.Target, targetId).Scan(score).RowsAffected == 0 {
		return 0, false
	}
	return *score, true
}

func (service *TargetService) CreateList(list []*questionBank.Target) (int64, error) {
	tx := global.GVA_DB.Create(&list)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}
