package teachplan

import (
	"exam-system/global"
	"exam-system/model/common/request"
	"exam-system/model/teachplan"
	teachplanReq "exam-system/model/teachplan/request"
)

type ScoreService struct {
}

// CreateScore 创建Score记录
// Author [piexlmax](https://github.com/piexlmax)
func (scoreService *ScoreService) CreateScore(score teachplan.Score) (err error) {
	err = global.GVA_DB.Create(&score).Error
	return err
}

// DeleteScore 删除Score记录
// Author [piexlmax](https://github.com/piexlmax)
func (scoreService *ScoreService) DeleteScore(score teachplan.Score) (err error) {
	err = global.GVA_DB.Delete(&score).Error
	return err
}

// DeleteScoreByIds 批量删除Score记录
// Author [piexlmax](https://github.com/piexlmax)
func (scoreService *ScoreService) DeleteScoreByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]teachplan.Score{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateScore 更新Score记录
// Author [piexlmax](https://github.com/piexlmax)
func (scoreService *ScoreService) UpdateScore(score teachplan.Score) (err error) {
	err = global.GVA_DB.Save(&score).Error
	return err
}

// GetScore 根据id获取Score记录
// Author [piexlmax](https://github.com/piexlmax)
func (scoreService *ScoreService) GetScore(id uint) (score teachplan.Score, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&score).Error
	return
}

// GetScoreInfoList 分页获取Score记录
// Author [piexlmax](https://github.com/piexlmax)
func (scoreService *ScoreService) GetScoreInfoList(info teachplanReq.ScoreSearch) (list []teachplan.Score, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&teachplan.Score{})
	var scores []teachplan.Score
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StudentId != nil {
		db = db.Where("student_id = ?", info.StudentId)
	}
	if info.CourseId != nil {
		db = db.Where("course_id = ?", info.CourseId)
	}
	if info.CourseName != "" {
		db = db.Where("course_name LIKE ?", "%"+info.CourseName+"%")
	}
	if info.TeachClassName != "" {
		db = db.Where("teach_class_name LIKE ?", "%"+info.TeachClassName+"%")
	}
	if info.TeachClassId != nil {
		db = db.Where("teach_class_id = ?", info.TeachClassId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&scores).Error
	return scores, total, err
}
