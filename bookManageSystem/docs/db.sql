create table users(
                      `user_id` int auto_increment comment '用户ID',
                      `user_name` varchar(10) unique not null comment '用户名称',
                      `password` varchar(100) not null comment '用户密码',
                      `create_at` datetime not null comment '用户创建时间',
                      `user_type` varchar(6) not null comment '用户类型 0普通用户 1管理员',
                      primary key (`user_id`)
);

create table books(
                      `book_id` int auto_increment comment '书籍ID',
                      `book_name` varchar(10) unique not null comment '书籍名称',
                      `book_type` varchar(100) not null comment '书籍类型',
                      `book_author` varchar(100) not null comment '书籍作者',
                      `book_stock` int not null comment '书籍库存',
                      primary key (`book_id`)
);

create table `user_book_records`(
                      `record_id` int auto_increment comment '记录ID',
                      `user_id`int  not null comment '用户ID',
                      `book_id` int  not null comment '书籍ID',
                      `book_name` varchar(10)  not null comment '书籍名称',
                      `start_date` datetime not null comment '借书日期',
                      `days` int not null comment '借书时长：日',
                      `book_status` varchar(10)  default '0' comment '书籍状态 0未持有 1已归还 2未归还',
                      `pass_status` varchar(10) default '0' comment '审核状态 0未审核 1已通过 2未通过',
                      primary key (`record_id`)
);

