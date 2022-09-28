package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"

	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	"go.uber.org/zap"
	"time"
)

var ctx = context.Background()

func SetStudentToRedis(students []uint) error {
	for i := 0; i < len(students); i++ {
		err := global.GVA_REDIS.Set(ctx, strconv.Itoa(int(students[i])), "考生即将考试或正在考试,题库功能关闭", 2*time.Hour).Err()
		if err != nil {
			global.GVA_LOG.Error("RedisStoreSetError!", zap.Error(err))
		}
	}
	return nil
}
func GaSStudentsOfTeachClass(ids []uint) (students []uint, err error) {
	var studentAndTeachClass []basicdata.StudentAndTeachClass
	for i := 0; i < len(ids); i++ {
		err = global.GVA_DB.Where("teach_class_id = ?", i).Find(&studentAndTeachClass).Error
		if err != nil {
			global.GVA_LOG.Error("RedisStoreSetError!", zap.Error(err))
			return
		}
	}
	for _, v := range studentAndTeachClass {
		students = append(students, v.StudentId)
	}
	err = SetStudentToRedis(students)
	if err != nil {
		return
	}
	return
}
func IsExistInRedis(studentId uint) bool {
	_, err := global.GVA_REDIS.Get(ctx, strconv.Itoa(int(studentId))).Result()
	if err == redis.Nil {
		return true
	} else if err != nil {
		panic(err)
	} else {
		return false
	}
}
