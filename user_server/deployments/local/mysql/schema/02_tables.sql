USE hands_on_go;

CREATE TABLE `hands_on_go`.`users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `first_name` varchar(100) NOT NULL,
  `last_name` varchar(100) NOT NULL,
  `age` int(10) unsigned NOT NULL,
  `phone_number` varchar(255) NOT NULL,
  `phone_verification_status` tinyint(1) NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4;

