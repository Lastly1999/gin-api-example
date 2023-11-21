package service

import (
	"fmt"
	"github.com/gin-api-example/api/response"
	"github.com/gin-api-example/global"
	"mime/multipart"
	"strconv"
	"time"
)

type UploadService struct{}

// UploadFile 上传文件
func (s *UploadService) UploadFile(file *multipart.FileHeader) (*response.UploadFileRes, error) {
	fileHandle, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fileHandle.Close()
	ossService := OssService{}
	fileExt := ossService.getFileExtension(file.Filename)
	fileName := s.GetTimestarpFileName() + fileExt
	err = ossService.UploadFile2AliyunOss(fileHandle, fileName)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://%s.%s/%s", global.GLOBAL_CONFIG.Aliyun.Oss.Bucket, global.GLOBAL_CONFIG.Aliyun.Oss.Endpoint, fileName)
	return &response.UploadFileRes{
		Url:  url,
		Size: file.Size,
		Name: fileName,
	}, nil
}

// GetTimestarpFileName 获取时间戳文件名
func (s *UploadService) GetTimestarpFileName() string {
	t := time.Now()
	fileName := fmt.Sprintf("%s%s%s%s%s%s%s", strconv.Itoa(t.Year()), strconv.Itoa(int(t.Month())), strconv.Itoa(t.Day()), strconv.Itoa(t.Hour()), strconv.Itoa(t.Minute()), strconv.Itoa(t.Second()), strconv.Itoa(int(t.Unix())))
	return fileName
}
