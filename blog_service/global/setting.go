package global

import (
	"learn.go/blog_service/pkg/logger"
	"learn.go/blog_service/pkg/setting"
)

var (
	ServerSetting *setting.ServerSettings
	AppSetting *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	Logger *logger.Logger
)
