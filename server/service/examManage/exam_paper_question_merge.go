package examManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/examManage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    examManageReq "github.com/flipped-aurora/gin-vue-admin/server/model/examManage/request"
)

type PaperQuestionMergeService struct {
}

// CreatePaperQuestionMerge 创建PaperQuestionMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (试卷题目表Service *PaperQuestionMergeService) CreatePaperQuestionMerge(试卷题目表 examManage.PaperQuestionMerge) (err error) {
	err = global.GVA_DB.Create(&试卷题目表).Error
	return err
}

// DeletePaperQuestionMerge 删除PaperQuestionMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (试卷题目表Service *PaperQuestionMergeService)DeletePaperQuestionMerge(试卷题目表 examManage.PaperQuestionMerge) (err error) {
	err = global.GVA_DB.Delete(&试卷题目表).Error
	return err
}

// DeletePaperQuestionMergeByIds 批量删除PaperQuestionMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (试卷题目表Service *PaperQuestionMergeService)DeletePaperQuestionMergeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]examManage.PaperQuestionMerge{},"id in ?",ids.Ids).Error
	return err
}

// UpdatePaperQuestionMerge 更新PaperQuestionMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (试卷题目表Service *PaperQuestionMergeService)UpdatePaperQuestionMerge(试卷题目表 examManage.PaperQuestionMerge) (err error) {
	err = global.GVA_DB.Save(&试卷题目表).Error
	return err
}

// GetPaperQuestionMerge 根据id获取PaperQuestionMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (试卷题目表Service *PaperQuestionMergeService)GetPaperQuestionMerge(id uint) (试卷题目表 examManage.PaperQuestionMerge, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&试卷题目表).Error
	return
}

// GetPaperQuestionMergeInfoList 分页获取PaperQuestionMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (试卷题目表Service *PaperQuestionMergeService)GetPaperQuestionMergeInfoList(info examManageReq.PaperQuestionMergeSearch) (list []examManage.PaperQuestionMerge, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&examManage.PaperQuestionMerge{})
    var 试卷题目表s []examManage.PaperQuestionMerge
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.PaperId != nil {
        db = db.Where("paper_id = ?",info.PaperId)
    }
    if info.QuestionId != nil {
        db = db.Where("question_id = ?",info.QuestionId)
    }
    if info.Score != nil {
        db = db.Where("score = ?",info.Score)
    }
    if info.QuestionType != nil {
        db = db.Where("question_type = ?",info.QuestionType)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&试卷题目表s).Error
	return  试卷题目表s, total, err
}
