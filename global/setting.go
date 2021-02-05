package global

import (
	"ginblog_backend/pkg/logger"
	"ginblog_backend/pkg/setting"
)

var (
	ServerSetting   setting.ServerSettingS
	AppSetting      setting.AppSettingS
	DatabaseSetting setting.DatabaseSettingS
	JWTSetting      setting.JWTSettings
	Logger          *logger.Logger
)
