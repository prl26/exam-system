package questionBank

import "github.com/prl26/exam-system/server/global"

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/14 15:02

 * @Note:

 **/

type ChapterMerge struct {
	global.GVA_MODEL
	QuestionId   uint `json:"question_id" gorm:"question_id"`
	ChapterId    uint `json:"chapter_id" gorm:"chapter_id"`
	QuestionType uint `json:"question_type" gorm:"question_type"`
}

func (ChapterMerge) TableName() string {
	return "les_questionBank_chapter_merge"
}
