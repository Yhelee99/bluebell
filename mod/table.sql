--创建用户表
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

--创建社区表
create table  `community`(
    `id` int(11) not null auto_increment,
    `community_id` int(10) unsigned not null,
    `community_name` varchar(128) collate utf8mb4_general_ci not null ,
    `introduction` varchar(256) collate utf8mb4_general_ci not null ,
    `creat_time` timestamp not null DEFAULT  current_timestamp,
    `update_time` timestamp not null DEFAULT  current_timestamp on update current_timestamp,
    primary key (`id`),
    unique key `idx_community_id` (`community_id`),
    unique key `idx_community_name` (`community_name`)
) engine=Innodb default charset=utf8mb4 collate=utf8mb4_general_ci;

--创建帖子表
create table `post`(
    `id` bigint(20) not null auto_increment,
    `post_id` bigint(20) not null COMMENT '帖子id',
    `title` varchar(128) COLLATE utf8mb4_general_ci not null COMMENT '标题',
    `content` varchar(8192) COLLATE utf8mb4_general_ci not null COMMENT '内容',
    `author_id` bigint(20) not null COMMENT '作者的用户id',
    `community_id` bigint(20) not null COMMENT '所属的社区id',
    `status` tinyint(4) not null default '1' COMMENT '帖子状态',
    `create_time` timestamp not null DEFAULT  current_timestamp COMMENT '创建时间',
    `update_time` timestamp not null DEFAULT current_timestamp on update current_timestamp COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_post_id` (`post_id`),
    KEY `idx_author_id` (`author_id`),
    KEY `idx_community_id` (`community_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
