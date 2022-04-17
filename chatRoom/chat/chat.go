package chat

import (
	"context"
	"fmt"
	"goProjects/chatRoom/user"
)

type Chat struct{}

func (Chat) CreateChat(ctx context.Context, r *RequestRoom) (*Room, error) {
	//插入聊天室表
	db := user.SqlConn()
	defer db.Close()
	_, errs := db.Exec("insert into mydb.ChatRoom(userId1,userId2) values(?,?)",
		r.MyId, r.TargetId)
	if errs != nil {
		e := fmt.Errorf("创建聊天室失败:%v\n", errs.Error())
		return &Room{}, e
	}
	var roomId int32
	//查询聊天室ID并返回
	row := db.QueryRow("select id from mydb.ChatRoom where userId1=? and userId2=?", r.MyId, r.TargetId)
	if err := row.Scan(&roomId); err != nil {
		e := fmt.Errorf("获取RoomId数据库读取失败:%v\n", err.Error())
		return &Room{}, e
	}
	return &Room{RoomId: roomId}, nil
}

func (Chat) GetRoom(ctx context.Context, r *RequestRoom) (*Room, error) {
	db := user.SqlConn()
	defer db.Close()
	//获取根据我的ID获取到所有与我相关到聊天室，匹配ID2，若匹配成功则返回聊天室ID，否则
	rows, err := db.Query("select id,userId2 from mydb.ChatRoom where userId1=?", r.MyId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var roomId, userId int32
		err := rows.Scan(&roomId, &userId)
		if err != nil {
			return nil, err
		}
		if userId != r.TargetId {
			continue
		}
		room1 := Room{RoomId: roomId}
		return &room1, nil
	}
	return nil, nil
}
func (Chat) LoadHistory(ctx context.Context, r *Room) (*ChatMsgs, error) {
	DB := user.SqlConn()
	defer DB.Close()
	//查询历史20条聊天记录
	rows, err := DB.Query("select a.userId,b.name,a.chatMessage,a.createTime,a.id from mydb.ChatMessage a left join mydb.User b on a.userId=b.id order by createTime desc limit 20;")
	if err != nil {
		return nil, err
	}
	// 初始化动态结构体slice，将查询结果集写入slice，并返回
	ChatMessages := make([]*ChatMsg, 0, 20)
	for rows.Next() {
		chatMsg := ChatMsg{}
		err := rows.Scan(&chatMsg.UserId, &chatMsg.UserName, &chatMsg.ChatMessage, &chatMsg.CreateTime, &chatMsg.MessageId)
		if err != nil {
			return nil, err
		}
		ChatMessages = append(ChatMessages, &chatMsg)
	}
	return &ChatMsgs{ChatMessages: ChatMessages}, nil
}
func (Chat) SendChat(ctx context.Context, r *RequestSendChatMsg) (*Empty, error) {
	//发送聊天，将用户ID聊天信息存入库表
	db := user.SqlConn()
	defer db.Close()
	_, errs := db.Exec("insert into mydb.ChatMessage(roomId,userId,chatMessage) values(?,?,?)",
		r.RoomId, r.MyId, r.ChatMessage)
	if errs != nil {
		e := fmt.Errorf("发送聊天失败:%v\n", errs.Error())
		return &Empty{}, e
	}
	return &Empty{}, nil
}
func (Chat) GetChat(ctx context.Context, r *GetNew) (*ChatMsgs, error) {
	//获取聊天
	db := user.SqlConn()
	defer db.Close()
	rows, err := db.Query("select a.userId,b.name,a.chatMessage,a.createTime,a.id from mydb.ChatMessage a left join mydb.User b on a.userId=b.id where a.id>?", r.Lastid)
	if err != nil {
		return nil, err
	}
	// 初始化动态结构体slice，将查询结果集写入slice，并返回
	ChatMessages := make([]*ChatMsg, 0, 1000)
	for rows.Next() {
		chatMsg := ChatMsg{}
		err := rows.Scan(&chatMsg.UserId, &chatMsg.UserName, &chatMsg.ChatMessage, &chatMsg.CreateTime, &chatMsg.MessageId)
		if err != nil {
			return nil, err
		}
		ChatMessages = append(ChatMessages, &chatMsg)
	}
	return &ChatMsgs{ChatMessages: ChatMessages}, nil
}
