protoc -I .  user.proto  --go_out=plugins=grpc:.
protoc -I .  chat.proto  --go_out=plugins=grpc:.



-- 创建用户表
create table User(
    `id` int not null,
    `name` varchar(10) unique not null,
    `password` varchar(100) not null,
    `status` int default 0,
    `createTime` datetime not null default current_timestamp ,
    primary key (`id`)
);

-- 创建聊天室表
create table ChatRoom(
    `id` int auto_increment,
    `userId1` int not null,
    `userId2` int not null,
    `createTime` datetime not null default current_timestamp ,
    primary key (`id`)
);

-- 创建聊天记录表
create table ChatMessage(
    `id` int auto_increment,
    `roomId` int not null,
    `userId` int not null,
    `chatMessage` varchar(255) not null,
    `createTime` datetime not null default current_timestamp,
    primary key (`id`)
);