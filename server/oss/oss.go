package oss

import (
	"io"
	"mime/multipart"
)

type OSS interface {
	UploadMultipartFile(file *multipart.FileHeader) (string, string, error)
	UploadMultipartFileWithPrefix(multipartFile *multipart.FileHeader, prefix string) (string, string, error)
	UploadFileWithPrefix(fileName string, file io.Reader, fileSize int64, prefix string) (string, string, error)
	DeleteFile(key string) error
}
