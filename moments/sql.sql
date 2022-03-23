-- ##database_info(mysql):
-- host&port=>localhost:3306
-- name&pwd=>root:11111111
-- db=>test


-- 创建test库并使用
create database test;
use test;
-- 创建用户信息表
create table persons(
    `personId` int auto_increment,
    `name` varchar(255) unique not null,
    `sex` enum('男','女') not null default '男',
    `age` int not null,
    `tall` float not null,
    `weight` float not null,
    `fatRate` float,
    `createTime` datetime not null default current_timestamp,
    `updateTime` datetime not null default current_timestamp, 
    primary key (`personId`)
)engine =InnoDB default charset =utf8mb4;


-- 创建动态信息表
create table states(
    `statesId` int auto_increment,
    `personId` int not null,
    `content` varchar(255) not null,
    `visable` bool not null default true,
    `createTime` datetime not null default current_timestamp ,
    PRIMARY KEY(`statesId`),
    index emp_id (`personId`),
    foreign key (`personId`) references persons (`personId`)
)engine =InnoDB default charset =utf8mb4;
