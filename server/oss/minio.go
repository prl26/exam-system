package oss

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"mime/multipart"
	"path"
	"strings"
	"time"
)

type Minio struct {
	*minio.Client
	defaultBucketName string // 默认桶的名字
}

func NewMinio(endPoint string, accessKeyID string, secretAccessKey string, bucketName string) (*Minio, error) {
	minioClient, err := minio.New(endPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}
	m := &Minio{
		defaultBucketName: bucketName,
		Client:            minioClient,
	}

	return m, nil
}

func (m *Minio) UploadMultipartFile(multipartFile *multipart.FileHeader) (string, string, error) {
	return m.UploadMultipartFileWithPrefix(multipartFile, "default")
}

func (m *Minio) UploadMultipartFileWithPrefix(multipartFile *multipart.FileHeader, prefix string) (string, string, error) {
	fileName := multipartFile.Filename
	file, err := multipartFile.Open()
	if err != nil {
		return "", "", err
	}
	return m.UploadFileWithPrefix(fileName, file, multipartFile.Size, prefix)
}

func (m *Minio) UploadFileWithPrefix(fileName string, file io.Reader, fileSize int64, prefix string) (string, string, error) {
	ext := path.Ext(fileName)
	// 读取文件名并加密
	name := strings.TrimSuffix(fileName, ext)
	name = MD5V([]byte(name))
	// 拼接新文件名
	filename := prefix + "/" + name + "_" + time.Now().Format("20060102150405") + ext
	f, err := m.PutObject(context.Background(), m.defaultBucketName, filename, file, fileSize, minio.PutObjectOptions{})
	return filename, f.Bucket + "/" + f.Key, err
}

func (*Minio) DeleteFile(key string) error {
	// TODO  删除文件
	return nil
}

func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}
