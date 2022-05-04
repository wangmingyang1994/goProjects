package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"goProjects/blog_service/global"
	"goProjects/blog_service/internal/model"
	"goProjects/blog_service/internal/routers"
	"goProjects/blog_service/pkg/logger"
	setting "goProjects/blog_service/pkg/setting"
	"log"
	"net/http"
	"time"
)


func init(){
	err:= setupSetting()
	if err!=nil{
		log.Fatalf("init setupSetting err:%v\n",err)
	}
	err= setupDBEngine()
	if err!=nil{
		log.Fatalf("init setupDBEngine err:%v\n",err)
	}
	err= setupLogger()
	if err!=nil{
		log.Fatalf("init setupLogger err:%v\n",err)
	}
}


func main() {
	//global.Logger.Infof("%s","blog_service")
	gin.SetMode(global.ServerSetting.RunMode)
	router:=routers.NewRouter()
	s:= &http.Server{
		Addr:":"+global.ServerSetting.HttpPort,
		Handler: router,
		ReadTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout: global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1<<20,

	}
	s.ListenAndServe()

}



func setupSetting()error{
	setting,err:=setting.NewSetting()
	if err!=nil{
		return err
	}
	err =setting.ReadSection("Server", &global.ServerSetting)
	if err!=nil{
		return err
	}
	err = setting.ReadSection("App",&global.AppSetting)
	if err!=nil{
		return err
	}
	err=setting.ReadSection("Database" , &global.DatabaseSetting)
	if err!=nil{
		return err
	}
	global.ServerSetting.ReadTimeout *=time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}


func setupDBEngine()error{
	var err error
	global.DBEngine,err=model.NewDBEngine(global.DatabaseSetting)
	if err!=nil{
		return err
	}
	return nil
}

func setupLogger()error{
	fileName :=global.AppSetting.LogSavePath+"/"+global.AppSetting.LogFileName+global.AppSetting.LogFileExt
	global.Logger =logger.NewLogger(&lumberjack.Logger{
		Filename: fileName,
		MaxSize: 600,
		MaxAge: 10,
		LocalTime: true,
	},"",log.LstdFlags).WithCaller(2)
	return nil
}