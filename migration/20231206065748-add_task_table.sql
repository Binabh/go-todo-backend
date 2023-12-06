
-- +migrate Up
CREATE TABLE IF NOT EXISTS `tasks` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,  
  `title` VARCHAR(20) NOT NULL,
  `description` VARCHAR(255),
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS `tasks`;