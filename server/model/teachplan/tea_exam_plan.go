// 自动生成模板ExamPlan
package teachplan

import (
	"github.com/prl26/exam-system/server/global"
	"time"
)

// ExamPlan 结构体
type ExamPlan struct {
	global.GVA_MODEL
	Name          string     `json:"name" form:"name" gorm:"column:name;comment:考试名称;size:32;"`
	TeachClassId  *uint      `json:"teachClassId" form:"teachClassId" gorm:"column:teach_class_id;comment:教学班id;size:32;"`
	Time          *int64     `json:"time" form:"time" gorm:"column:time;comment:考试时长;"`
	StartTime     *time.Time `json:"startTime" form:"startTime" gorm:"column:start_time;comment:考试时间;"`
	EndTime       *time.Time `json:"endTime" form:"endTime" gorm:"column:end_time;comment:结束时间;"`
	LessonId      *int       `json:"lessonId" form:"lessonId" gorm:"column:lesson_id;comment:课程Id;size:32;"`
	TemplateId    *uint      `json:"templateId" form:"templateId" gorm:"column:template_id;comment:考试模板Id;size:32;"`
	State         *int       `json:"state" form:"state" gorm:"column:state;comment:状态;size:8;"`
	Audit         *int       `json:"audit" form:"audit" gorm:"column:audit;comment:是否审核;size:8;"`
	Type          int        `json:"type" form:"type" gorm:"column:type;comment:考试类型;size:8;"`
	PassScore     *float64   `json:"passScore" form:"passScore" gorm:"column:pass_score;comment:通过分数;size:8;"`
	Weight        *int       `json:"weight" form:"weight" gorm:"column:weight;comment:权重;size:8;"`
	TermId        *uint      `json:"termId" form:"termId" gorm:"column:term_id;comment:学期id"`
	IsDistributed bool       `json:"isDistributed" gorm:"is_distributed"`
	UserId        *uint      `json:"userId" form:"userId" gorm:"column:user_id;comment:创建id;size:32;"`
	PrePlanId     string     `json:"prePlanId" form:"prePlanId" gorm:"pre_plan_id"`
	IsLimitTime   bool       `json:"isLimitTime" form:"isLimitTime" gorm:"is_limit_time;comment:是否限制考试时长"`
	IsReady       bool       `json:"isReady" form:"isReady" gorm:"is_ready"`
}
type CoverRq struct {
	StudentId uint `json:"studentId"`
	PlanId    uint `json:"planId"`
}

//type ExamPlan

// TableName ExamPlan 表名
func (ExamPlan) TableName() string {
	return "tea_examplan"
}
