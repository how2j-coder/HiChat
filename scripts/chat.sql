CREATE TABLE `users`
(
    `id` BIGINT PRIMARY KEY,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    `deleted_at` TIMESTAMP NULL DEFAULT NULL,
    `username` VARCHAR(255),
    `account` VARCHAR(6) UNIQUE NOT NULL,
    `password` VARCHAR(100),
    `avatar_url` VARCHAR(255),
    `email` VARCHAR(100) UNIQUE ,
    `gender` VARCHAR(10) CHECK ( gender in ('Male', 'Female', 'Other')),
    CONSTRAINT idx_username_password UNIQUE (account, password),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4;

# 角色表
CREATE TABLE `role`
(
    `id` BIGINT PRIMARY KEY,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    `deleted_at` TIMESTAMP NULL DEFAULT NULL,
    `role_name` VARCHAR(56) DEFAULT NULL COMMENT '角色名称',
    `remark` VARCHAR(255) DEFAULT NULL COMMENT '备注'
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;

# 用户于角色表
CREATE TABLE `user_role`
(
    `id` BIGINT PRIMARY KEY,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    `deleted_at` TIMESTAMP NULL DEFAULT NULL,
    `role_id` BIGINT NOT NULL COMMENT '角色id',
    `user_id` BIGINT NOT NULL COMMENT '用户id'
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;

# 角色于权限表
CREATE TABLE `auth_role`
(
    `id` BIGINT PRIMARY KEY,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    `deleted_at` TIMESTAMP NULL DEFAULT NULL,
    `role_id` BIGINT NOT NULL COMMENT '角色id',


)