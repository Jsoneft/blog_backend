package global

import (
	"ginblog_backend/pkg/logger"
	"ginblog_backend/pkg/setting"
)

var (
	ServerSetting   setting.ServerSettingS
	AppSetting      setting.AppSettingS
	DatabaseSetting setting.DatabaseSettingS
	EmailSetting    setting.EmailSettingS
	JWTSetting      setting.JWTSettingS
	Logger          *logger.Logger
)
