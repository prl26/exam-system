package system

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	basicdataReq "github.com/prl26/exam-system/server/model/basicdata/request"
)

type FrontSystemService struct {
}

func (frontSystemService *FrontSystemService) GetTermInfoList(info basicdataReq.FrontTermSearch) (list []basicdata.Term, total int64, err error) {

	// 创建db
	db := global.GVA_DB.Model(&basicdata.Term{})
	var terms []basicdata.Term
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Find(&terms).Error
	return terms, total, err
}
func (frontSystemService *FrontSystemService) GetLessonInfoList(info basicdataReq.FrontLessonSearch) (list []basicdata.Lesson, total int64, err error) {

	// 创建db
	db := global.GVA_DB.Model(&basicdata.Lesson{})
	var lessons []basicdata.Lesson
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Find(&lessons).Error
	return lessons, total, err
}
