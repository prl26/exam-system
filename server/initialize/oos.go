package initialize

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/oss"
)

func NewOss() oss.OSS {
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
	case "minio":
		minio := global.GVA_CONFIG.Minio
		newMinio, err := oss.NewMinio(minio.EndPoint, minio.AccessKeyId, minio.SecretAccessKey, minio.Bucket)
		if err != nil {
			global.GVA_LOG.Panic("newMinio 初始化失败")
		}
		return newMinio
	default:
		global.GVA_LOG.Panic("oss类型输入错误")
		return nil
	}
}
