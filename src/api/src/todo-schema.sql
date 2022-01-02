  CREATE TABLE `todos` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `message` VARCHAR(200) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `message_UNIQUE` (`message` ASC));