/*
*

	@author: qianyi  2022/8/24 19:18:00
	@note:
*/
package request

// 接收 教学班id 和学生id 的结构体
type StuTeachClass struct {
	Teach_class_id int64   `json:"teach_class_id"`
	Student_ids    []int64 `json:"student_ids"`
}
