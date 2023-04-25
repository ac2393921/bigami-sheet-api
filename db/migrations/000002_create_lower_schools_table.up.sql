CREATE TABLE `lower_schools` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `school_id` int(11) NOT NULL,
  `name` VARCHAR(191) NULL,
  `style` TEXT NULL,
  `enemy` VARCHAR(191) NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`) ON DELETE CASCADE
);