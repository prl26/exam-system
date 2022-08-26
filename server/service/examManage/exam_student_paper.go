package examManage

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/examManage"
	examManageReq "github.com/prl26/exam-system/server/model/examManage/request"
)

type ExamStudentPaperService struct {
}

// CreateExamStudentPaper 创建ExamStudentPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examstudentPaperService *ExamStudentPaperService) CreateExamStudentPaper(examstudentPaper examManage.ExamStudentPaper) (err error) {
	err = global.GVA_DB.Create(&examstudentPaper).Error
	return err
}

// DeleteExamStudentPaper 删除ExamStudentPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examstudentPaperService *ExamStudentPaperService) DeleteExamStudentPaper(examstudentPaper examManage.ExamStudentPaper) (err error) {
	err = global.GVA_DB.Delete(&examstudentPaper).Error
	return err
}

// DeleteExamStudentPaperByIds 批量删除ExamStudentPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examstudentPaperService *ExamStudentPaperService) DeleteExamStudentPaperByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]examManage.ExamStudentPaper{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateExamStudentPaper 更新ExamStudentPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examstudentPaperService *ExamStudentPaperService) UpdateExamStudentPaper(examstudentPaper examManage.ExamStudentPaper) (err error) {
	err = global.GVA_DB.Save(&examstudentPaper).Error
	return err
}

// GetExamStudentPaper 根据id获取ExamStudentPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examstudentPaperService *ExamStudentPaperService) GetExamStudentPaper(id uint) (examstudentPaper examManage.ExamStudentPaper, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&examstudentPaper).Error
	return
}

// GetExamStudentPaperInfoList 分页获取ExamStudentPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examstudentPaperService *ExamStudentPaperService) GetExamStudentPaperInfoList(info examManageReq.ExamStudentPaperSearch) (list []examManage.ExamStudentPaper, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&examManage.ExamStudentPaper{})
	var examstudentPapers []examManage.ExamStudentPaper
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.PaperId != nil {
		db = db.Where("paper_id = ?", info.PaperId)
	}
	if info.QuestionId != nil {
		db = db.Where("question_id = ?", info.QuestionId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&examstudentPapers).Error
	return examstudentPapers, total, err
}
