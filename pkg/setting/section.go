package setting

import (
	"time"
)

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize      int
	MaxPageSize          int
	LogSavePath          string
	LogFileName          string
	LogFileExt           string
	UploadSavePath       string
	UploadServerURL      string
	UploadImageMaxSize   int
	UploadImageAllowExts []string
}

type DatabaseSettingS struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	Port         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type JWTSettings struct {
	Secret string
	Issuer string
	Expire time.Duration
}

func (a *Setting) ReadSection(k string, v interface{}) error {
	err := a.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
