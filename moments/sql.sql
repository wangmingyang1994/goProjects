-- create table
create table persons(
    `personId` int auto_increment,
    `name` varchar(255) unique not null,
    `sex` enum('男','女') not null default '男',
    `age` int not null,
    `tall` float not null,
    `weight` float not null,
    `fatRate` float,
    primary key (`personId`)
    
)engine =InnoDB default charset =utf8mb4;

insert into persons(name, sex, age, tall, weight, fatRate) 
values('xiaoming','男', 18,1.75,70,0.2834);


create table states(
    `statesId` int auto_increment,
    `personId` int not null,
    `content` varchar(255) not null,
    `visable` bool not null default true,
    `createTime` datetime not null default current_timestamp ,
    PRIMARY KEY(`statesId`)
)engine =InnoDB default charset =utf8mb4;