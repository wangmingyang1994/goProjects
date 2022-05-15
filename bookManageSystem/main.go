package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goProjects/bookManageSystem/router"
	"goProjects/bookManageSystem/utils"
	"log"
	"net/http"
)

func init(){
	err:= utils.SetupSetting()
	if err!=nil{
		log.Fatalf("init setupSetting err:%v\n",err)
	}
	err= utils.SetupDB()
	if err!=nil{
		log.Fatalf("init SetupDB err:%v\n",err)
	}
	err= utils.SetupLogger(utils.LogMsg)
	if err!=nil{
		log.Fatalf("init SetupLogger err:%v\n",err)
	}
}

func main() {
	gin.SetMode(utils.ServerMsg.RunMode)
	routers := router.NewRouter()
	zap.L().Info("hello ,debug zap")
	s := &http.Server{
		Addr:           ":" + utils.ServerMsg.HttpPort,
		Handler:        routers,
		ReadTimeout:    utils.ServerMsg.ReadTimeout,
		WriteTimeout:   utils.ServerMsg.WriteTimeout,
		MaxHeaderBytes: 1 >> 20,
	}
	if err:=s.ListenAndServe();err!=nil{
		log.Fatalf("start server error:%v\n", err.Error())
	}


}



