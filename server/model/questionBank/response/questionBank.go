package response

import "github.com/prl26/exam-system/server/global"

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/28 15:44

 * @Note:

 **/

type CourseSupport struct {
	global.GVA_MODEL
	ChapterId   uint
	ChapterName string
	LessonId    uint
	LessonName  string
}
