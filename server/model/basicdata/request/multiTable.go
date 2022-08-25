/*
*

	@author: qianyi  2022/8/24 19:18:00
	@note:
*/
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// 接收 教学班id 和学生id 的结构体
type StuTeachClass struct {
	TeachClassId uint   `json:"teachClassId"`
	StudentIds   []uint `json:"studentIds"`
}

type TeachClassStudent struct {
	TeachClassId uint `json:"teachClassId"`
	request.PageInfo
}
