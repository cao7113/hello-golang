CREATE TABLE `users` ( 
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT PRIMARY KEY, 
  `name` BIGINT(20) NOT NULL, 
  `email` varchar(100) NOT NULL,
  `birthday` date,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  UNIQUE KEY `user_email_index` (`email`) 
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;