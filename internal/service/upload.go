package service

import (
	"errors"
	"ginblog_backend/global"
	"ginblog_backend/pkg/upload"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessURL string
}

func (s *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	uploadPath := upload.GetSavePath()
	dst := uploadPath + "/" + fileName
	if !upload.CheckContainExts(fileType, dst) {
		return nil, errors.New("file EXT not supported")
	}
	if upload.CheckIsPermissionDenied(dst) {
		return nil, errors.New("failed to create save directory, permission denied")
	}
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded the maximum file limit")
	}
	if !upload.CheckSavePathExisted(dst) {
		return nil, errors.New("file already existed")
	}
	if err := upload.CreateSavePath(uploadPath, os.ModePerm); err != nil {
		return nil, errors.New("failed to create save directory.")
	}
	accessUrl := global.AppSetting.UploadServerURL + "/" + fileName
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}
	return &FileInfo{
		Name:      fileName,
		AccessURL: accessUrl,
	}, nil
}
