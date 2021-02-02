package service

import (
	"errors"
	"ginblog_backend/global"
	"ginblog_backend/pkg/upload"
	"mime/multipart"
)

type FileInfo struct {
	Name      string
	AccessURL string
}

func (s *Service) UploadFile(fileType upload.FileType, file multipart.File, header *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(header.Filename)
	uploadPath := upload.GetSavePath()
	dst := uploadPath + "/" + fileName
	if !upload.CheckContainExts(fileType, dst) {
		return nil, errors.New("file EXT not supported")
	}
	if !upload.CheckIsPermission(dst) {
		return nil, errors.New("failed to create save directory, permission denied")
	}
	if !upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded the maximum file limit")
	}
	if !upload.CheckSavePathExisted(dst) {
		return nil, errors.New("file already existed")
	}
	accessUrl := global.AppSetting.UploadServerURL + "/" + fileName
	return &FileInfo{
		Name:      fileName,
		AccessURL: accessUrl,
	}, nil
}
