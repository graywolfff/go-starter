CREATE DATABASE banking;
USE banking; 

DROP TABLE IF EXISTS `customer`;
CREATE TABLE `customer` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `date_of_birth` date NOT NULL,
  `city` varchar(100) NOT NULL,
  `zipcode` varchar(10) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `customer` (`id`, `name`, `date_of_birth`, `city`, `zipcode`, `status`) VALUES
(2000,	'Steve',	'1978-12-15',	'Delhi',	'110075',	1),
(2001,	'Arian',	'1978-12-15',	'Newburgh, NY',	'12550',	1),
(2002,	'Hadley',	'1978-12-15',	'Delhi',	'07631',	0),
(2003,	'Ben',	'1988-04-30',	'Manchester, NH',	'03102',	0),
(2004,	'Nina',	'1988-05-14',	'Clarkston, MI',	'48348',	1),
(2005,	'Osmam',	'1988-11-08',	'Hyattsville, MD',	'20782',	0);

DROP TABLE IF EXISTS `account`;
CREATE TABLE `account` (
  `id` int NOT NULL AUTO_INCREMENT,
  `customer_id` int NOT NULL,
  `openning_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `account_type` varchar(10) NOT NULL,
  `pin` varchar(10) NOT NULL,
  `status` tinyint NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `account_fk` (`customer_id`),
  CONSTRAINT `account_fk` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `account` (`id`, `customer_id`, `openning_at`, `account_type`, `pin`, `status`) VALUES
(95470,	2000,	'2024-03-17 17:37:11',	'Saving',	'1075',	1),
(95471,	2001,	'2024-03-17 17:37:33',	'Saving',	'1255',	1),
(95472,	2002,	'2024-03-17 17:38:05',	'Checking',	'0763',	1),
(95473,	2000,	'2024-03-17 17:38:26',	'Saving',	'0310',	1),
(95474,	2004,	'2024-03-17 17:38:49',	'Checking',	'4834',	1),
(95475,	2005,	'2024-03-17 17:40:05',	'Saving',	'2078',	0);



DROP TABLE IF EXISTS `transaction`;
CREATE TABLE `transaction` (
  `id` int NOT NULL AUTO_INCREMENT,
  `account_id` int NOT NULL,
  `amount` int NOT NULL,
  `transaction_type` varchar(10) NOT NULL,
  `transaction_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `transaction_fk` (`account_id`),
  CONSTRAINT `transaction_fk` FOREIGN KEY (`account_id`) REFERENCES `account` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
