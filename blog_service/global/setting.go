package global

import (
	"goProjects/blog_service/pkg/logger"
	"goProjects/blog_service/pkg/setting"
)

var (
	ServerSetting *setting.ServerSettings
	AppSetting *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	Logger *logger.Logger
)