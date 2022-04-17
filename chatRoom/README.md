聊天服务说明:\

user服务:\
SignUp(UserRegister) returns (UserDetail);//注册接口\
SignIn(UserLogin) returns (UserDetail);//登录接口\
Logout(UserLogout) returns (UserDetail);//退出登录接口\
GetUserInfos(Needusers)returns(UserInfos);//获取在线的用户列表接口\
chat服务:\
CreateChat(RequestRoom)returns(Room);//创建聊天室接口\
GetRoom(RequestRoom)returns(Room);//查询是否已创建聊天室\
LoadHistory(Room)returns(ChatMsgs);//加载历史聊天记录\
SendChat(RequestSendChatMsg)returns(Empty);//发送聊天\
GetChat(GetNew)returns(ChatMsgs);//获取聊天

聊天服务思路：\
创建好聊天室后，每个用户发送的聊天都会根据聊天室ID保存到聊天记录表里，聊天记录ID自增。每隔一秒，聊天室的双方都会获取一次比上次获取的聊天记录Id更大的聊天

server入口位置:\
goProjects/chatRoom/server/server.go

client示例:\
cd chatRoom\
注册\
go run client.go signUp [name] [password]\
登录\
go run client.go login [name] [password]\
获取用户列表\
go run client.go getUsers [page] \
开始聊天\
go run client.go startChat [userId] [targetUserId]\