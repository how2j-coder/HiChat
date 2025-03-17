CREATE TABLE `users`
(
    `id`         BIGINT PRIMARY KEY,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP                             NOT NULL,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    `deleted_at` TIMESTAMP                                                       NULL DEFAULT NULL,
    `username`   VARCHAR(255),
    `account`    VARCHAR(6) UNIQUE                                               NOT NULL,
    `password`   VARCHAR(100),
    `avatar_url` VARCHAR(255),
    `email`      VARCHAR(100) UNIQUE,
    `gender`     VARCHAR(10) CHECK ( gender in ('Male', 'Female', 'Other')),
    CONSTRAINT idx_username_password UNIQUE (account, password),
    INDEX idx_deleted_at (deleted_at)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4;

# 角色表
CREATE TABLE `role`
(
    `id`         BIGINT PRIMARY KEY,
    `created_at` TIMESTAMP    DEFAULT CURRENT_TIMESTAMP                             NOT NULL,
    `updated_at` TIMESTAMP    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    `deleted_at` TIMESTAMP                                                          NULL DEFAULT NULL,
    `role_name`  VARCHAR(56)  DEFAULT NULL COMMENT '角色名称',
    `remark`     VARCHAR(255) DEFAULT NULL COMMENT '备注'
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4;

# 用户于角色表
CREATE TABLE `user_role`
(
    `id`         BIGINT PRIMARY KEY,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP                             NOT NULL,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    `deleted_at` TIMESTAMP                                                       NULL DEFAULT NULL,
    `role_id`    BIGINT                                                          NOT NULL COMMENT '角色id',
    `role_code`  VARCHAR(6)                                                      NOT NULL COMMENT 'code',
    `user_id`    BIGINT                                                          NOT NULL COMMENT '用户id'
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4;

# 角色于权限表
CREATE TABLE `auth_role`
(
    `id`         BIGINT PRIMARY KEY,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP                             NOT NULL,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    `deleted_at` TIMESTAMP                                                       NULL DEFAULT NULL
);

CREATE TABLE `casbin_rule`
(
    `ptype` VARCHAR(10)  NULL DEFAULT NULL COMMENT '策略类型，通常是 p（策略规则）或 g（分组规则）',
    `v0`    VARCHAR(256) NULL DEFAULT NULL COMMENT '策略主体，表示用户或角色（sub）',
    `v1`    VARCHAR(256) NULL DEFAULT NULL COMMENT '策略对象，表示资源（obj）',
    `v2`    VARCHAR(256) NULL DEFAULT NULL COMMENT '策略操作，表示操作行为（act）',
    `v3`    VARCHAR(256) NULL DEFAULT NULL COMMENT 'Casbin 允许最多 6 个参数扩展使用',
    `v4`    VARCHAR(256) NULL DEFAULT NULL,
    `v5`    VARCHAR(256) NULL DEFAULT NULL
);

DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`
(
    `id`             BIGINT PRIMARY KEY,
    `created_at`     TIMESTAMP             DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `updated_at`     TIMESTAMP             DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    `deleted_at`     TIMESTAMP    NULL     DEFAULT NULL,
    `parent_menu_id` BIGINT COMMENT '父级菜单ID',
    `menu_name` VARCHAR(20) COMMENT '菜单名称',
    `menu_code`      VARCHAR(20)  NOT NULL COMMENT '菜单唯一Code',
    `menu_path`      VARCHAR(255) NOT NULL COMMENT '菜单对于地址',
    `menu_source`    VARCHAR(255) NOT NULL COMMENT '菜单对应前端的文件页面地址',
    `is_enable`      TINYINT      NOT NULL DEFAULT 1 COMMENT '菜单状态（1：启用，0：禁用）',
    `type`           TINYINT      NOT NULL DEFAULT 1 COMMENT '菜单类型（1：菜单，2：按钮，3：其他）',
    `is_refresh`     TINYINT      NOT NULL DEFAULT 1 COMMENT '页面刷新（1：刷新，0：不刷新）',
    `is_visible`     TINYINT               DEFAULT 1 COMMENT '是否可见(0 隐藏 1 显示)',
    UNIQUE KEY `idx_menu_code` (`menu_code`)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4;
