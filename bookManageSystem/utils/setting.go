package utils

import (
	"github.com/spf13/viper"
	"time"
)

var (
	ServerMsg *ServerSettings
	DatabaseMsg *DatabaseSettings
	LogMsg *LogSettings
)

type Setting struct{
	vp *viper.Viper
}

func (s *Setting)ReadSection(k string, v interface{})error{
	// 将默认配置写入结构体
	err:=s.vp.UnmarshalKey(k,v)
	if err!=nil{
		return err
	}
	return nil
}

func NewSetting()(*Setting,error){
	// 初始化viper默认的文件
	vp:=viper.New()
	vp.AddConfigPath("config/")
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	// 读取默认配置
	err:=vp.ReadInConfig()
	if err!=nil{
		return nil, err
	}
	return &Setting{vp}, nil
}

func SetupSetting()error{
	// 初始化setting
	setting,err:=NewSetting()
	if err!=nil{
		return  err
	}
	// 读取配置
	err =setting.ReadSection("Server", &ServerMsg)
	if err!=nil{
		return  err
	}
	err=setting.ReadSection("Database" , &DatabaseMsg)
	if err!=nil{
		return  err
	}
	err=setting.ReadSection("Log" , &LogMsg)
	if err!=nil{
		return  err
	}
	return nil
}

type ServerSettings struct{
	RunMode string
	HttpPort string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

type DatabaseSettings struct{
	DBType string
	UserName string
	Password string
	Host string
	DBName string
	TablePrefix string
	Charset string
	ParseTime bool
	MaxIdleConns int
	MaxOpenConns int
}

type LogSettings struct {
	Level string `json:"level"`
	Filename string `json:"filename"`
	MaxSize int `json:"maxsize"`
	MaxAge int `json:"max_age"`
	MaxBackups int `json:"max_backups"`
}

