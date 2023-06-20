package upload

import (
	"io"
	"mime/multipart"

	"github.com/prl26/exam-system/server/global"
)

// OSS 对象存储接口
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ccfish86](https://github.com/ccfish86)
type OSS interface {
	UploadMultipartFile(file *multipart.FileHeader) (string, string, error)
	UploadMultipartFileWithPrefix(multipartFile *multipart.FileHeader, prefix string) (string, string, error)
	UploadFileWithPrefix(fileName string, file io.Reader, fileSize int64, prefix string) (string, string, error)
	DeleteFile(key string) error
}

// NewOss OSS的实例化方法
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ccfish86](https://github.com/ccfish86)
func NewOss() OSS {
	switch global.GVA_CONFIG.System.OssType {
	//case "local":
	//	return &Local{}
	//case "qiniu":
	//	return &Qiniu{}
	//case "tencent-cos":
	//	return &TencentCOS{}
	//case "aliyun-oss":
	//	return &AliyunOSS{}
	//case "huawei-obs":
	//	return HuaWeiObs
	//case "aws-s3":
	//	return &AwsS3{}
	//case "minio":
	//	minio := global.GVA_CONFIG.Minio
	//	newMinio, err := NewMinio(minio.EndPoint, minio.AccessKeyId, minio.SecretAccessKey, minio.Bucket)
	//	if err != nil {
	//		global.GVA_LOG.Error("连接minio 失败:" + err.Error())
	//	}
	//	return newMinio
	default:
		return nil
	}
}
