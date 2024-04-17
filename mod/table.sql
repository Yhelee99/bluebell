create table `user`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL ,
    `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL ,
    `email` varchar(64) COLLATE utf8mb4_general_ci,
    `gender` tinyint(4) NOT NULL DEFAULT '0',
    `creat_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=Innodb DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

create table  `community`(
    `id` int(11) not null auto_increment,
    `community_id` int(10) unsigned not null,
    `community_name` varchar(128) collate utf8mb4_general_ci not null ,
    `introduction` varchar(256) collate utf8mb4_general_ci not null ,
    `creat_time` timestamp not null default current_timestamp,
    `update_time` timestamp not null default current_timestamp on update current_timestamp,
    primary key (`id`),
    unique key `idx_community_id` (`community_id`),
    unique key `idx_community_name` (`community_name`)
) engine=Innodb default charset=utf8mb4 collate=utf8mb4_general_ci;

