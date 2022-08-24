/*
*

	@author: qianyi  2022/8/24 19:18:00
	@note:
*/
package request

// 接收 教学班id 和学生id 的结构体
type StuTeachClass struct {
	TeachClassId int   `json:"teach_class_id"`
	StudentIds   []int `json:"student_ids"`
}
