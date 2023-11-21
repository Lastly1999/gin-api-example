package service

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-api-example/global"
	"mime/multipart"
	"path/filepath"
)

type OssService struct{}

// UploadFile2AliyunOss 上传文件到阿里云OSS
func (s *OssService) UploadFile2AliyunOss(file multipart.File, fileName string) error {
	bucket, err := s.GetAliyunOssBucket()
	if err != nil {
		return err
	}
	if err := bucket.PutObject(fileName, file); err != nil {
		return err
	}
	return nil
}

// GetAliyunOssBucket 获取阿里云OSS Bucket
func (s *OssService) GetAliyunOssBucket() (*oss.Bucket, error) {
	bucket, err := global.GLOBAL_ALIYUN_OSS_CLIENT.Bucket(global.GLOBAL_CONFIG.Aliyun.Oss.Bucket)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}

func (s *OssService) getFileExtension(filename string) string {
	// 检查文件名是否为空
	if filename == "" {
		return ""
	}

	// 使用 filepath.Ext 获取文件的后缀名
	ext := filepath.Ext(filename)

	// 检查是否没有点号，或者点号在文件名的开头
	if ext == "" || ext == "." {
		return ""
	}

	// 去除文件名中的点号
	return ext
}
