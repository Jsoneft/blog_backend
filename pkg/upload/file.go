package upload

import (
	"ginblog_backend/global"
	"ginblog_backend/pkg/utils"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const (
	TypeImage = iota + 1
)

func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.EncodeMD5(fileName)
	return fileName + ext
}

func CheckSavePathExisted(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

func CheckIsPermissionDenied(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		return size >= global.AppSetting.UploadImageMaxSize*1024*1024
	}
	return false
}

func CheckContainExts(t FileType, name string) bool {
	ext := GetFileExt(name)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
				return true
			}
		}
	}
	return false
}

func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.Mkdir(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	return err
}
