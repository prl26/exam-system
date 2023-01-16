package basicdata

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	commonError "github.com/prl26/exam-system/server/model/common/error"
	"github.com/prl26/exam-system/server/model/lessondata"
	"github.com/prl26/exam-system/server/model/questionBank/enum/questionType"
	questionBank "github.com/prl26/exam-system/server/model/questionBank/po"
	"github.com/prl26/exam-system/server/model/teachplan"
)

func (lessonService *LessonService) FindLessonDetail(id uint, all bool, studentId uint) (*basicdata.Lesson, error) {
	var lesson basicdata.Lesson
	if err := global.GVA_DB.Model(&lesson).Where("id=?", id).Preload("Chapters").Find(&lesson).Error; err != nil {
		return nil, err
	}
	if lesson.ID == 0 {
		return nil, commonError.NotFoundError
	}
	if all {
		for _, chapter := range lesson.Chapters {
			if err := global.GVA_DB.Model(&chapter).Preload("Knowledges").Find(&chapter).Error; err != nil {
				return nil, err
			}
			chapterQuestionNumber := int64(0)
			chapterDoneNumber := int64(0)
			for _, knowledge := range chapter.Knowledges {
				ids := []uint{}
				if err := global.GVA_DB.Model(&questionBank.Target{}).Where("is_check=? and can_practice=? and knowledge_id=?", 1, 1, knowledge.ID).Select("id").Find(&ids).Error; err != nil {
					return nil, err
				}
				questionNumber := int64(len(ids))
				knowledge.QuestionCount = &questionNumber
				doneNumber := int64(0)
				if questionNumber != 0 {
					if err := global.GVA_DB.Model(&teachplan.PracticeAnswer{}).Where("student_id=? and question_type=? and question_id in ?", studentId, questionType.Target, ids).Count(&doneNumber).Error; err != nil {
						return nil, err
					}
					knowledge.DoneCount = &doneNumber
				}
				knowledge.DoneCount = &doneNumber
				chapterQuestionNumber += questionNumber
				chapterDoneNumber += doneNumber
			}
			chapter.QuestionCount = &chapterQuestionNumber
			chapter.DoneCount = &chapterDoneNumber
		}
	}
	return &lesson, nil
}

func (lessonService *LessonService) FindKnowledge(idUint uint) ([]*lessondata.Knowledge, error) {
	var chapters []*lessondata.Knowledge
	if err := global.GVA_DB.Where("chapter_id=?", idUint).Find(&chapters).Error; err != nil {
		return nil, err
	}
	return chapters, nil
}
