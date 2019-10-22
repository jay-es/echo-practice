-- Generation time: Mon, 21 Oct 2019 14:42:10 +0000
-- Host: mysql.hostinger.ro
-- DB name: u574849695_25
/*!40030 SET NAMES UTF8 */;
/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `last_name` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `email` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `birthdate` date NOT NULL,
  `added` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

INSERT INTO `users` VALUES ('1','Guadalupe','Bergstrom','ebergstrom@example.com','1993-09-22','2014-01-21 05:47:01'),
('2','Maurine','Powlowski','schuster.robert@example.com','1990-05-08','1993-12-10 05:14:37'),
('3','Sonny','Weissnat','gturcotte@example.org','1999-07-07','1982-02-13 03:02:07'),
('4','Amelie','Gislason','bgorczany@example.net','2016-10-04','1971-05-30 03:52:58'),
('5','Susie','Pacocha','sydney94@example.org','2008-12-23','2006-12-24 11:10:41'),
('6','Vella','Heidenreich','hickle.sophie@example.net','2013-01-19','2000-04-17 08:34:45'),
('7','Urban','Bruen','koss.maryam@example.com','1999-08-01','1974-03-03 03:14:11'),
('8','Aletha','Stracke','virginia.friesen@example.org','2012-07-12','2008-03-28 19:04:28'),
('9','Ila','Doyle','darion26@example.com','2014-11-04','2007-06-21 23:25:44'),
('10','Cydney','Kirlin','schinner.ursula@example.com','1987-06-03','1976-12-14 13:16:05'); 




/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

