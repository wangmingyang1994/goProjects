package main

import (
	"fmt"
	"goProjects/chatRoom/chat"
	"goProjects/chatRoom/user"
	"google.golang.org/grpc"
	"net"
)

func main() {
	g := grpc.NewServer()
	user.RegisterUserServiceServer(g, &user.User{})
	chat.RegisterChatServer(g, &chat.Chat{})
	lis, err := net.Listen("tcp", "0.0.0.0:8088")
	if err != nil {
		fmt.Println("监听端口失败：" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		fmt.Println("开启grpc失败：" + err.Error())
	}

}
