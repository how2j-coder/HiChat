CREATE DATABASE IF NOT EXISTS hi_chat;

INSERT INTO hi_chat.users (id, created_at, updated_at, deleted_at, name, pass_word, avatar, gender, phone, email, identity, client_ip, client_port, salt, login_time, heart_beat_time, login_out_time, is_login_out, device_info) VALUES ('7280510090929717248', '2025-01-02 17:07:51.014', '2025-01-02 17:07:51.014', null, 'how2j', 'eaae5777bb6148166153b1b7587cf23b$753730105', '', 'male', '', 'how2j.online@online.com', '', '', '', '753730105', null, null, null, 0, '');



INSERT INTO hi_chat.platform (
id, created_at, updated_at, deleted_at, platform_name,
platform_code, platform_url, version, is_enable)
VALUES ('7280756958339219456', '2025-01-03 09:28:48.792',
        '2025-01-03 09:28:48.792', null,
        '后台管理系统',
        'plat_4Ng4ssFLh8Tgo8WbTord2Cw4',
        'https://127.0.0.1:5217', '', 1);


INSERT INTO hi_chat.menus (id, created_at, updated_at, deleted_at, platform_id, parent_menu_id, menu_name, menu_code, menu_type, menu_path, menu_file_path, is_visible, is_enabled, is_refresh, sort_order, menu_icon, is_single) VALUES ('7281858506603577344', '2025-01-06 10:25:58.364', '2025-01-06 10:25:58.364', null, '7280756958339219456', '', '系统管理', 'SystemLayout', 0, '/system', 'Layout', 1, 1, 0, 0, 'Setting', null);
INSERT INTO hi_chat.menus (id, created_at, updated_at, deleted_at, platform_id, parent_menu_id, menu_name, menu_code, menu_type, menu_path, menu_file_path, is_visible, is_enabled, is_refresh, sort_order, menu_icon, is_single) VALUES ('7281858708215382016', '2025-01-06 10:26:46.431', '2025-01-06 10:26:46.431', null, '7280756958339219456', '7281858506603577344', '用户管理', 'User', 1, 'user', 'views/system/user/index.vue', 1, 1, 0, 0, null, null);
INSERT INTO hi_chat.menus (id, created_at, updated_at, deleted_at, platform_id, parent_menu_id, menu_name, menu_code, menu_type, menu_path, menu_file_path, is_visible, is_enabled, is_refresh, sort_order, menu_icon, is_single) VALUES ('7281858754625355776', '2025-01-06 10:26:57.497', '2025-01-06 10:26:57.497', null, '7280756958339219456', '7281858506603577344', '单位管理', 'Unit', 1, 'unit', 'views/system/unit/index.vue', 1, 1, 0, 1, null, null);
INSERT INTO hi_chat.menus (id, created_at, updated_at, deleted_at, platform_id, parent_menu_id, menu_name, menu_code, menu_type, menu_path, menu_file_path, is_visible, is_enabled, is_refresh, sort_order, menu_icon, is_single) VALUES ('7283001505781858304', '2025-01-09 14:07:50.612', '2025-01-09 14:07:50.612', null, '7280756958339219456', '7281858506603577344', '菜单管理', 'Menu', 1, 'menu', 'views/system/menu/index.vue', 1, 1, 0, 2, null, null);
INSERT INTO hi_chat.menus (id, created_at, updated_at, deleted_at, platform_id, parent_menu_id, menu_name, menu_code, menu_type, menu_path, menu_file_path, is_visible, is_enabled, is_refresh, sort_order, menu_icon, is_single) VALUES ('7283046949706743808', '2025-01-09 17:08:25.280', '2025-01-09 17:08:25.280', null, '7283045581981958144', '', 'create', 'CreateSystem', 0, '/create', 'Layout', 1, 1, 0, 0, null, null);
