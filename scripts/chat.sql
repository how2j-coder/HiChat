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
    `type`       VARCHAR(10) CHECK ( type in ('Admin', 'Ordinary')),
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

# 用户角色表
CREATE TABLE `user_role`
(
    `id`      BIGINT PRIMARY KEY,
    `role_id` BIGINT NOT NULL COMMENT '角色id',
    `user_id` BIGINT NOT NULL COMMENT '用户id',
    CONSTRAINT fk_emp_user
        FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_emp_role
        FOREIGN KEY (role_id) REFERENCES role (id)
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
    `menu_name`      VARCHAR(20) COMMENT '菜单名称',
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

CREATE DATABASE `hi_chat`;
DROP table casbin_rule, menu, role_menu, role_user, roles, users;

# 初始化数据

INSERT INTO hi_chat.platform (id, created_at, updated_at, deleted_at, platform_name,
                              platform_code, is_enabled, platform_url)
VALUES (558136173290065920, '2025-03-21 11:53:53.381',
        '2025-03-21 11:53:53.381', null, 'IAM',
        '9z9f9ZkquA', 1, 'https://how2j.online');

INSERT INTO hi_chat.menu (id, created_at, updated_at, deleted_at, parent_menu_id, menu_name, menu_code, menu_path, menu_source, is_enable, type, is_refresh, is_visible, menu_icon, sort, is_single, platform_id) VALUES (557480499400937472, '2025-03-19 16:28:28.545', '2025-03-24 15:14:04.152', null, 0, '系统管理', 'System', '/system', 'Layout', 1, 1, 1, 1, 'setting', 0, 0, 558136173290065920);
INSERT INTO hi_chat.menu (id, created_at, updated_at, deleted_at, parent_menu_id, menu_name, menu_code, menu_path, menu_source, is_enable, type, is_refresh, is_visible, menu_icon, sort, is_single, platform_id) VALUES (557480542895869952, '2025-03-19 16:28:38.915', '2025-03-25 13:40:28.469', null, 557480499400937472, '菜单管理', 'Menu', 'menu', 'views/system/menu/index.vue', 1, 1, 1, 1, 'cus:menu', 0, 0, 558136173290065920);
INSERT INTO hi_chat.menu (id, created_at, updated_at, deleted_at, parent_menu_id, menu_name, menu_code, menu_path, menu_source, is_enable, type, is_refresh, is_visible, menu_icon, sort, is_single, platform_id) VALUES (557480562357440512, '2025-03-19 16:28:43.555', '2025-03-25 13:47:12.155', null, 557480499400937472, '角色管理', 'Role', 'role', 'views/system/role/index.vue', 1, 1, 1, 1, 'cus:role', 1, 0, 558136173290065920);

INSERT INTO hi_chat.users (id, created_at, updated_at, deleted_at, username, account, password, email, avatar_url, gender, type) VALUES (557479636460638208, '2025-03-19 16:25:02.804', '2025-03-19 16:25:02.804', null, '', 'lHQbeH', '$2a$10$wAK6ETAET4ozXeszpqksZOKrB9nGRNT4kbLdkP6BtYhlr0Fv6O0wa', 'how2jCoder@linux.do', '', 'Other', 'Admin');
INSERT INTO hi_chat.users (id, created_at, updated_at, deleted_at, username, account, password, email, avatar_url, gender, type) VALUES (557479644073299968, '2025-03-19 16:25:04.618', '2025-03-19 16:25:04.618', null, '', 'eqwNdt', '$2a$10$VZSSLIDLPamtHl7clxdRge7X58mhptMSg5Q/0ygQduoC9kWwFs6pq', 'how2j@linux.do', '', 'Other', 'Ordinary');
