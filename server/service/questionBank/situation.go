package questionBank

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/questionBank/enum/questionType"
	questionBank "github.com/prl26/exam-system/server/model/questionBank/po"
	questionBankVoResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
	"github.com/prl26/exam-system/server/model/teachplan"
)

type SituationService struct {
}

func (s SituationService) FindTeachClassSituation(info request.PageInfo, lessonId uint, teachClassId uint) (data []questionBankVoResp.QuestionSituation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	err = global.GVA_DB.Model(&basicdata.StudentAndTeachClass{}).Where("teach_class_id=?", teachClassId).Count(&total).Error
	if err != nil || total == 0 {
		return
	}
	// 创建db
	sql := `select  b.student_id,b.record_count,count(tea_practice_answer.id) as question_count from
			(select a.student_id,count(tea_practice_record.student_id) as record_count from tea_practice_record right join (
				select bas_student.id as student_id,bas_student.name
				from bas_student_teach_classes
				left join bas_student
				on bas_student.id=bas_student_teach_classes.student_id
				where bas_student_teach_classes.teach_class_id=?
				limit ?,?)  a 
			on a.student_id=tea_practice_record.student_id and tea_practice_record.lesson_id=? where isnull(tea_practice_record.deleted_at) GROUP BY a.student_id) b 
		  LEFT JOIN tea_practice_answer on b.student_id=tea_practice_answer.student_id and tea_practice_answer.lesson_id=? 
          LEFT JOIN bas_student on bas_student.id=b.student_id GROUP BY b.student_id`
	err = global.GVA_DB.Raw(sql, teachClassId, offset, limit, lessonId, lessonId).Scan(&data).Error

	return
}

func (s SituationService) FindStudentSituation(info request.PageInfo, lessonId uint, studentId uint) (data []teachplan.PracticeRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&teachplan.PracticeRecord{})
	db = db.Where("lesson_id=?", lessonId)
	db = db.Where("student_id=?", studentId)
	err = db.Count(&total).Error
	if err != nil || total == 0 {
		return
	}
	db = db.Order("begin_time desc")
	err = db.Limit(limit).Offset(offset).Find(&data).Error
	return
}
func (p SituationService) FindTheLatestRecord(lessonId, studentId uint) *teachplan.PracticeRecord {
	t := teachplan.PracticeRecord{}
	global.GVA_DB.Where("lesson_id=? and student_id=?", lessonId, studentId).Order("begin_time desc").First(&t)
	return &t
}

func (s SituationService) UpdateSituation(lessonId uint, studentId uint) {
	record := s.FindTheLatestRecord(lessonId, studentId)
	if record.ID != 0 {
		a := 0
		global.GVA_DB.Raw("update tea_practice_record a \njoin (select count(*) as count\nfrom tea_practice_item\nwhere tea_practice_item.record_id=?) b\nset a.question_count=b.count \nwhere a.id=?", record.ID, record.ID).Scan(&a)
	}
	return

}

func (s SituationService) FindDetail(info request.PageInfo, detailId uint) (data []questionBankVoResp.SituationDetail, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&teachplan.PracticeItem{})
	db = db.Where("record_id=?", detailId)
	err = db.Count(&total).Error
	if err != nil || total == 0 {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&data).Error
	if err != nil {
		return
	}
	table := make(map[questionType.QuestionType][]uint)
	table2 := make(map[questionType.QuestionType]map[uint]*questionBank.BasicModelWith)
	for i := 0; i < len(data); i++ {
		table[data[i].QuestionType] = append(table[data[i].QuestionType], data[i].QuestionId)
	}
	for key, details := range table {
		db := global.GVA_DB
		if key == questionType.SINGLE_CHOICE || key == questionType.MULTIPLE_CHOICE {
			db = db.Model(&questionBank.MultipleChoice{})
		} else if key == questionType.PROGRAM {
			db = db.Model(&questionBank.Program{})
		} else if key == questionType.JUDGE {
			db = db.Model(&questionBank.Judge{})
		} else if key == questionType.Target {
			db = db.Model(&questionBank.Target{})
		}
		table2[key] = make(map[uint]*questionBank.BasicModelWith)
		models := make([]*questionBank.BasicModelWith, len(details))
		if db.Where("id in ?", details).Find(&models).Error != nil {

		}
		for _, model := range models {
			table2[key][model.ID] = model
		}
	}
	for i := 0; i < len(data); i++ {
		data[i].BasicModelWith = table2[data[i].QuestionType][data[i].QuestionId]
	}
	return
}
