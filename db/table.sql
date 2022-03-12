CREATE TABLE `users` (
  `id` varchar(36) NOT NULL,
  `display_name` varchar(256) NOT NULL,
  `user_name` varchar(256) NOT NULL,
  `avatar` text,
  `created_at` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `todos` (
  `id` bigint unsigned AUTO_INCREMENT,
  `user_id` varchar(36),
  `category` varchar(36),
  `done` boolean DEFAULT false,
  `priority` bigint NOT NULL,
  `title` varchar(128) NOT NULL,
  `description` varchar(256),
  `created_at` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY idx_user_id (`user_id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- The following will not work 😢
-- ALTER TABLE todos ADD FOREIGN KEY (`user_id`) REFERENCES `users`(`id`);

INSERT INTO users (id, display_name, user_name) VALUES ('hogehoge', 'しまぶー', 'shimabu');