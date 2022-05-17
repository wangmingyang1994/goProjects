create table users(
                      `user_id` int auto_increment,
                      `user_name` varchar(10) unique not null,
                      `password` varchar(100) not null,
                      `create_at` datetime not null,
                      `user_type` varchar(6) not null,
                      primary key (`user_id`)
);
