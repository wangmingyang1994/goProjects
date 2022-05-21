package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goProjects/bookManageSystem/router"
	"goProjects/bookManageSystem/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	s := &http.Server{
		Addr:           ":" + utils.ServerMsg.HttpPort,
		Handler:        routers,
		ReadTimeout:    utils.ServerMsg.ReadTimeout*time.Second,
		WriteTimeout:   utils.ServerMsg.WriteTimeout*time.Second,
		MaxHeaderBytes: 1 >> 20,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal(err.Error())
		}
	}()
	zap.L().Info("server is running...")
	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)  // 此处不会阻塞
	<-quit  // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := s.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown")
		zap.L().Fatal(err.Error())
	}

	zap.L().Info("Server exiting")


}



