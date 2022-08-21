package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	listenAddr string
)

func main() {
	// 开启服务
	flag.StringVar(&listenAddr, "listen-addr", ":8080", "server listen address")
	flag.Parse()
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	server := newWebserver(logger)
	go gracefulShutdown(server, logger, quit, done)
	logger.Println("Server is ready to handle requests at", listenAddr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}
	// 阻塞，接收关闭channel的信号
	<-done
	logger.Println("Server stopped")
}

func gracefulShutdown(server *http.Server, logger *log.Logger, quit <-chan os.Signal, done chan<- bool) {
	// 接收ctrl+c的信号
	<-quit
	logger.Println("Server is shutting down...")
	//创建context，设置10秒超时
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	server.SetKeepAlivesEnabled(false)
	// 刷入缓存
	time.Sleep(time.Second * 2)
	log.Println("Brush the cache success~")
	time.Sleep(time.Second * 2)
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
	//关闭channel
	close(done)
}

func newWebserver(logger *log.Logger) *http.Server {
	// 注册一个server
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	})

	return &http.Server{
		Addr:         listenAddr,
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}
