package main

import (
	"github.com/gin-gonic/gin"
	"goProjects/bookManageSystem/global"
	"goProjects/bookManageSystem/router"
	"goProjects/bookManageSystem/utils"
	"log"
	"net/http"
	"time"
)

func init(){
	err:= setupSetting()
	if err!=nil{
		log.Fatalf("init setupSetting err:%v\n",err)
	}
}

func main() {
	gin.SetMode(global.ServerSettings.RunMode)
	router := router.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSettings.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSettings.ReadTimeout,
		WriteTimeout:   global.ServerSettings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}


func setupSetting()error{
	setting,err:=utils.NewSetting()
	if err!=nil{
		return err
	}
	err =setting.ReadSection("Server", &global.ServerSettings)
	if err!=nil{
		return err
	}
	err=setting.ReadSection("Database" , &global.DatabaseSettings)
	if err!=nil{
		return err
	}
	global.ServerSettings.ReadTimeout *=time.Second
	global.ServerSettings.WriteTimeout *= time.Second
	return nil
}