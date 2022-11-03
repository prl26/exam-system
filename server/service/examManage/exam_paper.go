package examManage

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/examManage"
	examManageReq "github.com/prl26/exam-system/server/model/examManage/request"
)

type ExamPaperService struct {
}

// CreateExamPaper 创建ExamPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examStatusServices *ExamPaperService) CreateExamPaper(examPaper examManage.ExamPaper) (err error) {
	err = global.GVA_DB.Create(&examPaper).Error
	return err
}

// DeleteExamPaper 删除ExamPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examStatusServices *ExamPaperService) DeleteExamPaper(examPaper examManage.ExamPaper) (err error) {
	err = global.GVA_DB.Where("id = ?", examPaper.ID).Delete(&examPaper).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Delete(&examManage.PaperQuestionMerge{}, "paper_id = ?", examPaper.ID).Error
	return err
}

// DeleteExamPaperByIds 批量删除ExamPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examStatusServices *ExamPaperService) DeleteExamPaperByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]examManage.ExamPaper{}, "id in ?", ids.Ids).Error
	err = global.GVA_DB.Delete(&examManage.PaperQuestionMerge{}, "paper_id in ?", ids.Ids).Error
	return err
}

// UpdateExamPaper 更新ExamPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examStatusServices *ExamPaperService) UpdateExamPaper(examPaper examManage.ExamPaper) (err error) {
	err = global.GVA_DB.Where("id = ?", examPaper.ID).Updates(&examPaper).Error
	return err
}

// GetExamPaper 根据id获取ExamPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examStatusServices *ExamPaperService) GetExamPaper(id uint) (examPaper examManage.ExamPaper, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&examPaper).Error

	return
}

// GetExamPaperInfoList 分页获取ExamPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examStatusServices *ExamPaperService) GetExamPaperInfoList(info examManageReq.ExamPaperSearch) (list []examManage.ExamPaper, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&examManage.ExamPaper{})
	var examPapers []examManage.ExamPaper
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.PlanId != nil {
		db = db.Where("plan_id = ?", info.PlanId)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.TemplateId != nil {
		db = db.Where("template_id = ?", info.TemplateId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&examPapers).Error
	return examPapers, total, err
}
func (examStatusServices *ExamPaperService) GetTemplate(info examManage.ExamPaper) (list []examManage.PaperTemplateItem, err error) {
	err = global.GVA_DB.Where("template_id = ?", info.TemplateId).Find(&list).Error
	return
}
