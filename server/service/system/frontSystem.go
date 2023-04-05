package system

import (
	"errors"
	"fmt"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	basicdataReq "github.com/prl26/exam-system/server/model/basicdata/request"
	"github.com/prl26/exam-system/server/model/system"
	"github.com/prl26/exam-system/server/model/teachplan"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

type FrontSystemService struct {
}

func (frontSystemService *FrontSystemService) GetTermInfoList(info basicdataReq.FrontTermSearch) (list []basicdata.Term, total int64, err error) {

	// 创建db
	db := global.GVA_DB.Model(&basicdata.Term{})
	var terms []basicdata.Term
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Find(&terms).Error
	return terms, total, err
}
func (frontSystemService *FrontSystemService) GetLessonInfoList(info basicdataReq.FrontLessonSearch) (list []basicdata.Lesson, total int64, err error) {

	// 创建db
	db := global.GVA_DB.Model(&basicdata.Lesson{})
	var lessons []basicdata.Lesson
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Find(&lessons).Error
	return lessons, total, err
}
func (frontSystemService *FrontSystemService) UploadFile(header *multipart.FileHeader, noSave string, planId uint, studentId uint) (file system.ExaFileUploadAndDownload, err error) {
	var planDetail teachplan.ExamPlan
	global.GVA_DB.Model(teachplan.ExamPlan{}).Where("id = ?", planId).Find(&planDetail)
	filePath, key, uploadErr := frontSystemService.UploadFile1(header, planDetail.Name, studentId)
	if uploadErr != nil {
		panic(err)
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := system.ExaFileUploadAndDownload{
			Url:  filePath,
			Name: header.Filename,
			Tag:  s[len(s)-1],
			Key:  key,
		}
		return f, nil
	}
	return
}
func (frontSystemService *FrontSystemService) UploadFile1(file *multipart.FileHeader, planName string, studentId uint) (string, string, error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(fmt.Sprintf("%s/%s/%d", global.GVA_CONFIG.Local.StorePath, planName, studentId), os.ModePerm)
	if mkdirErr != nil {
		global.GVA_LOG.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := fmt.Sprintf("%s/%s/%d", global.GVA_CONFIG.Local.StorePath, planName, studentId) + "/" + filename
	filepath := fmt.Sprintf("%s/%s/%d", global.GVA_CONFIG.Local.StorePath, planName, studentId) + "/" + filename

	f, openError := file.Open() // 读取文件
	if openError != nil {
		global.GVA_LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p)
	if createErr != nil {
		global.GVA_LOG.Error("function os.Create() Filed", zap.Any("err", createErr.Error()))

		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		global.GVA_LOG.Error("function io.Copy() Filed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return filepath, filename, nil
}
