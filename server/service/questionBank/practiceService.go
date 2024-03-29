package questionBank

import (
	"context"
	"fmt"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/questionBank/enum/questionType"
	questionBankVoResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
	"github.com/prl26/exam-system/server/model/teachplan"
	teachplanResp "github.com/prl26/exam-system/server/model/teachplan/response"
	"time"
)

type PracticeService struct {
}

func (p PracticeService) FindTheLatestRecord(lessonId, studentId uint) *teachplan.PracticeRecord {
	t := teachplan.PracticeRecord{}
	global.GVA_DB.Where("lesson_id=? and student_id=?", lessonId, studentId).Order("begin_time desc").First(&t)
	return &t
}

func (p PracticeService) CreatePracticeRecord(r *teachplan.PracticeRecord) error {
	return global.GVA_DB.Create(r).Error
}

func (p PracticeService) FindTheLatestItemId(lessonId, studentId uint) *uint {
	result := new(uint)
	global.GVA_DB.Model(&teachplan.PracticeItem{}).Select("id").Where("lessonId=? and studentId=?", lessonId, studentId).Order("answer_time desc").Find(result)
	return result
}

func (p PracticeService) RemoveRecord(id uint) {
	global.GVA_DB.Where("id=?", id).Delete(&teachplan.PracticeRecord{})
}

func (p PracticeService) FindItemCount(lessonId, studentId uint, beginIndex, endIndex *uint) (result int64) {
	global.GVA_DB.Where("id>=? && id<=? && lessonId=? and studentId=?", beginIndex, endIndex, lessonId, studentId).Count(&result)
	return
}

func (p PracticeService) UpdatePracticeRecord(lessonId, studentId uint) {
	record := p.FindTheLatestRecord(lessonId, studentId)
	if record.ID != 0 {
		a := 0
		global.GVA_DB.Raw("update tea_practice_record a \njoin (select count(*) as count\nfrom tea_practice_item\nwhere tea_practice_item.record_id=?) b\nset a.question_count=b.count,a.deleted_at=if(b.count=0,NOW(),null) \nwhere a.id=?", record.ID, record.ID).Scan(&a)
	}
	return
}

func (p PracticeService) CreatePracticeItem(questionType questionType.QuestionType, questionId, lessonId, studentId uint, score uint, answer string) {
	record := p.FindTheLatestRecord(lessonId, studentId)
	if record.ID != 0 {
		a := 0
		global.GVA_DB.Raw("INSERT INTO tea_practice_item(student_id,question_type,lesson_id,question_id,record_id,score,answer,commit_time) VALUES (?,?,?,?,?,?,?,now())\n  ON DUPLICATE KEY UPDATE commit_time=now(),score=?",
			studentId, uint(questionType), lessonId, questionId, record.ID, score, answer, score).Scan(&a)
	}
}

func (p PracticeService) UpdatePracticeAnswer(questionType questionType.QuestionType, questionId, lessonId, studentId uint, score uint, answer string) {
	global.GVA_DB.Raw("INSERT INTO tea_practice_answer(student_id,question_type, question_id, lesson_id,score,answer)\nVALUES ( ?, ?, ?,?,?,?)\nON DUPLICATE KEY UPDATE score = GREATEST(score,?) , answer=?", studentId, questionType, questionId, lessonId, score, answer, score, answer).Scan(nil)
	return
}

func (p PracticeService) FindHistoryAnswer(questionType questionType.QuestionType, ids []uint, studentId uint) *teachplanResp.History {
	histories := []teachplan.PracticeAnswer{}
	global.GVA_DB.Where("student_id=? and question_type=? and question_id in ?", studentId, questionType, ids).Find(&histories)
	history := teachplanResp.History{History: map[uint]*teachplanResp.HistoryItem{}}
	for _, t := range histories {
		item := teachplanResp.HistoryItem{
			Exist:  true,
			Score:  t.Score,
			Answer: t.Answer,
		}
		history.History[t.QuestionId] = &item
	}
	for _, id := range ids {
		if history.History[id] == nil {
			item := teachplanResp.HistoryItem{
				Exist: false,
				Score: 0,
			}
			history.History[id] = &item
		}
	}
	return &history
}

func (p PracticeService) CanNewPracticeRecord(lessonId uint, studentId uint) bool {
	str := fmt.Sprintf("newPracticeRecord:%d:%d", studentId, lessonId)
	_, err := global.GVA_REDIS.Get(context.Background(), str).Result()
	if err == nil {
		return false
	} else {
		global.GVA_REDIS.Set(context.Background(), str, true, 20*time.Minute)
		return true
	}
}

func (p PracticeService) RankingList(lessonId uint, info request.PageInfo) (list []questionBankVoResp.RankingListItem, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	err = global.GVA_DB.Model(&teachplan.PracticeAnswer{}).Group("student_id").Where("lesson_id=?", lessonId).Count(&total).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Raw("select a.student_id,sum(a.score) as total_score,count(a.score) as problem_count,b.name from tea_practice_answer a left join bas_student  b on a.student_id=b.id where lesson_id=? GROUP BY student_id ORDER BY total_score desc,student_id limit ?  OFFSET ?", lessonId, limit, offset).Find(&list).Error
	if err != nil {
		return
	}
	if len(list) != 0 {
		var rank int64
		err = global.GVA_DB.Raw("select count(total_score)+1 as count from (select sum(score) total_score    from tea_practice_answer  where lesson_id=? GROUP BY student_id having total_score>?) a", lessonId, list[0].TotalScore).First(&rank).Error
		list[0].Rank = uint(rank)
		next := offset + 1
		for i := 1; i < len(list); i++ {
			next++
			if list[i].TotalScore == list[i-1].TotalScore {
				list[i].Rank = list[i-1].Rank
			} else {
				list[i].Rank = uint(next)
			}
		}
	}
	return
}

func (p PracticeService) GetMyRank(lessonId int, studentId uint) (item questionBankVoResp.RankingListItem, err error) {
	err = global.GVA_DB.Raw("select a.student_id,sum(a.score) as total_score,count(a.score) as problem_count,b.name from tea_practice_answer a left join bas_student  b on a.student_id=b.id where lesson_id=? and student_id=? GROUP BY student_id ", lessonId, studentId).Find(&item).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Raw("select count(total_score)+1 as count from (select DISTINCT sum(score) total_score    from tea_practice_answer  where lesson_id=? GROUP BY student_id having total_score>?) a", lessonId, item.TotalScore).First(&item.Rank).Error
	return
}

func (p PracticeService) FindAnswer(questionT questionType.QuestionType, id uint) (answer questionBankVoResp.Answer) {
	if questionT == questionType.PROGRAM {
		global.GVA_DB.Raw("select reference_answers from les_questionbank_programm where id=?", id).Scan(&answer.Answer)
	}
	if answer.Answer == "" || answer.Answer == "{}" {
		answer.Answer = "暂无参考答案"
	}
	return
}
