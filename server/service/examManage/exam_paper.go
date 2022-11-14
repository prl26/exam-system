package examManage

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/examManage"
	examManageReq "github.com/prl26/exam-system/server/model/examManage/request"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type ExamPaperService struct {
}

// CreateExamPaper 创建ExamPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPaperService *ExamPaperService) CreateExamPaper(examPaper examManage.ExamPaper) (err error) {
	err = global.GVA_DB.Create(&examPaper).Error
	return err
}

// DeleteExamPaper 删除ExamPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPaperService *ExamPaperService) DeleteExamPaper(examPaper examManage.ExamPaper) (err error) {
	err = global.GVA_DB.Where("id = ?", examPaper.ID).Delete(&examPaper).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Delete(&examManage.PaperQuestionMerge{}, "paper_id = ?", examPaper.ID).Error
	return err
}

// DeleteExamPaperByIds 批量删除ExamPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPaperService *ExamPaperService) DeleteExamPaperByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]examManage.ExamPaper{}, "id in ?", ids.Ids).Error
	err = global.GVA_DB.Delete(&examManage.PaperQuestionMerge{}, "paper_id in ?", ids.Ids).Error
	return err
}

// UpdateExamPaper 更新ExamPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPaperService *ExamPaperService) UpdateExamPaper(examPaper examManage.ExamPaper) (err error) {
	err = global.GVA_DB.Where("id = ?", examPaper.ID).Updates(&examPaper).Error
	return err
}

// GetExamPaper 根据id获取ExamPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPaperService *ExamPaperService) GetExamPaper(id uint) (examPaper examManage.ExamPaper, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&examPaper).Error

	return
}

// GetExamPaperInfoList 分页获取ExamPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPaperService *ExamPaperService) GetExamPaperInfoList(info examManageReq.ExamPaperSearch) (list []examManage.ExamPaper, total int64, err error) {
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
func (examPaperService *ExamPaperService) GetTemplate(info examManage.ExamPaper) (list []examManage.PaperTemplateItem, err error) {
	err = global.GVA_DB.Where("template_id = ?", info.TemplateId).Find(&list).Error
	return
}
func (examPaperService *ExamPaperService) PaperDistribution(PlanId uint) (err error) {
	var number int64
	global.GVA_DB.Table("exam_paper").Where("plan_id = ?", PlanId).Count(&number)
	var studentList []int64
	global.GVA_DB.Raw("SELECT student_id FROM bas_student_teach_classes join tea_examplan on  tea_examplan.teach_class_id = bas_student_teach_classes.teach_class_id and bas_student_teach_classes.teach_class_id = ? GROUP BY student_id ", PlanId).
		Scan(&studentList)
	rand.Seed(time.Now().UnixNano())
	for _, v := range studentList {
		go func() {
			a := rand.Intn(int(number))
			global.GVA_DB.Transaction(func(tx *gorm.DB) error {
				err = tx.Table("exam_student_paper").Raw("INSERT INTO exam_student_paper(question_id,score,question_type,problem_type,paper_id) SELECT question_id,score,question_type,problem_type,paper_id from exam_paper_question_merge WHERE paper_id = ?", a).Error
				if err != nil {
					return err
				}
				err = tx.Table("exam_student_paper").Raw("UPDATE exam_student_paper SET id = ?,plan_id =? ", v, PlanId).Where("paper_id = ?", a).Error
				if err != nil {
					return err
				}
				return nil
			})
		}()
	}
	return
}
