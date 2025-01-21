CREATE TABLE `message` (
	`id` int NOT NULL AUTO_INCREMENT,
	`chat_id` int NOT NULL,
	`sender_id` int NOT NULL,
	`content` text,
	`has_files` tinyint(1) DEFAULT '0',
	`file_count` int DEFAULT '0',
	`created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NULL DEFAULT NULL,
	`deleted_at` TIMESTAMP NULL DEFAULT NULL,
	PRIMARY KEY (`id`),
	KEY `chat_id` (`chat_id`),
	KEY `sender_id` (`sender_id`),
	CONSTRAINT `message_ibfk_1` FOREIGN KEY (`chat_id`) REFERENCES `chat` (`id`) ON DELETE CASCADE,
	CONSTRAINT `message_ibfk_2` FOREIGN KEY (`sender_id`) REFERENCES `user` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci