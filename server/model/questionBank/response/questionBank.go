package response

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/questionBank"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/28 15:44

 * @Note:

 **/

type CourseSupport struct {
	global.GVA_MODEL
	ChapterId   uint   `json:"chapterId" gorm:"chapter_id"`
	ChapterName string `json:"chapterName" gorm:"chapter_name"`
	LessonId    uint   `json:"lessonId" gorm:"lesson_id"`
	LessonName  string `json:"lessonName" gorm:"lesson_name"`
}

type QuestionSupport struct {
	global.GVA_MODEL
	questionBank.BasicModel
}
