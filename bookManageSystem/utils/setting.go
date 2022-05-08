package utils

import (
	"github.com/spf13/viper"
	"time"
)

type Setting struct{
	vp *viper.Viper
}
func NewSetting()(*Setting,error){
	vp:=viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("config/")
	vp.SetConfigType("yaml")
	err:=vp.ReadInConfig()
	if err!=nil{
		return nil, err
	}
	return &Setting{vp}, nil
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


func (s *Setting)ReadSection(k string, v interface{})error{
	err:=s.vp.UnmarshalKey(k,v)
	if err!=nil{
		return err
	}
	return nil
}