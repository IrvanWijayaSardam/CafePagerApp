-- --------------------------------------------------------
-- Host:                         localhost
-- Server version:               5.7.24 - MySQL Community Server (GPL)
-- Server OS:                    Win64
-- HeidiSQL Version:             10.2.0.5599
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- Dumping database structure for mantrapager_db
CREATE DATABASE IF NOT EXISTS `mantrapager_db` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `mantrapager_db`;

-- Dumping structure for table mantrapager_db.notes
CREATE TABLE IF NOT EXISTS `notes` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `date` varchar(255) DEFAULT NULL,
  `user_id` bigint(20) unsigned NOT NULL,
  `image` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_notes_user` (`user_id`),
  CONSTRAINT `fk_notes_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

-- Dumping data for table mantrapager_db.notes: ~0 rows (approximately)
/*!40000 ALTER TABLE `notes` DISABLE KEYS */;
INSERT INTO `notes` (`id`, `title`, `description`, `date`, `user_id`, `image`) VALUES
	(1, 'Test notes update 2', 'test', '11/03/2023', 1, '-');
/*!40000 ALTER TABLE `notes` ENABLE KEYS */;

-- Dumping structure for table mantrapager_db.pagers
CREATE TABLE IF NOT EXISTS `pagers` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `pager_name` varchar(255) DEFAULT NULL,
  `pager_status` tinyint(1) DEFAULT NULL,
  `ss_id` varchar(255) DEFAULT NULL,
  `user_id` bigint(20) unsigned NOT NULL,
  `ss_id_pass` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_pagers_user` (`user_id`),
  CONSTRAINT `fk_pagers_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

-- Dumping data for table mantrapager_db.pagers: ~4 rows (approximately)
/*!40000 ALTER TABLE `pagers` DISABLE KEYS */;
INSERT INTO `pagers` (`id`, `pager_name`, `pager_status`, `ss_id`, `user_id`, `ss_id_pass`) VALUES
	(1, 'Pager 1', 0, 'Sukasto House', 1, 'AminiVan12!'),
	(2, 'Pager 1 Updated', 0, 'Sukasto House', 1, 'AminiVan12!'),
	(3, 'Pager 1 Updated', 0, 'Sukasto House', 1, 'AminiVan12!'),
	(4, 'Pager 1 Updated', 0, 'Sukasto House', 1, 'AminiVan12!');
/*!40000 ALTER TABLE `pagers` ENABLE KEYS */;

-- Dumping structure for table mantrapager_db.users
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` longtext NOT NULL,
  `profile` varchar(255) DEFAULT NULL,
  `jk` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_users_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

-- Dumping data for table mantrapager_db.users: ~0 rows (approximately)
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` (`id`, `name`, `email`, `password`, `profile`, `jk`) VALUES
	(1, 'Irvan Wijaya', 'aminivan@gmail.com', '$2a$04$ZBC7voeezeuQrrxI70FaeuHeKLamDDTWmlSEa/0TRLhfWYybrhb6i', '-', 'L'),
	(2, 'Irvan Wijaya', 'aminidev@gmail.com', '$2a$04$cf6s9J3cgFa.Ol04FVgoVeLDNBdwZ4ZWfBm3et935s5aAsRw9HI02', '-', 'L');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
