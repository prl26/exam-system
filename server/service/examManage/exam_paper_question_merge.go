package examManage

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/examManage"
	examManageReq "github.com/prl26/exam-system/server/model/examManage/request"
)

type PaperQuestionMergeService struct {
}

// CreatePaperQuestionMerge 创建PaperQuestionMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (paperQuestionMergeService *PaperQuestionMergeService) CreatePaperQuestionMerge(paperQuestionMerge examManage.PaperQuestionMerge) (err error) {
	err = global.GVA_DB.Create(&paperQuestionMerge).Error
	return err
}

// DeletePaperQuestionMerge 删除PaperQuestionMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (paperQuestionMergeService *PaperQuestionMergeService) DeletePaperQuestionMerge(paperQuestionMerge examManage.PaperQuestionMerge) (err error) {
	err = global.GVA_DB.Delete(&paperQuestionMerge).Error
	return err
}

// DeletePaperQuestionMergeByIds 批量删除PaperQuestionMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (paperQuestionMergeService *PaperQuestionMergeService) DeletePaperQuestionMergeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]examManage.PaperQuestionMerge{}, "id in ?", ids.Ids).Error
	return err
}

// UpdatePaperQuestionMerge 更新PaperQuestionMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (paperQuestionMergeService *PaperQuestionMergeService) UpdatePaperQuestionMerge(paperQuestionMerge examManage.PaperQuestionMerge) (err error) {
	err = global.GVA_DB.Updates(&paperQuestionMerge).Error
	return err
}

// GetPaperQuestionMerge 根据id获取PaperQuestionMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (paperQuestionMergeService *PaperQuestionMergeService) GetPaperQuestionMerge(id uint) (paperQuestionMerge examManage.PaperQuestionMerge, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&paperQuestionMerge).Error
	return
}

// GetPaperQuestionMergeInfoList 分页获取PaperQuestionMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (paperQuestionMergeService *PaperQuestionMergeService) GetPaperQuestionMergeInfoList(info examManageReq.PaperQuestionMergeSearch) (list []examManage.PaperQuestionMerge, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&examManage.PaperQuestionMerge{})
	var paperQuestionMerges []examManage.PaperQuestionMerge
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.PaperId != nil {
		db = db.Where("paper_id = ?", info.PaperId)
	}
	if info.QuestionId != nil {
		db = db.Where("question_id = ?", info.QuestionId)
	}
	if info.Score != nil {
		db = db.Where("score = ?", info.Score)
	}
	if info.QuestionType != nil {
		db = db.Where("question_type = ?", info.QuestionType)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&paperQuestionMerges).Error
	return paperQuestionMerges, total, err
}

//func (paperQuestionMergeService *PaperQuestionMergeService) SetChoiceQuestionMerge(paper frontExamManage.ExamPaper, info []questionBank.MultipleChoice) (err error) {
//	for i := 0; i < len(info); i++ {
//		questionMerge := frontExamManage.PaperQuestionMerge{
//			GVA_MODEL:    global.GVA_MODEL{},
//			PaperId:      paper.ID,
//			QuestionId:   &info[i].ID,
//			Score:        0.,
//			QuestionType: nil,
//			ProblemType:  nil,
//		}
//	}
//}
