/*
*

	@author: qianyi  2022/8/24 18:56:00
	@note:
*/
package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/basicdata"
	"github.com/flipped-aurora/gin-vue-admin/server/model/basicdata/request"
)

type MultiTableService struct {
}

// 根据传入 教学班id和全部学生id 关联教学班，即更新bas_teach_class_student
func (multiTableService *MultiTableService) UpdateTeachClassStudents(info request.StuTeachClass) error {

	// 将数据整合到 表的结构体中方便
	n := len(info.StudentIds)

	list := make([]*basicdata.TeachClassStudent, n)
	for i := 0; i < n; i++ {
		list[i] = &basicdata.TeachClassStudent{
			StudentId:    &info.StudentIds[i],
			TeachClassId: &info.TeachClassId,
		}
	}

	err := global.GVA_DB.Create(&list).Error

	return err
}
