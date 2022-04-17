package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"goProjects/chatRoom/chat"
	"goProjects/chatRoom/user"
	"google.golang.org/grpc"
	"strconv"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	userService := user.NewUserServiceClient(conn)
	chatService := chat.NewChatClient(conn)
	//注册服务
	SignUp := &cobra.Command{
		Use:  "signUp",
		Long: "请输入您的姓名和密码！",
		Run: func(cmd *cobra.Command, args []string) {
			SignMessage := &user.UserRegister{
				UserName: args[0],
				Password: args[1],
			}
			fmt.Println(SignMessage)
			r, err := userService.SignUp(context.Background(), SignMessage)
			if err != nil {
				fmt.Printf("注册失败，err:%v\n", err.Error())
				return
			}
			fmt.Printf("注册成功！您的用户id: %d,您的name:%s\n ", r.UserId, r.UserName)
		},
	}
	//登录服务
	Login := &cobra.Command{
		Use:  "login",
		Long: "请输入您的用户Id和密码！",
		Run: func(cmd *cobra.Command, args []string) {
			result, _ := strconv.Atoi(args[0])
			userId := int32(result)
			loginMessage := &user.UserLogin{
				UserId:   userId,
				Password: args[1],
			}
			r1, err := userService.SignIn(context.Background(), loginMessage)
			if err != nil {
				fmt.Printf("登录失败，err:%v\n", err.Error())
				return
			}
			fmt.Printf("登录成功！您的用户id: %d,您的name:%s\n", r1.UserId, r1.UserName)
		},
	}
	//获取用户列表
	GetUsers := &cobra.Command{
		Use:  "getUsers",
		Long: "请输入要获取的页码！",
		Run: func(cmd *cobra.Command, args []string) {
			result, _ := strconv.Atoi(args[0])
			page1 := int32(result)
			needUsers := &user.Needusers{
				Page:     page1,
				Pagenums: 20,
			}
			infos, err := userService.GetUserInfos(context.Background(), needUsers)
			if err != nil {
				return
			}
			if err != nil {
				fmt.Printf("获取用户列表失败，err:%v\n", err.Error())
				return
			}
			if len(infos.Users) == 0 {
				fmt.Println("暂时无可聊天用户，请稍后再试～")
				return
			}
			for _, v := range infos.Users {
				fmt.Printf("UserId:%d,UserName:%s\n", v.UserId, v.UserName)
			}
		},
	}
	//开始聊天
	StartChat := &cobra.Command{
		Use:  "startChat",
		Long: "请输入您的用户Id和好友Id,开启聊天！",
		Run: func(cmd *cobra.Command, args []string) {
			result, _ := strconv.Atoi(args[0])
			result1, _ := strconv.Atoi(args[1])
			userId := int32(result)
			targetId := int32(result1)
			requestRoom := chat.RequestRoom{MyId: userId, TargetId: targetId}
			Room, err := chatService.GetRoom(context.Background(), &requestRoom)
			if err != nil {
				fmt.Printf("查询聊天失败：err:%v\n", err.Error())
				return
			}
			if Room != nil {
				fmt.Println("聊天室Id:", Room.RoomId)
				fmt.Println("=======可以开始聊天啦=======")
				room := chat.Room{RoomId: Room.RoomId}
				msgs, err := chatService.LoadHistory(context.Background(), &room)
				if err != nil {
					fmt.Println("获取历史聊天失败", err.Error())
					return
				}
				var maxId int32
				for _, v := range msgs.ChatMessages {
					fmt.Printf("%s/(%s):%s\n", v.CreateTime, v.UserName, v.ChatMessage)
					if v.MessageId > maxId {
						maxId = v.MessageId
					}
				}
				for {
					var msg string
					time.Sleep(time.Second)
					getNew := chat.GetNew{Lastid: maxId, RoomId: Room.RoomId}
					msgs, err := chatService.GetChat(context.Background(), &getNew)
					if err != nil {
						fmt.Println("获取聊天失败", err.Error())
						return
					}
					for _, v := range msgs.ChatMessages {
						fmt.Printf("%s/(%s):%s\n", v.CreateTime, v.UserName, v.ChatMessage)
					}
					fmt.Scanln(&msg)
					requestSendChatMsg := chat.RequestSendChatMsg{RoomId: Room.RoomId, MyId: userId, ChatMessage: msg}
					_, err = chatService.SendChat(context.Background(), &requestSendChatMsg)
					if err != nil {
						fmt.Println("发送聊天失败", err.Error())
						return
					}
					if msg == "exit" {
						fmt.Println("bye~")
						return
					}

				}

			} else {
				Room1, err := chatService.CreateChat(context.Background(), &requestRoom)
				if err != nil {
					fmt.Printf("创建聊天失败：err:%v\n", err.Error())
					return
				}
				fmt.Println("聊天室Id:", Room1.RoomId)
				fmt.Println("=======可以开始聊天啦=======")
				for {
					var msg string
					time.Sleep(time.Second)
					getNew := chat.GetNew{Lastid: 0, RoomId: Room.RoomId}
					msgs, err := chatService.GetChat(context.Background(), &getNew)
					if err != nil {
						fmt.Println("获取聊天失败", err.Error())
						return
					}
					for _, v := range msgs.ChatMessages {
						fmt.Printf("%s/(%s):%s\n", v.CreateTime, v.UserName, v.ChatMessage)
					}
					fmt.Scanln(&msg)
					requestSendChatMsg := chat.RequestSendChatMsg{RoomId: Room.RoomId, MyId: userId, ChatMessage: msg}
					_, err = chatService.SendChat(context.Background(), &requestSendChatMsg)
					if err != nil {
						fmt.Println("发送聊天失败", err.Error())
						return
					}
					if msg == "exit" {
						fmt.Println("bye~")
						return
					}

				}

			}

		},
	}
	// 命令行主入口
	var rootCmd = &cobra.Command{Use: "app"}
	// 注册子命令
	rootCmd.AddCommand(SignUp, Login, GetUsers, StartChat)
	rootCmd.Execute()

}
