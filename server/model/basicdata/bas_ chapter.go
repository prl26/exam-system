// 自动生成模板Chapter
package basicdata

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/lessondata"
)

// Chapter 结构体
type Chapter struct {
	global.GVA_MODEL
	Name          string                  `json:"name" form:"name" gorm:"column:name;comment:章节名称;size:32;"`
	LessonId      *int                    `json:"lessonId" form:"lessonId" gorm:"column:lesson_id;comment:所属课程id;size:32;"`
	Knowledges    []*lessondata.Knowledge `json:"knowledges" gorm:"foreignKey:ChapterId"`
	QuestionCount *int64                  `json:"questionCount,omitempty"`
	DoneCount     *int64                  `json:"doneCount,omitempty"`
}

// TableName Chapter 表名
func (Chapter) TableName() string {
	return "bas_chapter"
}
