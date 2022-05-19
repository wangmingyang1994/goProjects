create table users(
                      `user_id` int auto_increment,
                      `user_name` varchar(10) unique not null,
                      `password` varchar(100) not null,
                      `create_at` datetime not null,
                      `user_type` varchar(6) not null,
                      primary key (`user_id`)
);

create table books(
                      `book_id` int auto_increment,
                      `book_name` varchar(10) unique not null,
                      `book_type` varchar(100) not null,
                      `book_author` varchar(100) not null,
                      `book_stock` int not null,
                      primary key (`book_id`)
);

create table `user_book_records`(
                      `record_id` int auto_increment,
                      `user_id`int  not null,
                      `book_id` int  not null,
                      `book_name` varchar(10)  not null,
                      `start_date` datetime not null,
                      `days` int not null,
                      `book_status` varchar(10)  default '0',
                      `pass_status` varchar(10) default '0',
                      primary key (`record_id`)
);

