package lessondata

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/lessondata"
	lessondataReq "github.com/prl26/exam-system/server/model/lessondata/request"
)

type ResourcesTestService struct {
}

// CreateResourcesTest 创建ResourcesTest记录
// Author [piexlmax](https://github.com/piexlmax)
func (resourcesTestService *ResourcesTestService) CreateResourcesTest(resourcesTest lessondata.ResourcesTest) (err error) {
	err = global.GVA_DB.Create(&resourcesTest).Error
	return err
}

// DeleteResourcesTest 删除ResourcesTest记录
// Author [piexlmax](https://github.com/piexlmax)
func (resourcesTestService *ResourcesTestService) DeleteResourcesTest(resourcesTest lessondata.ResourcesTest) (err error) {
	err = global.GVA_DB.Delete(&resourcesTest).Error
	return err
}

// DeleteResourcesTestByIds 批量删除ResourcesTest记录
// Author [piexlmax](https://github.com/piexlmax)
func (resourcesTestService *ResourcesTestService) DeleteResourcesTestByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]lessondata.ResourcesTest{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateResourcesTest 更新ResourcesTest记录
// Author [piexlmax](https://github.com/piexlmax)
func (resourcesTestService *ResourcesTestService) UpdateResourcesTest(resourcesTest lessondata.ResourcesTest) (err error) {
	err = global.GVA_DB.Updates(&resourcesTest).Error
	return err
}

// GetResourcesTest 根据id获取ResourcesTest记录
// Author [piexlmax](https://github.com/piexlmax)
func (resourcesTestService *ResourcesTestService) GetResourcesTest(id uint) (resourcesTest lessondata.ResourcesTest, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&resourcesTest).Error
	return
}

// GetResourcesTestInfoList 分页获取ResourcesTest记录
// Author [piexlmax](https://github.com/piexlmax)
func (resourcesTestService *ResourcesTestService) GetResourcesTestInfoList(info lessondataReq.ResourcesTestSearch) (list []lessondata.ResourcesTest, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&lessondata.ResourcesTest{})
	var resourcesTests []lessondata.ResourcesTest
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ResourcesId != nil {
		db = db.Where("resources_id = ?", info.ResourcesId)
	}
	if info.QuestionId != nil {
		db = db.Where("question_id = ?", info.QuestionId)
	}
	if info.Weight != nil {
		db = db.Where("weight = ?", info.Weight)
	}
	if info.Orders != nil {
		db = db.Where("ordes = ?", info.Orders)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&resourcesTests).Error
	return resourcesTests, total, err
}
