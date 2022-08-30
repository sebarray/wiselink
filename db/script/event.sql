CREATE TABLE `event` (
	`id` VARCHAR(36) NOT NULL COLLATE 'utf8_general_ci',
	`Description` LONGTEXT NOT NULL COLLATE 'utf8_general_ci',
	`descriptionshort` TINYTEXT NOT NULL COLLATE 'utf8_general_ci',
	`state` VARCHAR(50) NOT NULL DEFAULT 'borrador' COLLATE 'utf8_general_ci',
	`organizer` VARCHAR(60) NOT NULL COLLATE 'utf8_general_ci',
	`place` VARCHAR(60) NOT NULL COLLATE 'utf8_general_ci',
	`date` DATETIME NOT NULL,
	PRIMARY KEY (`id`) USING BTREE
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB
;