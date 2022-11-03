package examManage

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/examManage"
	examManageReq "github.com/prl26/exam-system/server/model/examManage/request"
	"github.com/prl26/exam-system/server/model/examManage/response"
	"github.com/prl26/exam-system/server/model/teachplan"
)

type ExamStudentPaperService struct {
}

// CreateExamStudentPaper 创建ExamStudentPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examstudentPaperService *ExamStudentPaperService) CreateExamStudentPaper(examStudentPaper examManageReq.ExamComing) (studentPaper examManage.ExamStudentPaper, err error) {
	var examPlan teachplan.ExamPlan
	err = global.GVA_DB.Where("id = ?", examStudentPaper.PlanId).Find(&examPlan).Error
	if err != nil {
		return
	}
	var examPaper []examManage.ExamPaper
	err = global.GVA_DB.Where("template_id = ?", examPlan.TemplateId).Find(&examPaper).Error
	if err != nil {
		return
	}
	var paperItemsNum int64
	var AllMerge []response.AllPaperMerge
	for i := 0; i < len(examPaper); i++ {
		var paperMerge []examManage.PaperQuestionMerge
		err = global.GVA_DB.Table("exam_paper_question_merge").Where("paper_id = ?", examPaper[0].ID).Find(&paperMerge).Count(&paperItemsNum).Error
		if err != nil {
			return
		}
		AllMerge[i].PaperMerge = paperMerge
	}
	var studentIds []uint
	var numOfStudents int64
	err = global.GVA_DB.Table("bas_student_teach_classes").Select("student_id").Where("teach_class_id = ?", examPlan.TeachClassId).Find(&studentIds).Count(&numOfStudents).Error
	if err != nil {
		return
	}
	order := 0
	for i := 0; i < int(numOfStudents); i++ {
		for j := 0; j < int(paperItemsNum); j++ {
			studentPaper1 := examManage.ExamStudentPaper{
				GVA_MODEL:    global.GVA_MODEL{},
				PaperId:      &examPaper[order].ID,
				QuestionId:   AllMerge[order].PaperMerge[j].QuestionId,
				StudentId:    AllMerge[order].PaperMerge[j].QuestionId,
				Answer:       "",
				PlanId:       &examPlan.ID,
				Score:        nil,
				QuestionType: AllMerge[order].PaperMerge[j].QuestionType,
				ProblemType:  AllMerge[order].PaperMerge[j].ProblemType,
			}
			global.GVA_DB.Create(&studentPaper1)
			studentPaper = studentPaper1
		}
		order++
		if order == len(examPaper)-1 {
			order = 0
		}
	}
	return
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
	err = global.GVA_DB.Updates(&examstudentPaper).Error
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
