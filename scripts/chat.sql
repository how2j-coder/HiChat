CREATE TABLE users
(
    `id` BIGINT PRIMARY KEY,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    `deleted_at` TIMESTAMP NULL DEFAULT NULL,
    `username` VARCHAR(255),
    `account` VARCHAR(6) UNIQUE NOT NULL,
    `password` VARCHAR(100),
    `avatar_url` VARCHAR(255),
    `email` VARCHAR(100) UNIQUE ,
    `gender` VARCHAR(10) CHECK ( gender in ('Male', 'Female', 'Other')),
    CONSTRAINT idx_username_password UNIQUE (account, password),
    INDEX idx_deleted_at (deleted_at)
) CHARSET=utf8mb4;