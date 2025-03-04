CREATE TABLE users
(
    `id` BIGINT PRIMARY KEY,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    `deleted_at` TIMESTAMP NULL DEFAULT NULL,
    `name` VARCHAR(255),
    `password_hash` VARCHAR(100) UNIQUE,
    `avatar_url` VARCHAR(255),
    `email` VARCHAR(100) UNICODE,
    `gender` VARCHAR(10) CHECK ( gender in ('Male', 'Female', 'Other')),
    INDEX idx_deleted_at (deleted_at)
) CHARSET=utf8mb4;