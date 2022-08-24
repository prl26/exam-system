/*
*

	@author: qianyi  2022/8/24 18:56:00
	@note:
*/
package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/basicdata/request"
)

type MultiTableService struct {
}

// 根据传入 教学班id和全部学生id 关联教学班，即更新bas_teach_class_student
func (operation *MultiTableService) UpdateTeachClassStudents(info request.StuTeachClass) error {
	//classID := info.Teach_class_id
	//stuIDs := info.Student_ids
	//
	//for i := 0; i < len(stuIDs); i++ {
	//	global.GVA_DB.Create(&basicdata.TeachClassStudent{
	//		Student_id: ,
	//		Teach_class_id: ,
	//	})
	//}

	return nil
}
