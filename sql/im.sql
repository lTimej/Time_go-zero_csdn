SET
@@auto_increment_increment=9;

CREATE TABLE `contact` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `owner_id` char(20)  DEFAULT NULL,
  `target_id` char(20)  DEFAULT NULL,
  `type` bigint(20) DEFAULT NULL,
  `desc` varchar(255),
  PRIMARY KEY (`id`),
  KEY `idx_contact_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=185 DEFAULT CHARSET=utf8;

CREATE TABLE `message` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `form_id` longtext,
  `target_id` char(20),
  `user_id` char(20),
  `type` longtext,
  `media` bigint(20) DEFAULT NULL,
  `content` longtext,
  `create_time` bigint(20) unsigned,
	`read_time`   bigint(20) unsigned,
  `pic` longtext,
  `url` longtext,
  `desc` longtext,
  `amount` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_message_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;